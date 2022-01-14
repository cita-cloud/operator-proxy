package command

import (
	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	allinonepb "github.com/cita-cloud/operator-proxy/api/allinone"
	chainpb "github.com/cita-cloud/operator-proxy/api/chain"
	nodepb "github.com/cita-cloud/operator-proxy/api/node"
)

type printer interface {
	// chain
	InitChain(response *chainpb.ChainSimpleResponse)
	ListChain(list *chainpb.ChainList)

	// node
	InitNode(response *nodepb.NodeSimpleResponse)
	ListNode(list *nodepb.NodeList)
	StartNode(response *nodepb.NodeSimpleResponse)

	// account
	CreateAccount(account *accountpb.Account)
	ListAccount(list *accountpb.AccountList)

	// all in one
	CreateAllInOne(response *allinonepb.AllInOneCreateResponse)
}

func NewPrinter(printerType string) printer {
	switch printerType {
	case "simple":
		return &simplePrinter{}
	case "table":
		return &tablePrinter{printer: &simplePrinter{}}
	}
	return nil
}

func makeAccountListTable(list *accountpb.AccountList) (header []string, rows [][]string) {
	header = []string{"Name", "Namespace", "Chain", "Role", "Domain"}
	for _, account := range list.Accounts {
		rows = append(rows, []string{
			account.Name,
			account.Namespace,
			account.Chain,
			accountpb.Role_name[int32(account.Role)],
			account.Domain,
		})
	}
	return header, rows
}

func makeChainListTable(list *chainpb.ChainList) (header []string, rows [][]string) {
	header = []string{"Name", "Namespace", "Status"}
	for _, chain := range list.Chains {
		rows = append(rows, []string{
			chain.Name,
			chain.Namespace,
			chain.Status.String(),
		})
	}
	return header, rows
}

func makeNodeListTable(list *nodepb.NodeList) (header []string, rows [][]string) {
	header = []string{"Name", "Namespace", "Status"}
	for _, node := range list.Nodes {
		rows = append(rows, []string{
			node.Name,
			node.Namespace,
			"",
		})
	}
	return header, rows
}
