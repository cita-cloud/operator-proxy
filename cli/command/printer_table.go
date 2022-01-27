package command

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	chainpb "github.com/cita-cloud/operator-proxy/api/chain"
	nodepb "github.com/cita-cloud/operator-proxy/api/node"
)

type tablePrinter struct{ printer }

func (t *tablePrinter) ListAccount(list *accountpb.AccountList) {
	header, rows := makeAccountListTable(list)
	t.printTable(header, rows)
}

func (t *tablePrinter) ListChain(list *chainpb.ChainList) {
	header, rows := makeChainListTable(list)
	t.printTable(header, rows)
}

func (t *tablePrinter) ListNode(list *nodepb.NodeList) {
	header, rows := makeNodeListTable(list.GetNodes())
	t.printTable(header, rows)
}

func (t *tablePrinter) DescribeChain(response *chainpb.ChainDescribeResponse) {
	fmt.Println("Chain Base Info:")
	data := [][]string{
		{"Name", response.GetName()},
		{"Namespace", response.GetNamespace()},
		{"Id", response.GetId()},
		{"Timestamp", strconv.FormatInt(response.GetTimestamp(), 10)},
		{"PrevHash", response.GetPrevHash()},
		{"BlockInterval", strconv.FormatInt(int64(response.GetBlockInterval()), 10)},
		{"BlockLimit", strconv.FormatInt(int64(response.GetBlockLimit()), 10)},
		{"EnableTls", strconv.FormatBool(response.GetEnableTls())},
		{"ConsensusType", response.GetConsensusType().String()},
		{"NetworkImage", response.GetNetworkImage()},
		{"ConsensusImage", response.GetConsensusImage()},
		{"ExecutorImage", response.GetExecutorImage()},
		{"StorageImage", response.GetStorageImage()},
		{"ControllerImage", response.GetControllerImage()},
		{"KmsImage", response.GetKmsImage()},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()

	fmt.Println("Node Info:")
	header, rows := makeNodeListTable(response.Nodes)
	t.printTable(header, rows)
}

func (t *tablePrinter) printTable(header []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, row := range rows {
		table.Append(row)
	}
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render()
}
