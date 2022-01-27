package command

import (
	"fmt"
	"strings"

	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	pb "github.com/cita-cloud/operator-proxy/api/allinone"
	chainpb "github.com/cita-cloud/operator-proxy/api/chain"
	"github.com/cita-cloud/operator-proxy/api/node"
)

type simplePrinter struct{}

func (s *simplePrinter) ListNode(list *node.NodeList) {
	_, rows := makeNodeListTable(list.Nodes)
	for _, row := range rows {
		fmt.Println(strings.Join(row, ", "))
	}
}

func (s *simplePrinter) StartNode(response *node.NodeSimpleResponse) {
	fmt.Println(fmt.Sprintf("start node [%s/%s] success", response.GetNamespace(), response.GetName()))
}

func (s *simplePrinter) InitNode(response *node.NodeSimpleResponse) {
	fmt.Println(fmt.Sprintf("init node [%s/%s] success", response.GetNamespace(), response.GetName()))
}

func (s *simplePrinter) CreateAllInOne(response *pb.AllInOneCreateResponse) {
	fmt.Println(fmt.Sprintf("create chain [%s/%s] success by one click", response.GetNamespace(), response.GetName()))
}

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

func (s *simplePrinter) OnlineChain(response *chainpb.ChainSimpleResponse) {
	fmt.Println(fmt.Sprintf("online chain [%s/%s] success", response.GetNamespace(), response.GetName()))
}

func (s *simplePrinter) ListChain(list *chainpb.ChainList) {
	_, rows := makeChainListTable(list)
	for _, row := range rows {
		fmt.Println(strings.Join(row, ", "))
	}
}

func (s *simplePrinter) DescribeChain(response *chainpb.ChainDescribeResponse) {
	panic("implement me")
}
