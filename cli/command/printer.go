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
	accountpb "github.com/cita-cloud/operator-proxy/api/account"
	allinonepb "github.com/cita-cloud/operator-proxy/api/allinone"
	chainpb "github.com/cita-cloud/operator-proxy/api/chain"
	nodepb "github.com/cita-cloud/operator-proxy/api/node"
	"k8s.io/apimachinery/pkg/api/resource"
)

type printer interface {
	// chain
	InitChain(response *chainpb.ChainSimpleResponse)
	OnlineChain(response *chainpb.ChainSimpleResponse)
	ListChain(list *chainpb.ChainList)
	DescribeChain(response *chainpb.ChainDescribeResponse)
	DeleteChain(request *chainpb.ChainDeleteRequest)

	// node
	InitNode(response *nodepb.NodeSimpleResponse)
	ListNode(list *nodepb.NodeList)
	StartNode(response *nodepb.NodeSimpleResponse)
	StopNode(request *nodepb.NodeStopRequest)
	ReloadConfig(request *nodepb.ReloadConfigRequest)

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

func makeNodeListTable(list []*nodepb.Node) (header []string, rows [][]string) {
	header = []string{"Name", "Namespace", "Chain", "Account", "Size", "Status"}
	for _, node := range list {
		rows = append(rows, []string{
			node.Name,
			node.Namespace,
			node.GetChain(),
			node.GetAccount(),
			resource.NewQuantity(node.GetStorageSize(), resource.BinarySI).String(),
			node.GetStatus().String(),
		})
	}
	return header, rows
}
