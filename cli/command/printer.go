package command

import (
	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	pb "github.com/cita-cloud/operator-proxy/api/citacloud"
)

type printer interface {
	InitChain(chainSimple *pb.ChainConfigSimple)

	CreateAccount(account *accountpb.Account)
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
