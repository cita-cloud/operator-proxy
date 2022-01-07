package client

import (
	"context"

	pb "github.com/cita-cloud/operator-proxy/api/account"
	"google.golang.org/grpc"
)

type AccountInterface interface {
	CreateAccount(ctx context.Context, in *pb.Account) (*pb.Account, error)
	ListAccount(ctx context.Context, in *pb.ListAccountRequest) (*pb.AccountList, error)
}

type account struct {
	remote   pb.AccountServiceClient
	callOpts []grpc.CallOption
}

func (a account) ListAccount(ctx context.Context, in *pb.ListAccountRequest) (*pb.AccountList, error) {
	resp, err := a.remote.ListAccount(ctx, in, a.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a account) CreateAccount(ctx context.Context, in *pb.Account) (*pb.Account, error) {
	resp, err := a.remote.CreateAccount(ctx, in, a.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewAccount(c *Client) AccountInterface {
	api := &account{remote: pb.NewAccountServiceClient(c.conn)}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
