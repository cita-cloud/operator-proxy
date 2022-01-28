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
	"time"

	"github.com/cita-cloud/operator-proxy/cli/client"
	"github.com/spf13/cobra"
)

type GlobalFlags struct {
	Endpoint         string
	DialTimeout      time.Duration
	CommandTimeOut   time.Duration
	KeepAliveTime    time.Duration
	KeepAliveTimeout time.Duration

	OutputFormat string
}

var display printer = &simplePrinter{}

func initDisplayFromCmd(cmd *cobra.Command) {
	outputType, err := cmd.Flags().GetString("write-out")
	if err != nil {
		ExitWithError(ExitError, err)
	}
	if display = NewPrinter(outputType); display == nil {
		ExitWithError(ExitBadFeature, errors.New("unsupported output format"))
	}
}

type secureCfg struct {
	cert               string
	key                string
	cacert             string
	insecureSkipVerify bool
}

type clientConfig struct {
	endpoint         string
	dialTimeout      time.Duration
	keepAliveTime    time.Duration
	keepAliveTimeout time.Duration
	scfg             *secureCfg
}

func clientConfigFromCmd(cmd *cobra.Command) *clientConfig {
	cfg := &clientConfig{}
	cfg.endpoint = endpointFromCmd(cmd)
	cfg.dialTimeout = dialTimeoutFromCmd(cmd)
	cfg.keepAliveTime = keepAliveTimeFromCmd(cmd)
	cfg.keepAliveTimeout = keepAliveTimeoutFromCmd(cmd)
	//cfg.scfg = secureCfgFromCmd(cmd)
	// init display printer
	initDisplayFromCmd(cmd)
	return cfg
}

func (cc *clientConfig) newClient() *client.Client {
	cfg, err := newClientCfg(cc.endpoint, cc.dialTimeout, cc.keepAliveTime, cc.keepAliveTimeout, cc.scfg)
	if err != nil {
		ExitWithError(ExitBadArgs, err)
	}

	client, err := client.New(*cfg)
	if err != nil {
		ExitWithError(ExitBadConnection, err)
	}

	return client
}

func newClientCfg(endpoint string, dialTimeout, keepAliveTime, keepAliveTimeout time.Duration, scfg *secureCfg) (*client.Config, error) {
	// set tls if any one tls option set
	//var cfgtls *client.TLSInfo
	//tlsinfo := client.TLSInfo{}
	//if scfg.cert != "" {
	//	tlsinfo.CertFile = scfg.cert
	//	cfgtls = &tlsinfo
	//}
	//
	//if scfg.key != "" {
	//	tlsinfo.KeyFile = scfg.key
	//	cfgtls = &tlsinfo
	//}
	//
	//if scfg.cacert != "" {
	//	tlsinfo.TrustedCAFile = scfg.cacert
	//	cfgtls = &tlsinfo
	//}

	cfg := &client.Config{
		Endpoint:             endpoint,
		DialTimeout:          dialTimeout,
		DialKeepAliveTime:    keepAliveTime,
		DialKeepAliveTimeout: keepAliveTimeout,
	}

	//if cfgtls != nil {
	//	clientTLS, err := cfgtls.ClientConfig()
	//	if err != nil {
	//		return nil, err
	//	}
	//	cfg.TLS = clientTLS
	//}

	// If the user wants to skip TLS verification then we should set
	// the InsecureSkipVerify flag in tls configuration.
	//if scfg.insecureSkipVerify && cfg.TLS != nil {
	//	cfg.TLS.InsecureSkipVerify = true
	//}

	return cfg, nil
}

func newClientFromCmd(cmd *cobra.Command) *client.Client {
	cfg := clientConfigFromCmd(cmd)
	return cfg.newClient()
}

func endpointFromCmd(cmd *cobra.Command) string {
	endpoint, err := cmd.Flags().GetString("endpoint")
	if err != nil {
		ExitWithError(ExitError, err)
	}
	return endpoint
}

func dialTimeoutFromCmd(cmd *cobra.Command) time.Duration {
	dialTimeout, err := cmd.Flags().GetDuration("dial-timeout")
	if err != nil {
		ExitWithError(ExitError, err)
	}
	return dialTimeout
}

func keepAliveTimeFromCmd(cmd *cobra.Command) time.Duration {
	keepAliveTime, err := cmd.Flags().GetDuration("keepalive-time")
	if err != nil {
		ExitWithError(ExitError, err)
	}
	return keepAliveTime
}

func keepAliveTimeoutFromCmd(cmd *cobra.Command) time.Duration {
	keepAliveTimeout, err := cmd.Flags().GetDuration("keepalive-timeout")
	if err != nil {
		ExitWithError(ExitError, err)
	}
	return keepAliveTimeout
}

func secureCfgFromCmd(cmd *cobra.Command) *secureCfg {
	cert, key, cacert := keyAndCertFromCmd(cmd)
	skipVerify := insecureSkipVerifyFromCmd(cmd)

	return &secureCfg{
		cert:               cert,
		key:                key,
		cacert:             cacert,
		insecureSkipVerify: skipVerify,
	}
}

func keyAndCertFromCmd(cmd *cobra.Command) (cert, key, cacert string) {
	var err error
	if cert, err = cmd.Flags().GetString("cert"); err != nil {
		ExitWithError(ExitBadArgs, err)
	}

	if key, err = cmd.Flags().GetString("key"); err != nil {
		ExitWithError(ExitBadArgs, err)
	}

	if cacert, err = cmd.Flags().GetString("cacert"); err != nil {
		ExitWithError(ExitBadArgs, err)
	}

	return cert, key, cacert
}

func insecureSkipVerifyFromCmd(cmd *cobra.Command) bool {
	skipVerify, err := cmd.Flags().GetBool("insecure-skip-tls-verify")
	if err != nil {
		ExitWithError(ExitError, err)
	}
	return skipVerify
}
