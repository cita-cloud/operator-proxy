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
