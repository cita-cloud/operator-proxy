package command

import (
	"os"

	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	chainpb "github.com/cita-cloud/operator-proxy/api/chain"
	nodepb "github.com/cita-cloud/operator-proxy/api/node"
	"github.com/olekukonko/tablewriter"
)

type tablePrinter struct{ printer }

func (t *tablePrinter) ListAccount(list *accountpb.AccountList) {
	header, rows := makeAccountListTable(list)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, row := range rows {
		table.Append(row)
	}
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.Render()
}

func (t *tablePrinter) ListChain(list *chainpb.ChainList) {
	header, rows := makeChainListTable(list)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, row := range rows {
		table.Append(row)
	}
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.Render()
}

func (t *tablePrinter) ListNode(list *nodepb.NodeList) {
	header, rows := makeNodeListTable(list)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, row := range rows {
		table.Append(row)
	}
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.Render()
}
