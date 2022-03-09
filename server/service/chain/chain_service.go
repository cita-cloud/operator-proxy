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
	"github.com/cita-cloud/operator-proxy/pkg/utils"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"

	pb "github.com/cita-cloud/operator-proxy/api/chain"
	"github.com/cita-cloud/operator-proxy/server/kubeapi"
	"github.com/cita-cloud/operator-proxy/server/service/account"
	"github.com/cita-cloud/operator-proxy/server/service/resource"
)

var _ pb.ChainServiceServer = &chainServer{}

type chainServer struct {
}

func (c chainServer) setDefault(request *pb.Chain) (*citacloudv1.ChainConfig, error) {
	chainConfig := &citacloudv1.ChainConfig{}
	chainConfig.Name = request.GetName()
	chainConfig.Namespace = request.GetNamespace()

	if request.GetId() == "" {
		request.Id = utils.GenerateChainId(request.GetName())
	}
	chainConfig.Spec.Id = request.GetId()

	if request.GetTimestamp() == 0 {
		request.Timestamp = time.Now().UnixMicro()
	}
	chainConfig.Spec.Timestamp = request.GetTimestamp()

	if request.GetPrevHash() == "" {
		request.PrevHash = "0x0000000000000000000000000000000000000000000000000000000000000000"
	}
	chainConfig.Spec.PrevHash = request.GetPrevHash()

	if request.GetBlockInterval() == 0 {
		request.BlockInterval = 3
	}
	chainConfig.Spec.BlockInterval = request.GetBlockInterval()

	if request.GetBlockLimit() == 0 {
		request.BlockLimit = 100
	}
	chainConfig.Spec.BlockLimit = request.GetBlockLimit()

	chainConfig.Spec.EnableTLS = request.GetEnableTls()
	chainConfig.Spec.ConsensusType = convertProtoToSpec(request.GetConsensusType())
	// default status is Publicizing
	chainConfig.Spec.Action = citacloudv1.Publicizing

	if request.GetVersion() == "" {
		chainConfig.Spec.Version = citacloudv1.LATEST_VERSION
	}
	chainConfig.Spec.Version = request.GetVersion()

	exactVersion, err := chainConfig.GetExactVersion()
	if err != nil {
		return nil, err
	}

	defaultImageInfo := citacloudv1.VERSION_MAP[exactVersion]
	// merge
	chainConfig.MergeFromDefaultImageInfo(defaultImageInfo)

	return chainConfig, nil
}

func (c chainServer) Init(ctx context.Context, request *pb.Chain) (*pb.ChainSimpleResponse, error) {
	chainConfig, err := c.setDefault(request)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set default chain cr", err)
	}

	err = kubeapi.K8sClient.Create(ctx, chainConfig)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create chain cr", err)
	}
	return &pb.ChainSimpleResponse{
		Name:      request.GetName(),
		Namespace: request.GetNamespace(),
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
			Status:    convertStatusFromSpecToProto(chainCr.Status.Status),
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

	adminAccount, err := account.GetAdminAccountByChain(ctx, request.GetNamespace(), request.GetName())
	if err != nil {
		return nil, err
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
		Status:          convertStatusFromSpecToProto(chain.Status.Status),
		Nodes:           nodeList,
		AdminAccount:    adminAccount,
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

func convertStatusFromSpecToProto(chainStatus citacloudv1.ChainStatus) pb.Status {
	switch chainStatus {
	case citacloudv1.Publicizing:
		return pb.Status_Publicizing
	case citacloudv1.Online:
		return pb.Status_Online
	default:
		return pb.Status_UnknownStatus
	}
}
