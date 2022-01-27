package resource

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sigs.k8s.io/controller-runtime/pkg/client"

	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"

	pb "github.com/cita-cloud/operator-proxy/api/node"
	"github.com/cita-cloud/operator-proxy/server/kubeapi"
)

func GetNodeList(ctx context.Context, namespace, chain string) ([]*pb.Node, error) {
	opts := []client.ListOption{
		client.InNamespace(namespace),
	}
	if chain != "" {
		opts = append(opts, client.MatchingFields{"spec.chainName": chain})
	}

	nodeCrList := &citacloudv1.ChainNodeList{}
	if err := kubeapi.K8sClient.List(ctx, nodeCrList, opts...); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list node cr", err)
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
			Status:           convertSpecToProto(nodeCr.Status.Status),
		}
		nodeList = append(nodeList, node)
	}
	return nodeList, nil
}

func convertSpecToProto(nodeStatus citacloudv1.NodeStatus) pb.Status {
	switch nodeStatus {
	case citacloudv1.NodeWaitChainOnline:
		return pb.Status_WaitChainOnline
	case citacloudv1.NodeInitialized:
		return pb.Status_Initialized
	case citacloudv1.NodeStarting:
		return pb.Status_Starting
	case citacloudv1.NodeRunning:
		return pb.Status_Running
	case citacloudv1.NodeWarning:
		return pb.Status_Warning
	case citacloudv1.NodeError:
		return pb.Status_Error
	case citacloudv1.NodeUpdating:
		return pb.Status_Updating
	case citacloudv1.NodeStopping:
		return pb.Status_Stopping
	case citacloudv1.NodeStopped:
		return pb.Status_Stopped
	default:
		return pb.Status_Unknown
	}
}
