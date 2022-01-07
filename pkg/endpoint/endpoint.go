package endpoint

import (
	"context"
	"net"
	"net/url"
	"strings"
)

// ParseEndpoint endpoint parses an endpoint of the form
// (http|https)://<host>*|(unix|unixs)://<path>)
// and returns a protocol ('tcp' or 'unix'),
// host (or filepath if a unix socket),
// scheme (http, https, unix, unixs).
func ParseEndpoint(endpoint string) (proto string, host string, scheme string) {
	proto = "tcp"
	host = endpoint
	url, uerr := url.Parse(endpoint)
	if uerr != nil || !strings.Contains(endpoint, "://") {
		return proto, host, scheme
	}
	scheme = url.Scheme

	// strip scheme:// prefix since grpc dials by host
	host = url.Host
	switch url.Scheme {
	case "http", "https":
	case "unix", "unixs":
		proto = "unix"
		host = url.Host + url.Path
	default:
		proto, host = "", ""
	}
	return proto, host, scheme
}

// Dialer dials a endpoint using net.Dialer.
// Context cancelation and timeout are supported.
func Dialer(ctx context.Context, dialEp string) (net.Conn, error) {
	proto, host, _ := ParseEndpoint(dialEp)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	dialer := &net.Dialer{}
	if deadline, ok := ctx.Deadline(); ok {
		dialer.Deadline = deadline
	}
	return dialer.DialContext(ctx, proto, host)
}