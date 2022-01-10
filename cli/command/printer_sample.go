package command

import (
	"fmt"
	"strings"

	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	chainpb "github.com/cita-cloud/operator-proxy/api/chain"
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

func (s *simplePrinter) InitChain(response *chainpb.ChainSimpleResponse) {
	fmt.Println(fmt.Sprintf("init chain [%s/%s] success", response.GetNamespace(), response.GetName()))
}

func (s *simplePrinter) ListChain(list *chainpb.ChainList) {
	_, rows := makeChainListTable(list)
	for _, row := range rows {
		fmt.Println(strings.Join(row, ", "))
	}
}
