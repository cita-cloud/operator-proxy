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
	"errors"
	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
	"time"

	"github.com/spf13/cobra"

	pb "github.com/cita-cloud/operator-proxy/api/chain"
)

var (
	IllegalEnumValue = errors.New("illegal enum value")
	initChainRequest = pb.Chain{}

	consensusType        string
	onlineChainRequest   = pb.ChainOnlineRequest{}
	listChainRequest     = pb.ListChainRequest{}
	describeChainRequest = pb.ChainDescribeRequest{}
	deleteChainRequest   = pb.ChainDeleteRequest{}
)

func NewChainCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "chain <subcommand>",
		Short: "Chain related commands",
	}

	cc.AddCommand(NewChainInitCommand())
	cc.AddCommand(NewChainOnlineCommand())
	cc.AddCommand(NewChainListCommand())
	cc.AddCommand(NewChainDescribeCommand())
	cc.AddCommand(NewChainDeleteCommand())

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
	cc.Flags().StringVarP(&initChainRequest.Version, "version", "v", citacloudv1.LATEST_VERSION, "The chain's main version.")
	cc.Flags().StringVarP(&initChainRequest.PullPolicy, "pullPolicy", "", string(corev1.PullIfNotPresent), "The pull image policy (IfNotPresent/Always).")

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

func NewChainOnlineCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "online <chain name>",
		Short: "Online a chain into the k8s cluster",

		Run: chainOnlineCommandFunc,
	}

	cc.Flags().StringVarP(&onlineChainRequest.Namespace, "namespace", "n", "cita", "The namespace that your chain in k8s.")

	return cc
}

func chainOnlineCommandFunc(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	onlineChainRequest.Name = args[0]

	resp, err := cli.ChainInterface.Online(ctx, &onlineChainRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}

	display.OnlineChain(resp)
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

func NewChainDescribeCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "describe [options]",
		Short: "Show chain detail in the k8s cluster",

		Run: chainDescribeCommandFunc,
	}

	cc.Flags().StringVarP(&describeChainRequest.Namespace, "namespace", "n", "cita", "The namespace that your chain in k8s.")

	return cc
}

func chainDescribeCommandFunc(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	describeChainRequest.Name = args[0]

	resp, err := cli.ChainInterface.Describe(ctx, &describeChainRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}

	display.DescribeChain(resp)
}

func NewChainDeleteCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "delete [options]",
		Short: "Delete a chain in the k8s cluster",

		Run: chainDeleteCommandFunc,
	}

	cc.Flags().StringVarP(&deleteChainRequest.Namespace, "namespace", "n", "cita", "The namespace that your chain in k8s.")

	return cc
}

func chainDeleteCommandFunc(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	deleteChainRequest.Name = args[0]

	_, err := cli.ChainInterface.Delete(ctx, &deleteChainRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}

	display.DeleteChain(&deleteChainRequest)
}
