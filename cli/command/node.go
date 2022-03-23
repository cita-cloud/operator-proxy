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
	pb "github.com/cita-cloud/operator-proxy/api/node"
	"github.com/spf13/cobra"
)

var (
	initNodeRequest = pb.Node{}

	listNodeRequest     = pb.ListNodeRequest{}
	startNodeRequest    = pb.NodeStartRequest{}
	stopNodeRequest     = pb.NodeStopRequest{}
	reloadConfigRequest = pb.ReloadConfigRequest{}
)

func NewNodeCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "node <subcommand>",
		Short: "Node related commands",
	}

	cc.AddCommand(NewNodeInitCommand())
	cc.AddCommand(NewNodeListCommand())
	cc.AddCommand(NewNodeStartCommand())
	cc.AddCommand(NewNodeStopCommand())
	cc.AddCommand(NewNodeReloadConfigCommand())

	return cc
}

func NewNodeInitCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "init <node name>",
		Short: "Init a node for chain",

		Run: nodeInitCommandFunc,
	}

	cc.Flags().StringVarP(&initNodeRequest.Namespace, "namespace", "n", "cita", "The namespace that your node create in k8s.")
	cc.Flags().StringVarP(&initNodeRequest.Cluster, "cluster", "", "", "The name of the k8s cluster where the node is located.")
	cc.Flags().StringVarP(&initNodeRequest.Chain, "chain", "", "", "The chain name corresponding to the node.")
	cc.Flags().StringVarP(&initNodeRequest.Account, "account", "a", "", "The account name corresponding to the node.")
	cc.Flags().StringVarP(&initNodeRequest.ExternalIp, "externalIp", "", "", "The external ip exposed by node.")
	cc.Flags().Int32VarP(&initNodeRequest.Port, "port", "", 9999, "The external port exposed by node.")
	cc.Flags().StringVarP(&initNodeRequest.StorageClassName, "storageClassName", "", "nas-client-provisioner", "The node's storage class name.")
	cc.Flags().Int64VarP(&initNodeRequest.StorageSize, "storageSize", "", 10737418240, "The node's storage size.")
	cc.Flags().StringVarP(&initNodeRequest.LogLevel, "logLevel", "", "info", "The node's log level.")

	return cc
}

func nodeInitCommandFunc(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	initNodeRequest.Name = args[0]

	resp, err := cli.NodeInterface.Init(ctx, &initNodeRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}

	display.InitNode(resp)
}

func NewNodeListCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "list [options]",
		Short: "List node in the k8s cluster",

		Run: nodeListCommandFunc,
	}

	cc.Flags().StringVarP(&listNodeRequest.Namespace, "namespace", "n", "cita", "The namespace that your node in k8s.")
	cc.Flags().StringVarP(&listNodeRequest.Chain, "chain", "c", "", "The chain name corresponding to the node.")

	return cc
}

func nodeListCommandFunc(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	resp, err := cli.NodeInterface.List(ctx, &listNodeRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}

	display.ListNode(resp)
}

func NewNodeStartCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "start <node name>",
		Short: "Start a node",

		Run: nodeStartCommandFunc,
	}

	cc.Flags().StringVarP(&startNodeRequest.Namespace, "namespace", "n", "cita", "The namespace that your node in k8s.")

	return cc
}

func nodeStartCommandFunc(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	startNodeRequest.Name = args[0]

	resp, err := cli.NodeInterface.Start(ctx, &startNodeRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}

	display.StartNode(resp)
}

func NewNodeStopCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "stop <node name>",
		Short: "Stop a node",

		Run: nodeStopCommandFunc,
	}

	cc.Flags().StringVarP(&stopNodeRequest.Namespace, "namespace", "n", "cita", "The namespace that your node in k8s.")

	return cc
}

func nodeStopCommandFunc(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	stopNodeRequest.Name = args[0]

	_, err := cli.NodeInterface.Stop(ctx, &stopNodeRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}

	display.StopNode(&stopNodeRequest)
}

func NewNodeReloadConfigCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "reload <node name>",
		Short: "Reload the node's config, usually used to add or delete nodes in a chain",

		Run: nodeReloadConfigCommand,
	}

	cc.Flags().StringVarP(&reloadConfigRequest.Namespace, "namespace", "n", "cita", "The namespace that your node in k8s.")

	return cc
}

func nodeReloadConfigCommand(cmd *cobra.Command, args []string) {
	// create grpc client
	ctx, cancel := commandCtx(cmd)
	defer func() {
		cancel()
	}()
	cli := newClientFromCmd(cmd)

	reloadConfigRequest.Name = args[0]

	_, err := cli.NodeInterface.ReloadConfig(ctx, &reloadConfigRequest)
	if err != nil {
		ExitWithError(ExitError, err)
	}

	display.ReloadConfig(&reloadConfigRequest)
}
