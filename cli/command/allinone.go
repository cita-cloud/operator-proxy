package command

import (
	"time"

	pb "github.com/cita-cloud/operator-proxy/api/allinone"
	chainpb "github.com/cita-cloud/operator-proxy/api/chain"
	"github.com/spf13/cobra"
)

var (
	allInOneCreateRequest = pb.AllInOneCreateRequest{}
)

func NewAllInOneCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "all-in-one <subcommand>",
		Short: "Create a chain with one click",
	}

	cc.AddCommand(NewAllInOneCreateCommand())

	return cc
}

func NewAllInOneCreateCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "create <chain name>",
		Short: "Create a node account for chain",

		Run: allInOneCreateCommandFunc,
	}

	cc.Flags().StringVarP(&allInOneCreateRequest.Namespace, "namespace", "n", "cita", "The namespace that your chain create in k8s.")
	cc.Flags().StringVarP(&allInOneCreateRequest.Id, "id", "i", "", "The chain id you create.")
	cc.Flags().Int64VarP(&allInOneCreateRequest.Timestamp, "timestamp", "t", time.Now().UnixMicro(), "The chain timestamp you create.")
	cc.Flags().StringVarP(&allInOneCreateRequest.PrevHash, "prevHash", "p", "", "The chain prevHash you create.")
	cc.Flags().Int32VarP(&allInOneCreateRequest.BlockInterval, "blockInterval", "", 3, "The chain blockInterval you create.")
	cc.Flags().Int32VarP(&allInOneCreateRequest.BlockLimit, "blockLimit", "", 100, "The chain blockLimit you create.")
	cc.Flags().BoolVarP(&allInOneCreateRequest.EnableTls, "enableTls", "", false, "enable tls")
	cc.Flags().StringVarP(&consensusType, "consensusType", "", "Raft", "The chain consensus type you create.")

	cc.Flags().StringVarP(&allInOneCreateRequest.NetworkImage, "networkImage", "", "", "The chain's network image.")
	cc.Flags().StringVarP(&allInOneCreateRequest.ConsensusImage, "consensusImage", "", "", "The chain's consensus image.")
	cc.Flags().StringVarP(&allInOneCreateRequest.ExecutorImage, "executorImage", "", "", "The chain's executor image.")
	cc.Flags().StringVarP(&allInOneCreateRequest.StorageImage, "storageImage", "", "", "The chain's storage image.")
	cc.Flags().StringVarP(&allInOneCreateRequest.ControllerImage, "controllerImage", "", "", "The chain's controller image.")
	cc.Flags().StringVarP(&allInOneCreateRequest.KmsImage, "kmsImage", "", "", "The chain's kms image.")

	cc.Flags().StringVarP(&allInOneCreateRequest.StorageClassName, "storageClassName", "", "", "The node's storage class.")
	cc.Flags().Int64VarP(&allInOneCreateRequest.StorageSize, "storageSize", "", 10737418240, "The chain's storage size.")
	cc.Flags().StringVarP(&allInOneCreateRequest.LogLevel, "logLevel", "", "info", "The node's log level.")

	cc.Flags().Int32VarP(&allInOneCreateRequest.NodeCount, "nodeCount", "", 3, "The node count for chain start.")

	return cc
}

func allInOneCreateCommandFunc(cmd *cobra.Command, args []string) {
	//var err error

	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	allInOneCreateRequest.Name = args[0]
	switch consensusType {
	case chainpb.ConsensusType_BFT.String():
		allInOneCreateRequest.ConsensusType = chainpb.ConsensusType_BFT
	case chainpb.ConsensusType_Raft.String():
		allInOneCreateRequest.ConsensusType = chainpb.ConsensusType_Raft
	default:
		ExitWithError(ExitBadArgs, IllegalEnumValue)
	}

	resp, err := cli.AllInOneInterface.Create(ctx, &allInOneCreateRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}
	// print resp info
	display.CreateAllInOne(resp)
}