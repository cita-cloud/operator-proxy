package command

import (
	"time"

	"github.com/spf13/cobra"
)

const (
	defaultDialTimeout      = 2 * time.Second
	defaultCommandTimeOut   = 5 * time.Second
	defaultKeepAliveTime    = 2 * time.Second
	defaultKeepAliveTimeOut = 6 * time.Second
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

	RootCmd.PersistentFlags().StringVar(&globalFlags.Endpoint, "endpoint", "127.0.0.1:8090", "gRPC server endpoints")

	RootCmd.PersistentFlags().StringVarP(&globalFlags.OutputFormat, "write-out", "w", "table", "set the output format (simple, table)")

	RootCmd.PersistentFlags().DurationVar(&globalFlags.DialTimeout, "dial-timeout", defaultDialTimeout, "dial timeout for client connections")
	RootCmd.PersistentFlags().DurationVar(&globalFlags.CommandTimeOut, "command-timeout", defaultCommandTimeOut, "timeout for short running command (excluding dial timeout)")
	RootCmd.PersistentFlags().DurationVar(&globalFlags.KeepAliveTime, "keepalive-time", defaultKeepAliveTime, "keepalive time for client connections")
	RootCmd.PersistentFlags().DurationVar(&globalFlags.KeepAliveTimeout, "keepalive-timeout", defaultKeepAliveTimeOut, "keepalive timeout for client connections")

	// add sub command here
	RootCmd.AddCommand(
		NewChainCommand(),
		NewAccountCommand())
}
