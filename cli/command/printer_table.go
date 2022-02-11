/*
 * Copyright Rivtower Technologies LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
		{"Status", response.GetStatus().String()},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()

	fmt.Println("Admin Account:")
	al := make([]*accountpb.Account, 0)
	if response.AdminAccount != nil {
		al = append(al, response.AdminAccount)
	}

	accountList := &accountpb.AccountList{Accounts: al}
	header, rows := makeAccountListTable(accountList)
	t.printTable(header, rows)

	fmt.Println("Node Info:")
	header, rows = makeNodeListTable(response.Nodes)
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
