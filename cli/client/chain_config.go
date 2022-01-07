package client

import (
	"context"

	pb "github.com/cita-cloud/operator-proxy/api/citacloud"
	"google.golang.org/grpc"
)

type ChainInterface interface {
	InitChain(ctx context.Context, in *pb.ChainConfig) (*pb.ChainConfigSimple, error)
}

type chain struct {
	remote   pb.CitaCloudServiceClient
	callOpts []grpc.CallOption
}

func NewChain(c *Client) ChainInterface {
	api := &chain{remote: pb.NewCitaCloudServiceClient(c.conn)}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}

func (c *chain) InitChain(ctx context.Context,in *pb.ChainConfig) (*pb.ChainConfigSimple, error) {
	resp, err := c.remote.CreateChainConfig(ctx, in, c.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
