package client

import (
	"context"

	pb "github.com/cita-cloud/operator-proxy/api/node"
	"google.golang.org/grpc"
)

type NodeInterface interface {
	Init(ctx context.Context, node *pb.Node) (*pb.NodeSimpleResponse, error)
	List(ctx context.Context, request *pb.ListNodeRequest) (*pb.NodeList, error)
	Start(ctx context.Context, request *pb.NodeStartRequest) (*pb.NodeSimpleResponse, error)
}

type node struct {
	remote   pb.NodeServiceClient
	callOpts []grpc.CallOption
}

func (n node) Init(ctx context.Context, node *pb.Node) (*pb.NodeSimpleResponse, error) {
	resp, err := n.remote.Init(ctx, node, n.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n node) List(ctx context.Context, request *pb.ListNodeRequest) (*pb.NodeList, error) {
	resp, err := n.remote.List(ctx, request, n.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n node) Start(ctx context.Context, request *pb.NodeStartRequest) (*pb.NodeSimpleResponse, error) {
	resp, err := n.remote.Start(ctx, request, n.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewNode(c *Client) NodeInterface {
	api := &node{remote: pb.NewNodeServiceClient(c.conn)}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
