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

	pb "github.com/cita-cloud/operator-proxy/api/chain"
	"google.golang.org/grpc"
)

type ChainInterface interface {
	Init(ctx context.Context, chain *pb.Chain) (*pb.ChainSimpleResponse, error)
	Online(ctx context.Context, request *pb.ChainOnlineRequest) (*pb.ChainSimpleResponse, error)
	List(ctx context.Context, request *pb.ListChainRequest) (*pb.ChainList, error)
	Describe(ctx context.Context, request *pb.ChainDescribeRequest) (*pb.ChainDescribeResponse, error)
}

type chain struct {
	remote   pb.ChainServiceClient
	callOpts []grpc.CallOption
}

func (c chain) Init(ctx context.Context, in *pb.Chain) (*pb.ChainSimpleResponse, error) {
	resp, err := c.remote.Init(ctx, in, c.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c chain) Online(ctx context.Context, request *pb.ChainOnlineRequest) (*pb.ChainSimpleResponse, error) {
	resp, err := c.remote.Online(ctx, request, c.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c chain) List(ctx context.Context, in *pb.ListChainRequest) (*pb.ChainList, error) {
	resp, err := c.remote.List(ctx, in, c.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c chain) Describe(ctx context.Context, request *pb.ChainDescribeRequest) (*pb.ChainDescribeResponse, error) {
	resp, err := c.remote.Describe(ctx, request, c.callOpts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}


func NewChain(c *Client) ChainInterface {
	api := &chain{remote: pb.NewChainServiceClient(c.conn)}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
