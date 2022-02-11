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

package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	allinonepb "github.com/cita-cloud/operator-proxy/api/allinone"
	chainpb "github.com/cita-cloud/operator-proxy/api/chain"
	nodepb "github.com/cita-cloud/operator-proxy/api/node"
	k8sclient "github.com/cita-cloud/operator-proxy/server/kubeapi"
	accountservice "github.com/cita-cloud/operator-proxy/server/service/account"
	allinoneservice "github.com/cita-cloud/operator-proxy/server/service/allinone"
	chainservice "github.com/cita-cloud/operator-proxy/server/service/chain"
	nodeservice "github.com/cita-cloud/operator-proxy/server/service/node"
)

const (
	port = ":8090"
)

func main() {
	err := k8sclient.InitK8sClient()
	if err != nil {
		log.Fatalf("failed to init k8s client: %v", err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	chainServer := chainservice.NewChainServer()
	chainpb.RegisterChainServiceServer(s, chainServer)

	accountServer := accountservice.NewAccountServer()
	accountpb.RegisterAccountServiceServer(s, accountServer)

	nodeServer := nodeservice.NewNodeServer()
	nodepb.RegisterNodeServiceServer(s, nodeServer)

	allInOneServer := allinoneservice.NewAllInOneServer()
	allinonepb.RegisterAllInOneServiceServer(s, allInOneServer)

	log.Printf("Starting gRPC listener on port " + port)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
