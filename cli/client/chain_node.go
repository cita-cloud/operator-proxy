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

package client

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/cita-cloud/operator-proxy/api/node"
	"google.golang.org/grpc"
)

type NodeInterface interface {
	Init(ctx context.Context, node *pb.Node) (*pb.NodeSimpleResponse, error)
	List(ctx context.Context, request *pb.ListNodeRequest) (*pb.NodeList, error)
	Start(ctx context.Context, request *pb.NodeStartRequest) (*pb.NodeSimpleResponse, error)
	Stop(ctx context.Context, request *pb.NodeStopRequest) (*emptypb.Empty, error)
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

func (n node) Stop(ctx context.Context, request *pb.NodeStopRequest) (*emptypb.Empty, error) {
	resp, err := n.remote.Stop(ctx, request, n.callOpts...)
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
