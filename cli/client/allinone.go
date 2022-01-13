package client

import (
	"context"

	pb "github.com/cita-cloud/operator-proxy/api/allinone"
	"google.golang.org/grpc"
)

type AllInOneInterface interface {
	Create(ctx context.Context, request *pb.AllInOneCreateRequest) (*pb.AllInOneCreateResponse, error)
}

type allInOne struct {
	remote   pb.AllInOneServiceClient
	callOpts []grpc.CallOption
}

func (a allInOne) Create(ctx context.Context, request *pb.AllInOneCreateRequest) (*pb.AllInOneCreateResponse, error) {
	resp, err := a.remote.Create(ctx, request, a.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewAllInOne(c *Client) AllInOneInterface {
	api := &allInOne{remote: pb.NewAllInOneServiceClient(c.conn)}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
