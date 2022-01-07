package service

import (
	"context"

	pb "github.com/cita-cloud/operator-proxy/api/citacloud"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ pb.CitaCloudServiceServer = &citaCloudServer{}

type citaCloudServer struct {
}

func NewCitaCloudServer() pb.CitaCloudServiceServer {
	return &citaCloudServer{}
}

func (c *citaCloudServer) CreateChainConfig(ctx context.Context, config *pb.ChainConfig) (*pb.ChainConfigSimple, error) {
	return &pb.ChainConfigSimple{
		Name:      "test-chain",
		Namespace: "cita",
	}, status.New(codes.OK, "").Err()
}
