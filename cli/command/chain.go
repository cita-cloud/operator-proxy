package command

import (
	"errors"
	"time"

	pb "github.com/cita-cloud/operator-proxy/api/chain"
	"github.com/spf13/cobra"
)

var (
	IllegalEnumValue = errors.New("illegal enum value")
	initChainRequest = pb.Chain{}

	consensusType    string
	listChainRequest = pb.ListChainRequest{}
)

func NewChainCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "chain <subcommand>",
		Short: "Chain related commands",
	}

	cc.AddCommand(NewChainInitCommand())
	cc.AddCommand(NewChainListCommand())

	return cc
}

func NewChainInitCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "init <chain name>",
		Short: "Initialize a chain into the k8s cluster",

		Run: chainInitCommandFunc,
	}
	cc.Flags().StringVarP(&initChainRequest.Namespace, "namespace", "n", "cita", "The namespace that your chain create in k8s.")
	cc.Flags().StringVarP(&initChainRequest.Id, "id", "i", "", "The chain id you create.")
	cc.Flags().Int64VarP(&initChainRequest.Timestamp, "timestamp", "t", time.Now().UnixMicro(), "The chain timestamp you create.")
	cc.Flags().StringVarP(&initChainRequest.PrevHash, "prevHash", "p", "", "The chain prevHash you create.")
	cc.Flags().Int32VarP(&initChainRequest.BlockInterval, "blockInterval", "", 3, "The chain blockInterval you create.")
	cc.Flags().Int32VarP(&initChainRequest.BlockLimit, "blockLimit", "", 100, "The chain blockLimit you create.")
	cc.Flags().BoolVarP(&initChainRequest.EnableTls, "enableTls", "", false, "enable tls")
	cc.Flags().StringVarP(&consensusType, "consensusType", "", "Raft", "The chain consensus type you create.")

	cc.Flags().StringVarP(&initChainRequest.NetworkImage, "networkImage", "", "", "The chain's network image.")
	cc.Flags().StringVarP(&initChainRequest.ConsensusImage, "consensusImage", "", "", "The chain's consensus image.")
	cc.Flags().StringVarP(&initChainRequest.ExecutorImage, "executorImage", "", "", "The chain's executor image.")
	cc.Flags().StringVarP(&initChainRequest.StorageImage, "storageImage", "", "", "The chain's storage image.")
	cc.Flags().StringVarP(&initChainRequest.ControllerImage, "controllerImage", "", "", "The chain's controller image.")
	cc.Flags().StringVarP(&initChainRequest.KmsImage, "kmsImage", "", "", "The chain's kms image.")

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
	switch consensusType {
	case pb.ConsensusType_BFT.String():
		initChainRequest.ConsensusType = pb.ConsensusType_BFT
	case pb.ConsensusType_Raft.String():
		initChainRequest.ConsensusType = pb.ConsensusType_Raft
	default:
		ExitWithError(ExitBadArgs, IllegalEnumValue)
	}

	resp, err := cli.ChainInterface.Init(ctx, &initChainRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}
	// print resp info
	display.InitChain(resp)
}

func NewChainListCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "list [options]",
		Short: "List chain in the k8s cluster",

		Run: chainListCommandFunc,
	}

	cc.Flags().StringVarP(&listChainRequest.Namespace, "namespace", "n", "cita", "The namespace that your chain in k8s.")

	return cc
}

func chainListCommandFunc(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	resp, err := cli.ChainInterface.List(ctx, &listChainRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}

	display.ListChain(resp)
}
