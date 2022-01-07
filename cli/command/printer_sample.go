package command

import (
	"fmt"

	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	pb "github.com/cita-cloud/operator-proxy/api/citacloud"
)

type simplePrinter struct{}

func (s *simplePrinter) CreateAccount(account *accountpb.Account) {
	fmt.Println("create account success")
}

func (s *simplePrinter) InitChain(chainSimple *pb.ChainConfigSimple)  {
	fmt.Println("WORLD")
}