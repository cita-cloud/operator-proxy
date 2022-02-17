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
	"os"
	"time"

	"github.com/spf13/cobra"
)

const (
	defaultDialTimeout      = 2 * time.Minute
	defaultCommandTimeOut   = 5 * time.Minute
	defaultKeepAliveTime    = 2 * time.Minute
	defaultKeepAliveTimeOut = 6 * time.Minute
)

var (
	globalFlags = GlobalFlags{}
)

var RootCmd = &cobra.Command{
	Use:   "cco-cli",
	Short: "The cita-cloud operator command line interface.",
	Long:  `The cita-cloud operator command line interface lets you create and manage CITA-CLOUD chain.`,
}

func init() {
	//cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&globalFlags.Endpoint, "endpoint", getEndPoint(), "gRPC server endpoints")

	RootCmd.PersistentFlags().StringVarP(&globalFlags.OutputFormat, "write-out", "w", "table", "set the output format (simple, table)")

	RootCmd.PersistentFlags().DurationVar(&globalFlags.DialTimeout, "dial-timeout", defaultDialTimeout, "dial timeout for client connections")
	RootCmd.PersistentFlags().DurationVar(&globalFlags.CommandTimeOut, "command-timeout", defaultCommandTimeOut, "timeout for short running command (excluding dial timeout)")
	RootCmd.PersistentFlags().DurationVar(&globalFlags.KeepAliveTime, "keepalive-time", defaultKeepAliveTime, "keepalive time for client connections")
	RootCmd.PersistentFlags().DurationVar(&globalFlags.KeepAliveTimeout, "keepalive-timeout", defaultKeepAliveTimeOut, "keepalive timeout for client connections")

	// add sub command here
	RootCmd.AddCommand(
		NewChainCommand(),
		NewNodeCommand(),
		NewAccountCommand(),
		NewAllInOneCommand(),
		NewVersionCommand())
}

func getEndPoint() string {
	endpoint := os.Getenv("OPERATOR_PROXY_ENDPOINT")
	if endpoint != "" {
		return endpoint
	} else {
		return "127.0.0.1:8090"
	}
}
