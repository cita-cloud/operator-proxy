package command

import (
	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	pb "github.com/cita-cloud/operator-proxy/api/citacloud"
)

type printer interface {
	InitChain(chainSimple *pb.ChainConfigSimple)

	CreateAccount(account *accountpb.Account)
	ListAccount(list *accountpb.AccountList)
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
	for _, account := range list.Account {
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
