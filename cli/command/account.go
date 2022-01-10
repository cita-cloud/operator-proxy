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
		Short: "account related commands",
	}

	cc.AddCommand(NewAccountCreateCommand())
	cc.AddCommand(NewAccountListCommand())

	return cc
}

func NewAccountCreateCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a node account for chain",

		Run: accountCreateCommandFunc,
	}

	cc.Flags().StringVarP(&createAccountRequest.Namespace, "namespace", "n", "cita", "The namespace that your node account create in k8s.")
	cc.Flags().StringVarP(&createAccountRequest.Chain, "chain", "c", "", "The chain name corresponding to the node account.")
	cc.Flags().StringVarP(&createAccountRequest.KmsPassword, "kmsPassword", "k", "", "The account kms password.")
	cc.Flags().StringVarP(&role, "role", "r", "", "The node account kms password.")
	convertRole(role)
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
