package client

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/cita-cloud/operator-proxy/pkg/endpoint"
	grpccredentials "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"

	"google.golang.org/grpc"
)

var (
	ErrNoAvailableEndpoint = errors.New("cita-cloud-ctl: no available endpoint")
)

type Config struct {
	// Endpoints is a list of URLs.
	Endpoint string `json:"endpoint"`

	// AutoSyncInterval is the interval to update endpoints with its latest members.
	// 0 disables auto-sync. By default auto-sync is disabled.
	AutoSyncInterval time.Duration `json:"auto-sync-interval"`

	// DialTimeout is the timeout for failing to establish a connection.
	DialTimeout time.Duration `json:"dial-timeout"`

	// DialKeepAliveTime is the time after which client pings the server to see if
	// transport is alive.
	DialKeepAliveTime time.Duration `json:"dial-keep-alive-time"`

	// DialKeepAliveTimeout is the time that the client waits for a response for the
	// keep-alive probe. If the response is not received in this time, the connection is closed.
	DialKeepAliveTimeout time.Duration `json:"dial-keep-alive-timeout"`

	// MaxCallSendMsgSize is the client-side request send limit in bytes.
	// If 0, it defaults to 2.0 MiB (2 * 1024 * 1024).
	// Make sure that "MaxCallSendMsgSize" < server-side default send/recv limit.
	// ("--max-request-bytes" flag to etcd or "embed.Config.MaxRequestBytes").
	MaxCallSendMsgSize int

	// MaxCallRecvMsgSize is the client-side response receive limit.
	// If 0, it defaults to "math.MaxInt32", because range response can
	// easily exceed request send limits.
	// Make sure that "MaxCallRecvMsgSize" >= server-side default send/recv limit.
	// ("--max-request-bytes" flag to etcd or "embed.Config.MaxRequestBytes").
	MaxCallRecvMsgSize int

	// TLS holds the client secure credentials, if any.
	TLS *tls.Config

	// DialOptions is a list of dial options for the grpc client (e.g., for interceptors).
	// For example, pass "grpc.WithBlock()" to block until the underlying connection is up.
	// Without this, Dial returns immediately and connecting the server happens in background.
	DialOptions []grpc.DialOption

	// Context is the default client context; it can be used to cancel grpc dial out and
	// other operations that do not have an explicit context.
	Context context.Context

	// PermitWithoutStream when set will allow client to send keepalive pings to server without any active streams(RPCs).
	PermitWithoutStream bool `json:"permit-without-stream"`
}

type Client struct {
	ChainInterface
	AccountInterface
	AllInOneInterface

	conn *grpc.ClientConn

	cfg   Config
	creds grpccredentials.TransportCredentials
	mu    *sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc

	callOpts []grpc.CallOption
}

func New(cfg Config) (*Client, error) {
	if cfg.Endpoint == "" {
		return nil, ErrNoAvailableEndpoint
	}

	return newClient(&cfg)
}

func newClient(cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = &Config{}
	}
	var creds grpccredentials.TransportCredentials
	//if cfg.TLS != nil {
	//	creds = credentials.NewBundle(credentials.Config{TLSConfig: cfg.TLS}).TransportCredentials()
	//}

	baseCtx := context.TODO()
	if cfg.Context != nil {
		baseCtx = cfg.Context
	}

	ctx, cancel := context.WithCancel(baseCtx)
	client := &Client{
		conn:     nil,
		cfg:      *cfg,
		creds:    creds,
		ctx:      ctx,
		cancel:   cancel,
		mu:       new(sync.RWMutex),
		callOpts: defaultCallOpts,
	}

	if cfg.MaxCallSendMsgSize > 0 || cfg.MaxCallRecvMsgSize > 0 {
		if cfg.MaxCallRecvMsgSize > 0 && cfg.MaxCallSendMsgSize > cfg.MaxCallRecvMsgSize {
			return nil, fmt.Errorf("gRPC message recv limit (%d bytes) must be greater than send limit (%d bytes)", cfg.MaxCallRecvMsgSize, cfg.MaxCallSendMsgSize)
		}
		callOpts := []grpc.CallOption{
			defaultFailFast,
			defaultMaxCallSendMsgSize,
			defaultMaxCallRecvMsgSize,
		}
		if cfg.MaxCallSendMsgSize > 0 {
			callOpts[1] = grpc.MaxCallSendMsgSize(cfg.MaxCallSendMsgSize)
		}
		if cfg.MaxCallRecvMsgSize > 0 {
			callOpts[2] = grpc.MaxCallRecvMsgSize(cfg.MaxCallRecvMsgSize)
		}
		client.callOpts = callOpts
	}

	// Use a provided endpoint target so that for https:// without any tls config given, then
	// grpc will assume the certificate server name is the endpoint host.
	conn, err := client.Dial(cfg.Endpoint)
	if err != nil {
		client.cancel()
		return nil, err
	}
	client.conn = conn

	// New grpc service client
	client.ChainInterface = NewChain(client)
	client.AccountInterface = NewAccount(client)
	client.AllInOneInterface = NewAllInOne(client)

	return client, nil
}

// Dial connects to a single endpoint using the client's config.
func (c *Client) Dial(ep string) (*grpc.ClientConn, error) {
	creds, err := c.directDialCreds(ep)
	if err != nil {
		return nil, err
	}
	// Use the grpc passthrough resolver to directly dial a single endpoint.
	// This resolver passes through the 'unix' and 'unixs' endpoints schemes used
	// by etcd without modification, allowing us to directly dial endpoints and
	// using the same dial functions that we use for load balancer dialing.
	return c.dial(fmt.Sprintf("passthrough:///%s", ep), creds)
}

func (c *Client) processCreds(scheme string) (creds grpccredentials.TransportCredentials) {
	creds = c.creds
	switch scheme {
	case "unix":
	case "http":
		creds = nil
	case "https", "unixs":
		if creds != nil {
			break
		}
		//creds = credentials.NewBundle(credentials.Config{}).TransportCredentials()
	default:
		creds = nil
	}
	return creds
}

func (c *Client) directDialCreds(ep string) (grpccredentials.TransportCredentials, error) {
	_, host, scheme := endpoint.ParseEndpoint(ep)
	creds := c.creds
	if len(scheme) != 0 {
		creds = c.processCreds(scheme)
		if creds != nil {
			clone := creds.Clone()
			// Set the server name must to the endpoint hostname without port since grpc
			// otherwise attempts to check if x509 cert is valid for the full endpoint
			// including the scheme and port, which fails.
			overrideServerName, _, err := net.SplitHostPort(host)
			if err != nil {
				// Either the host didn't have a port or the host could not be parsed. Either way, continue with the
				// original host string.
				overrideServerName = host
			}
			clone.OverrideServerName(overrideServerName)
			creds = clone
		}
	}
	return creds, nil
}

// dial configures and dials any grpc balancer target.
func (c *Client) dial(target string, creds grpccredentials.TransportCredentials, dopts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts, err := c.dialSetupOpts(creds, dopts...)
	if err != nil {
		return nil, fmt.Errorf("failed to configure dialer: %v", err)
	}

	opts = append(opts, c.cfg.DialOptions...)

	dctx := c.ctx
	if c.cfg.DialTimeout > 0 {
		var cancel context.CancelFunc
		dctx, cancel = context.WithTimeout(c.ctx, c.cfg.DialTimeout)
		defer cancel() // TODO: Is this right for cases where grpc.WithBlock() is not set on the dial options?
	}

	conn, err := grpc.DialContext(dctx, target, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// dialSetupOpts gives the dial opts prior to any authentication.
func (c *Client) dialSetupOpts(creds grpccredentials.TransportCredentials, dopts ...grpc.DialOption) (opts []grpc.DialOption, err error) {
	if c.cfg.DialKeepAliveTime > 0 {
		params := keepalive.ClientParameters{
			Time:                c.cfg.DialKeepAliveTime,
			Timeout:             c.cfg.DialKeepAliveTimeout,
			PermitWithoutStream: c.cfg.PermitWithoutStream,
		}
		opts = append(opts, grpc.WithKeepaliveParams(params))
	}
	opts = append(opts, dopts...)

	dialer := endpoint.Dialer
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithContextDialer(dialer))

	return opts, nil
}
