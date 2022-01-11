package node

import (
	"context"

	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pb "github.com/cita-cloud/operator-proxy/api/node"
	"github.com/cita-cloud/operator-proxy/server/kubeapi"
)

var _ pb.NodeServiceServer = &nodeServer{}

type nodeServer struct {
}

func (n nodeServer) Init(ctx context.Context, node *pb.Node) (*pb.NodeSimpleResponse, error) {
	nodeCr := &citacloudv1.ChainNode{}
	nodeCr.Name = node.GetName()
	nodeCr.Namespace = node.GetNamespace()
	nodeCr.Spec.Cluster = node.GetCluster()
	nodeCr.Spec.Account = node.GetAccount()
	nodeCr.Spec.ExternalIp = node.GetExternalIp()
	nodeCr.Spec.Port = node.GetPort()
	nodeCr.Spec.ChainName = node.GetChain()
	nodeCr.Spec.StorageSize = pointer.Int64(node.GetStorageSize())
	nodeCr.Spec.StorageClassName = pointer.String(node.GetStorageClassName())
	nodeCr.Spec.Action = citacloudv1.NodeInitialize

	err := kubeapi.K8sClient.Create(ctx, nodeCr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create node cr", err)
	}
	return &pb.NodeSimpleResponse{
		Name:      node.GetName(),
		Namespace: node.GetNamespace(),
		Status:    pb.Status_Initialized,
	}, status.New(codes.OK, "").Err()
}

func (n nodeServer) List(ctx context.Context, request *pb.ListNodeRequest) (*pb.NodeList, error) {
	nodeCrList := &citacloudv1.ChainNodeList{}
	nodeCrOpts := []client.ListOption{
		client.InNamespace(request.GetNamespace()),
	}
	if request.GetChain() != "" {
		nodeCrOpts = append(nodeCrOpts, client.MatchingFields{"spec.chain": request.GetChain()})
	}
	if err := kubeapi.K8sClient.List(ctx, nodeCrList, nodeCrOpts...); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list account cr", err)
	}
	nodeList := make([]*pb.Node, 0)
	for _, nodeCr := range nodeCrList.Items {
		node := &pb.Node{
			Name:             nodeCr.Name,
			Namespace:        nodeCr.Namespace,
			Cluster:          nodeCr.Spec.Cluster,
			Chain:            nodeCr.Spec.ChainName,
			Account:          nodeCr.Spec.Account,
			ExternalIp:       nodeCr.Spec.ExternalIp,
			Port:             nodeCr.Spec.Port,
			StorageSize:      *nodeCr.Spec.StorageSize,
			StorageClassName: *nodeCr.Spec.StorageClassName,
		}
		nodeList = append(nodeList, node)
	}
	return &pb.NodeList{Nodes: nodeList}, status.New(codes.OK, "").Err()
}

func (n nodeServer) Start(ctx context.Context, request *pb.NodeStartRequest) (*pb.NodeSimpleResponse, error) {
	node := &citacloudv1.ChainNode{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: request.Name, Namespace: request.Namespace}, node); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get node cr", err)
	}
	if node.Spec.Action != citacloudv1.NodeStart {
		node.Spec.Action = citacloudv1.NodeStart
		if err := kubeapi.K8sClient.Update(ctx, node); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update node start", err)
		}
	}
	return &pb.NodeSimpleResponse{
		Name:      request.Name,
		Namespace: request.Namespace,
		Status:    pb.Status_Starting,
	}, status.New(codes.OK, "").Err()
}

func NewNodeServer() pb.NodeServiceServer {
	return &nodeServer{}
}
