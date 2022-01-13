package allinone

import (
	"context"
	"fmt"
	"time"

	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"
	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	pb "github.com/cita-cloud/operator-proxy/api/allinone"
	chainpb "github.com/cita-cloud/operator-proxy/api/chain"
	nodepb "github.com/cita-cloud/operator-proxy/api/node"
	"github.com/cita-cloud/operator-proxy/server/kubeapi"
	accountsvc "github.com/cita-cloud/operator-proxy/server/service/account"
	chainsvc "github.com/cita-cloud/operator-proxy/server/service/chain"
	nodesvc "github.com/cita-cloud/operator-proxy/server/service/node"
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
	"github.com/tjfoc/gmsm/sm3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
)

var _ pb.AllInOneServiceServer = &allInOneServer{}

type allInOneServer struct {
}

func setDefault(request *pb.AllInOneCreateRequest) {
	if request.GetId() == "" {
		request.Id = generateChainId(request.GetName())
	}
	if request.GetTimestamp() == 0 {
		request.Timestamp = time.Now().UnixMicro()
	}
	if request.GetPrevHash() == "" {
		request.PrevHash = "0x0000000000000000000000000000000000000000000000000000000000000000"
	}
	if request.GetBlockInterval() == 0 {
		request.BlockInterval = 3
	}
	if request.GetBlockLimit() == 0 {
		request.BlockLimit = 100
	}
	if request.GetNetworkImage() == "" {
		if request.GetEnableTls() {
			request.NetworkImage = "citacloud/network_tls:v6.3.0"
		} else {
			request.NetworkImage = "citacloud/network_p2p:v6.3.0"
		}
	}
	if request.GetConsensusImage() == "" {
		if request.ConsensusType == chainpb.ConsensusType_Raft {
			request.ConsensusImage = "citacloud/consensus_raft:v6.3.0"
		} else if request.ConsensusType == chainpb.ConsensusType_BFT {
			request.ConsensusImage = "citacloud/consensus_bft:v6.3.0"
		}
	}
	if request.GetExecutorImage() == "" {
		request.ExecutorImage = "citacloud/executor_evm:v6.3.0"
	}
	if request.GetStorageImage() == "" {
		request.StorageImage = "citacloud/storage_rocksdb:v6.3.0"
	}
	if request.GetControllerImage() == "" {
		request.ControllerImage = "citacloud/controller:v6.3.0"
	}
	if request.GetKmsImage() == "" {
		request.KmsImage = "citacloud/kms_sm:v6.3.0"
	}
	if request.GetStorageSize() == 0 {
		// 10Gi
		request.StorageSize = 10737418240
	}
	if request.GetStorageClassName() == "" {
		request.StorageClassName = "nas-client-provisioner"
	}
	if request.GetLogLevel() == "" {
		request.LogLevel = "info"
	}
}

func (a allInOneServer) Create(ctx context.Context, request *pb.AllInOneCreateRequest) (*pb.AllInOneCreateResponse, error) {
	setDefault(request)
	// init chain
	initChainReq := &chainpb.Chain{
		Name:            request.GetName(),
		Namespace:       request.GetNamespace(),
		Id:              request.GetId(),
		Timestamp:       request.GetTimestamp(),
		PrevHash:        request.GetPrevHash(),
		BlockInterval:   request.GetBlockInterval(),
		BlockLimit:      request.GetBlockLimit(),
		EnableTls:       request.GetEnableTls(),
		ConsensusType:   request.GetConsensusType(),
		NetworkImage:    request.GetNetworkImage(),
		ConsensusImage:  request.GetConsensusImage(),
		ExecutorImage:   request.GetExecutorImage(),
		StorageImage:    request.GetStorageImage(),
		ControllerImage: request.GetControllerImage(),
		KmsImage:        request.GetKmsImage(),
	}
	_, err := chainsvc.NewChainServer().Init(ctx, initChainReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to init chain: %v", err)
	}
	adminPwd, err := generateAccountPassword()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate admin password: %v", err)
	}
	adminAccountReq := &accountpb.Account{
		Name:        generateAccountOrNodeName(request.GetName()),
		Namespace:   request.GetNamespace(),
		Chain:       request.GetName(),
		KmsPassword: adminPwd,
		Role:        accountpb.Role_Admin,
	}
	_, err = accountsvc.NewAccountServer().CreateAccount(ctx, adminAccountReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create admin account: %v", err)
	}

	nodeAccountNameList := make([]string, 0)

	index := int32(1)
	for index <= request.GetNodeCount() {
		accountPwd, err := generateAccountPassword()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to generate node account password: %v", err)
		}
		nodeAccountName := generateAccountOrNodeName(request.GetName())
		nodeAccountReq := &accountpb.Account{
			Name:        nodeAccountName,
			Namespace:   request.GetNamespace(),
			Chain:       request.GetName(),
			KmsPassword: accountPwd,
			Role:        accountpb.Role_Consensus,
			Domain:      "",
		}
		_, err = accountsvc.NewAccountServer().CreateAccount(ctx, nodeAccountReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create node account: %v", err)
		}
		index++
		nodeAccountNameList = append(nodeAccountNameList, nodeAccountName)
	}
	// wait admin account && node account in chain status
	err = wait.Poll(3*time.Second, 10*time.Second, func() (done bool, err error) {
		return a.checkChainAccount(ctx, request.GetNamespace(), request.GetName(), request.GetNodeCount())
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create node account: %v", err)
	}

	// set chain online
	onlineChainReq := &chainpb.ChainOnlineRequest{
		Name:      request.GetName(),
		Namespace: request.GetNamespace(),
	}
	_, err = chainsvc.NewChainServer().Online(ctx, onlineChainReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set chain online: %v", err)
	}

	// wait chain online
	err = wait.Poll(3*time.Second, 30*time.Second, func() (done bool, err error) {
		return a.checkChainOnline(ctx, request.GetNamespace(), request.GetName())
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to wait chain online: %v", err)
	}

	nodeNameList := make([]string, 0)
	for _, nodeAccountName := range nodeAccountNameList {
		nodeName := generateAccountOrNodeName(request.GetName())
		// create node
		nodeReq := &nodepb.Node{
			Name:      nodeName,
			Namespace: request.GetNamespace(),
			// todo modify this field
			Cluster:          "k8s-1",
			Account:          nodeAccountName,
			Chain:            request.GetName(),
			StorageSize:      request.GetStorageSize(),
			StorageClassName: request.GetStorageClassName(),
			LogLevel:         request.GetLogLevel(),
		}
		_, err = nodesvc.NewNodeServer().Init(ctx, nodeReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to init node: %v", err)
		}
		nodeNameList = append(nodeNameList, nodeName)
	}

	// concurrent start node
	g, _ := NewGroup(ctx)
	for _, nn := range nodeNameList {
		g.GoStartNode(a.startNode, ctx, request.GetNamespace(), nn)
	}
	err = g.Wait()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to start node: %v", err)
	}

	return &pb.AllInOneCreateResponse{
		Name:      request.GetName(),
		Namespace: request.GetNamespace(),
	}, status.New(codes.OK, "").Err()
}

func (a allInOneServer) checkChainAccount(ctx context.Context, namespace, name string, nodeCount int32) (bool, error) {
	chain := &citacloudv1.ChainConfig{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, chain); err != nil {
		return false, nil
	}
	if chain.Status.AdminAccount != nil && len(chain.Status.ValidatorAccountMap) == int(nodeCount) {
		return true, nil
	}
	return false, nil
}

func (a allInOneServer) checkChainOnline(ctx context.Context, namespace, name string) (bool, error) {
	chain := &citacloudv1.ChainConfig{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, chain); err != nil {
		return false, nil
	}
	if chain.Status.Status == citacloudv1.Online {
		return true, nil
	}
	return false, nil
}

func (a allInOneServer) startNode(ctx context.Context, namespace, name string) error {
	err := wait.Poll(3*time.Second, 60*time.Second, func() (done bool, err error) {
		return a.checkNodeInitialized(ctx, namespace, name)
	})
	if err != nil {
		return err
	}
	startNodeReq := &nodepb.NodeStartRequest{
		Name:      name,
		Namespace: namespace,
	}
	_, err = nodesvc.NewNodeServer().Start(ctx, startNodeReq)
	if err != nil {
		return err
	}
	return nil
}

func (a allInOneServer) checkNodeInitialized(ctx context.Context, namespace, name string) (bool, error) {
	node := &citacloudv1.ChainNode{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, node); err != nil {
		return false, nil
	}
	chain := &citacloudv1.ChainConfig{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: node.Spec.ChainName, Namespace: namespace}, chain); err != nil {
		return false, nil
	}
	if node.Status.Status == citacloudv1.NodeInitialized {
		if _, exist := chain.Status.NodeInfoMap[name]; exist {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, nil
}

func NewAllInOneServer() pb.AllInOneServiceServer {
	return &allInOneServer{}
}

func generateChainId(name string) string {
	h := sm3.New()
	h.Write([]byte(name))
	sum := h.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func generateAccountOrNodeName(name string) string {
	s := uuid.New().String()
	return name + "-" + s[len(s)-12:]
}

func generateAccountPassword() (string, error) {
	_, err := password.Generate(16, 4, 4, false, false)
	return "123456", err
}
