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
	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"
	pb "github.com/cita-cloud/operator-proxy/api/account"
	"github.com/spf13/cobra"
)

var (
	createAccountRequest = pb.Account{}
	role                 = ""

	listAccountRequest = pb.ListAccountRequest{}
)

func NewAccountCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "account <subcommand>",
		Short: "Account related commands",
	}

	cc.AddCommand(NewAccountCreateCommand())
	cc.AddCommand(NewAccountListCommand())

	return cc
}

func NewAccountCreateCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "create <account name>",
		Short: "Create a node account for chain",

		Run: accountCreateCommandFunc,
	}

	cc.Flags().StringVarP(&createAccountRequest.Namespace, "namespace", "n", "cita", "The namespace that your node account create in k8s.")
	cc.Flags().StringVarP(&createAccountRequest.Chain, "chain", "c", "", "The chain name corresponding to the node account.")
	cc.Flags().StringVarP(&createAccountRequest.KmsPassword, "kmsPassword", "k", "", "The account kms password.")
	cc.Flags().StringVarP(&role, "role", "r", "", "The role of node account.")
	cc.Flags().StringVarP(&createAccountRequest.Domain, "domain", "d", "", "The domain of node account.")

	return cc
}

func convertRole(role string) {
	if role == string(citacloudv1.Admin) {
		createAccountRequest.Role = pb.Role_Admin
	} else if role == string(citacloudv1.Consensus) {
		createAccountRequest.Role = pb.Role_Consensus
	} else if role == string(citacloudv1.Ordinary) {
		createAccountRequest.Role = pb.Role_Ordinary
	}
}

func accountCreateCommandFunc(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	createAccountRequest.Name = args[0]
	convertRole(role)
	resp, err := cli.AccountInterface.CreateAccount(ctx, &createAccountRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}
	// print resp info
	display.CreateAccount(resp)
}

func NewAccountListCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "list [options]",
		Short: "List node account in the k8s cluster",

		Run: accountListCommandFunc,
	}

	cc.Flags().StringVarP(&listAccountRequest.Namespace, "namespace", "n", "cita", "The namespace that your node account in k8s.")
	cc.Flags().StringVarP(&listAccountRequest.Chain, "chain", "c", "", "The chain name corresponding to the node account.")

	return cc
}

func accountListCommandFunc(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	resp, err := cli.AccountInterface.ListAccount(ctx, &listAccountRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}

	display.ListAccount(resp)
}
