package command

import (
	"os"

	accountpb "github.com/cita-cloud/operator-proxy/api/account"
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
