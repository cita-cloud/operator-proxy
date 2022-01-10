package command

import (
	"fmt"
	"strings"

	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	pb "github.com/cita-cloud/operator-proxy/api/citacloud"
)

type simplePrinter struct{}

func (s *simplePrinter) ListAccount(list *accountpb.AccountList) {
	_, rows := makeAccountListTable(list)
	for _, row := range rows {
		fmt.Println(strings.Join(row, ", "))
	}
}

func (s *simplePrinter) CreateAccount(account *accountpb.Account) {
	fmt.Println(fmt.Sprintf("create account [%s/%s] success", account.GetNamespace(), account.GetName()))
}

func (s *simplePrinter) InitChain(chainSimple *pb.ChainConfigSimple) {
	fmt.Println("WORLD")
}
