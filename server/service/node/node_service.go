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

package node

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	k8sresource "k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/pointer"

	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"
	operatorctrl "github.com/cita-cloud/cita-cloud-operator/controllers"

	pb "github.com/cita-cloud/operator-proxy/api/node"
	"github.com/cita-cloud/operator-proxy/server/kubeapi"
	"github.com/cita-cloud/operator-proxy/server/service/resource"
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
	nodeCr.Spec.LogLevel = convertProtoToSpec(node.GetLogLevel())
	nodeCr.Spec.Action = citacloudv1.NodeInitialize

	if node.GetCpuRequest() != "" || node.GetMemRequest() != "" {
		nodeCr.Spec.Resources.Requests = corev1.ResourceList{}
		nodeCr.Spec.Resources.Requests[corev1.ResourceCPU], _ = k8sresource.ParseQuantity(node.GetCpuRequest())
		nodeCr.Spec.Resources.Requests[corev1.ResourceMemory], _ = k8sresource.ParseQuantity(node.GetMemRequest())
	}

	if node.GetCpuLimit() != "" || node.GetCpuLimit() != "" {
		nodeCr.Spec.Resources.Limits = corev1.ResourceList{}
		nodeCr.Spec.Resources.Limits[corev1.ResourceCPU], _ = k8sresource.ParseQuantity(node.GetCpuLimit())
		nodeCr.Spec.Resources.Limits[corev1.ResourceMemory], _ = k8sresource.ParseQuantity(node.GetMemLimit())
	}

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
	nodeList, err := resource.GetNodeList(ctx, request.GetNamespace(), request.GetChain())
	if err != nil {
		return nil, err
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

func (n nodeServer) Stop(ctx context.Context, request *pb.NodeStopRequest) (*emptypb.Empty, error) {
	node := &citacloudv1.ChainNode{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: request.Name, Namespace: request.Namespace}, node); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get node cr", err)
	}
	if node.Spec.Action != citacloudv1.NodeStop {
		node.Spec.Action = citacloudv1.NodeStop
		if err := kubeapi.K8sClient.Update(ctx, node); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update node stop", err)
		}
	}
	return &empty.Empty{}, status.New(codes.OK, "").Err()
}

func (n nodeServer) ReloadConfig(ctx context.Context, request *pb.ReloadConfigRequest) (*emptypb.Empty, error) {
	node := &citacloudv1.ChainNode{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: request.Name, Namespace: request.Namespace}, node); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get node cr", err)
	}
	// find chain config
	chain := &citacloudv1.ChainConfig{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: node.Spec.ChainName, Namespace: request.Namespace}, chain); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get chain cr", err)
	}

	// find account
	account := &citacloudv1.Account{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: node.Spec.Account, Namespace: request.Namespace}, account); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get account cr", err)
	}

	// find node's configmap
	configmap := &corev1.ConfigMap{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: operatorctrl.GetNodeConfigName(request.Name), Namespace: request.Namespace}, configmap); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get configmap cr", err)
	}

	var cnService *operatorctrl.ChainNodeService

	if chain.Spec.EnableTLS {
		// todo: reflect
		// get chain ca secret
		caSecret := &corev1.Secret{}
		if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: operatorctrl.GetCaSecretName(chain.Name), Namespace: chain.Namespace}, caSecret); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get ca secret", err)
		}
		// get node secret
		nodeCertAndKeySecret := &corev1.Secret{}
		if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: operatorctrl.GetAccountCertAndKeySecretName(node.Spec.Account), Namespace: node.Namespace}, nodeCertAndKeySecret); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get node secret", err)
		}

		cnService = operatorctrl.NewChainNodeServiceForTls(chain, node, account, caSecret, nodeCertAndKeySecret)
	} else {
		cnService = operatorctrl.NewChainNodeServiceForP2P(chain, node, account)
	}

	configmap.Data = map[string]string{
		operatorctrl.NodeConfigFile: cnService.GenerateNodeConfig(),
	}

	if err := kubeapi.K8sClient.Update(ctx, configmap); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update configmap", err)
	}

	return &empty.Empty{}, status.New(codes.OK, "").Err()
}

func (n nodeServer) Delete(ctx context.Context, request *pb.NodeDeleteRequest) (*emptypb.Empty, error) {
	node := &citacloudv1.ChainNode{}
	if err := kubeapi.K8sClient.Get(ctx, types.NamespacedName{Name: request.Name, Namespace: request.Namespace}, node); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get node cr", err)
	}
	err := kubeapi.K8sClient.Delete(ctx, node, client.GracePeriodSeconds(0))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete node error", err)
	}
	return &empty.Empty{}, status.New(codes.OK, "").Err()
}

func NewNodeServer() pb.NodeServiceServer {
	return &nodeServer{}
}

// todo modify
func convertProtoToSpec(logLevel string) citacloudv1.LogLevel {
	if logLevel == "info" {
		return citacloudv1.Info
	} else if logLevel == "warn" {
		return citacloudv1.Warn
	}
	return ""
}
