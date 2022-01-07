package command

import (
	pb "github.com/cita-cloud/operator-proxy/api/citacloud"
	"github.com/spf13/cobra"
)

var (
	initChainRequest = pb.ChainConfig{}
)

func NewChainCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "chain <subcommand>",
		Short: "chain related commands",
	}

	cc.AddCommand(NewChainInitCommand())

	return cc
}

func NewChainInitCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "init <clusterName>",
		Short: "Initialize a chain config into the k8s cluster",

		Run: chainInitCommandFunc,
	}
	cc.Flags().StringVarP(&initChainRequest.Namespace, "namespace", "", "", "The namespace that your chain create in k8s.")
	cc.Flags().StringVarP(&initChainRequest.Id, "id", "", "", "The chain id you create.")
	cc.Flags().StringVarP(&initChainRequest.Timestamp, "timestamp", "", "", "The chain timestamp you create.")
	cc.Flags().StringVarP(&initChainRequest.Prevhash, "prevhash", "", "", "The chain prevhash you create.")
	cc.Flags().Int32VarP(&initChainRequest.BlockInterval, "blockInterval", "", 3, "The chain blockInterval you create.")
	cc.Flags().Int32VarP(&initChainRequest.BlockLimit, "blockLimit", "", 3, "The chain blockLimit you create.")
	cc.Flags().BoolVarP(&initChainRequest.EnableTls, "enableTls", "", false, "enable tls")
	cc.Flags().StringVarP(&initChainRequest.KmsPassword, "kmsPassword", "", "", "The chain kmsPassword you create.")
	//cc.Flags().StringArrayVarP(&initChainRequest.ConsensusType, "kmsPassword", "", "", "The chain kmsPassword you create.")

	return cc
}

func chainInitCommandFunc(cmd *cobra.Command, args []string) {
	//var err error

	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	initChainRequest.Name = args[0]

	resp, err := cli.ChainInterface.InitChain(ctx, &initChainRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}
	// print resp info
	display.InitChain(resp)
}
