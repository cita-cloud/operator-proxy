package main

import (
	"log"
	"net"

	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	citacloudpb "github.com/cita-cloud/operator-proxy/api/citacloud"
	k8sclient "github.com/cita-cloud/operator-proxy/server/kubeapi"
	"github.com/cita-cloud/operator-proxy/server/service"
	accountservice "github.com/cita-cloud/operator-proxy/server/service/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	ccServer := service.NewCitaCloudServer()
	citacloudpb.RegisterCitaCloudServiceServer(s, ccServer)

	accountServer := accountservice.NewAccountServer()
	accountpb.RegisterAccountServiceServer(s, accountServer)

	log.Printf("Starting gRPC listener on port " + port)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
