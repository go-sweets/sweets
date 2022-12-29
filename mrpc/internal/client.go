package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/mix-plus/go-mixplus/mrpc/internal/balancer/p2c"
	"github.com/mix-plus/go-mixplus/mrpc/internal/clientinterceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"time"
)

const (
	dialTimeout = time.Second * 3
	separator   = '/'
)

func init() {
	// TODO Register
}

type (
	// Client interface wraps the Conn method.
	Client interface {
		Conn() *grpc.ClientConn
	}

	// A ClientOptions is a client options.
	ClientOptions struct {
		NonBlock    bool
		Timeout     time.Duration
		Secure      bool
		DialOptions []grpc.DialOption
	}

	// ClientOption defines the method to customize a ClientOptions.
	ClientOption func(options *ClientOptions)

	client struct {
		conn *grpc.ClientConn
	}
)

func NewClient(target string, opts ...ClientOption) (Client, error) {
	var cli client
	svcCfg := fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, p2c.Name)
	balancerOpt := WithDialOption(grpc.WithDefaultServiceConfig(svcCfg))
	opts = append([]ClientOption{balancerOpt}, opts...)
	if err := cli.dial(target, opts...); err != nil {
		return nil, err
	}

	return &cli, nil
}

func (c *client) Conn() *grpc.ClientConn {
	return c.conn
}

func (c *client) buildDialOptions(opts ...ClientOption) []grpc.DialOption {
	var cliOpts ClientOptions
	for _, opt := range opts {
		opt(&cliOpts)
	}

	var options []grpc.DialOption
	if !cliOpts.Secure {
		options = append([]grpc.DialOption(nil), grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	if !cliOpts.NonBlock {
		options = append(options, grpc.WithBlock())
	}

	options = append(options,
		WithUnaryClientInterceptors(
			clientinterceptors.UnaryTracingInterceptor,
			clientinterceptors.DurationInterceptor,
			clientinterceptors.PrometheusInterceptor,
			clientinterceptors.BreakerInterceptor,
			clientinterceptors.TimeoutInterceptor(cliOpts.Timeout),
		),
		WithStreamClientInterceptors(
			clientinterceptors.StreamTracingInterceptor,
		),
	)

	return append(options, cliOpts.DialOptions...)
}

func (c *client) dial(server string, opts ...ClientOption) error {
	options := c.buildDialOptions(opts...)
	timeCtx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()
	conn, err := grpc.DialContext(timeCtx, server, options...)
	if err != nil {
		service := server
		if errors.Is(err, context.DeadlineExceeded) {
			pos := strings.LastIndexByte(server, separator)
			if pos > 0 && pos < len(server)-1 {
				service = server[pos+1:]
			}
		}
		return fmt.Errorf("rpc: dial: %s, erros: %s, make sure rpc service %q is already started",
			server, err.Error(), service)
	}

	c.conn = conn
	return nil
}

// WithDialOption returns a func to customize a ClientOptions with given dial option.
func WithDialOption(opt grpc.DialOption) ClientOption {
	return func(options *ClientOptions) {
		options.DialOptions = append(options.DialOptions, opt)
	}
}

// WithNonBlock sets the dialing to be nonblock.
func WithNonBlock() ClientOption {
	return func(options *ClientOptions) {
		options.NonBlock = true
	}
}

// WithStreamClientInterceptor returns func to customize a ClientOptions with given interceptor.
func WithStreamClientInterceptor(interceptor grpc.StreamClientInterceptor) ClientOption {
	return func(options *ClientOptions) {
		options.DialOptions = append(options.DialOptions, WithStreamClientInterceptors(interceptor))
	}
}

// WithTimeout returns a func to customize a ClientOptions with given timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(options *ClientOptions) {
		options.Timeout = timeout
	}
}

// WithTransportCredentials return a func to make the gRPC calls secured with given credentials.
func WithTransportCredentials(creds credentials.TransportCredentials) ClientOption {
	return func(options *ClientOptions) {
		options.Secure = true
		options.DialOptions = append(options.DialOptions, grpc.WithTransportCredentials(creds))
	}
}

// WithUnaryClientInterceptor returns a func to customize a ClientOptions with given interceptor.
func WithUnaryClientInterceptor(interceptor grpc.UnaryClientInterceptor) ClientOption {
	return func(options *ClientOptions) {
		options.DialOptions = append(options.DialOptions, WithUnaryClientInterceptors(interceptor))
	}
}
