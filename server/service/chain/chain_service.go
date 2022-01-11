package chain

import (
	"context"

	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pb "github.com/cita-cloud/operator-proxy/api/chain"
	"github.com/cita-cloud/operator-proxy/server/kubeapi"
)

var _ pb.ChainServiceServer = &chainServer{}

type chainServer struct {
}

func (c chainServer) Init(ctx context.Context, chain *pb.Chain) (*pb.ChainSimpleResponse, error) {
	chainConfig := &citacloudv1.ChainConfig{}
	chainConfig.Name = chain.GetName()
	chainConfig.Namespace = chain.GetNamespace()
	chainConfig.Spec.Id = chain.GetId()
	chainConfig.Spec.Timestamp = chain.GetTimestamp()
	chainConfig.Spec.PrevHash = chain.GetPrevHash()
	chainConfig.Spec.BlockInterval = chain.GetBlockInterval()
	chainConfig.Spec.BlockLimit = chain.GetBlockLimit()
	chainConfig.Spec.EnableTLS = chain.GetEnableTls()
	chainConfig.Spec.ConsensusType = citacloudv1.ConsensusType(chain.GetConsensusType())
	// default status is Publicizing
	chainConfig.Spec.Action = citacloudv1.Publicizing
	chainConfig.Spec.NetworkImage = chain.GetNetworkImage()
	chainConfig.Spec.ConsensusImage = chain.GetConsensusImage()
	chainConfig.Spec.ExecutorImage = chain.GetExecutorImage()
	chainConfig.Spec.StorageImage = chain.GetStorageImage()
	chainConfig.Spec.ControllerImage = chain.GetControllerImage()
	chainConfig.Spec.KmsImage = chain.GetKmsImage()

	err := kubeapi.K8sClient.Create(ctx, chainConfig)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create chain cr", err)
	}
	return &pb.ChainSimpleResponse{
		Name:      chain.GetName(),
		Namespace: chain.GetNamespace(),
		Status:    pb.Status_Publicizing,
	}, status.New(codes.OK, "").Err()
}

func (c chainServer) List(ctx context.Context, request *pb.ListChainRequest) (*pb.ChainList, error) {
	chainCrList := &citacloudv1.ChainConfigList{}
	chainCrOpts := []client.ListOption{
		client.InNamespace(request.GetNamespace()),
	}
	if err := kubeapi.K8sClient.List(ctx, chainCrList, chainCrOpts...); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list chain cr", err)
	}
	chainList := make([]*pb.ChainSimpleResponse, 0)
	for _, chainCr := range chainCrList.Items {
		c := &pb.ChainSimpleResponse{
			Name:      chainCr.Name,
			Namespace: chainCr.Namespace,
			//Status:    chainCr.Status,
		}
		chainList = append(chainList, c)
	}
	return &pb.ChainList{Chains: chainList}, status.New(codes.OK, "").Err()
}

func (c chainServer) Online(ctx context.Context, request *pb.ChainOnlineRequest) (*pb.ChainSimpleResponse, error) {
	chain := &citacloudv1.ChainConfig{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: request.Name, Namespace: request.Namespace}, chain); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get chain cr", err)
	}
	if chain.Spec.Action != citacloudv1.Online {
		chain.Spec.Action = citacloudv1.Online
		if err := kubeapi.K8sClient.Update(ctx, chain); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update chain online", err)
		}
	}

	return &pb.ChainSimpleResponse{
		Name:      request.Name,
		Namespace: request.Namespace,
		Status:    pb.Status_Online,
	}, status.New(codes.OK, "").Err()
}

func NewChainServer() pb.ChainServiceServer {
	return &chainServer{}
}