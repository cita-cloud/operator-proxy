/*
 * Copyright Rivtower Technologies LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package chain

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"

	pb "github.com/cita-cloud/operator-proxy/api/chain"
	"github.com/cita-cloud/operator-proxy/server/kubeapi"
	"github.com/cita-cloud/operator-proxy/server/service/resource"
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
	chainConfig.Spec.ConsensusType = convertProtoToSpec(chain.GetConsensusType())
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

func (c chainServer) Describe(ctx context.Context, request *pb.ChainDescribeRequest) (*pb.ChainDescribeResponse, error) {
	chain := &citacloudv1.ChainConfig{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: request.Name, Namespace: request.Namespace}, chain); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get chain cr", err)
	}

	nodeList, err := resource.GetNodeList(ctx, request.GetNamespace(), request.GetName())
	if err != nil {
		return nil, err
	}

	return &pb.ChainDescribeResponse{
		Name:            request.Name,
		Namespace:       request.Namespace,
		Id:              chain.Spec.Id,
		Timestamp:       chain.Spec.Timestamp,
		PrevHash:        chain.Spec.PrevHash,
		BlockInterval:   chain.Spec.BlockInterval,
		BlockLimit:      chain.Spec.BlockLimit,
		EnableTls:       chain.Spec.EnableTLS,
		ConsensusType:   convertSpecToProto(chain.Spec.ConsensusType),
		NetworkImage:    chain.Spec.NetworkImage,
		ConsensusImage:  chain.Spec.ConsensusImage,
		ExecutorImage:   chain.Spec.ExecutorImage,
		StorageImage:    chain.Spec.StorageImage,
		ControllerImage: chain.Spec.ControllerImage,
		KmsImage:        chain.Spec.KmsImage,
		Nodes:           nodeList,
	}, status.New(codes.OK, "").Err()
}

func NewChainServer() pb.ChainServiceServer {
	return &chainServer{}
}

// todo: modify
func convertProtoToSpec(consensusType pb.ConsensusType) citacloudv1.ConsensusType {
	if consensusType == pb.ConsensusType_Raft {
		return citacloudv1.Raft
	} else if consensusType == pb.ConsensusType_BFT {
		return citacloudv1.BFT
	}
	return ""
}

func convertSpecToProto(consensusType citacloudv1.ConsensusType) pb.ConsensusType {
	switch consensusType {
	case citacloudv1.BFT:
		return pb.ConsensusType_BFT
	case citacloudv1.Raft:
		return pb.ConsensusType_Raft
	default:
		return pb.ConsensusType_UnknownConsensusType
	}
}
