package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	chainpb "github.com/cita-cloud/operator-proxy/api/chain"
	k8sclient "github.com/cita-cloud/operator-proxy/server/kubeapi"
	accountservice "github.com/cita-cloud/operator-proxy/server/service/account"
	chainservice "github.com/cita-cloud/operator-proxy/server/service/chain"
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

	ccServer := chainservice.NewChainServer()
	chainpb.RegisterChainServiceServer(s, ccServer)

	accountServer := accountservice.NewAccountServer()
	accountpb.RegisterAccountServiceServer(s, accountServer)

	log.Printf("Starting gRPC listener on port " + port)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
