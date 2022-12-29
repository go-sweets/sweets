package mrpc

import (
	"github.com/mix-plus/go-mixplus/mrpc/internal"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	WithDialOption = internal.WithDialOption

	WithNonBlock = internal.WithNonBlock

	WithStreamClientInterceptor = internal.WithStreamClientInterceptor

	WithTimeout = internal.WithTimeout

	WithTransportCredentials = internal.WithTransportCredentials

	WithUnaryClientInterceptor = internal.WithUnaryClientInterceptor
)

type (
	Client internal.Client

	ClientOption = internal.ClientOption

	RpcClient struct {
		client Client
	}
)

func MustNewClient(c RpcClientConf, options ...ClientOption) Client {
	cli, err := NewClient(c, options...)
	if err != nil {
		log.Fatal(err)
	}

	return cli
}

func NewClient(c RpcClientConf, options ...ClientOption) (Client, error) {
	var opts []ClientOption

	if c.NonBlock {
		opts = append(opts, WithNonBlock())
	}

	if c.Timeout > 0 {
		opts = append(opts, WithTimeout(time.Duration(c.Timeout)*time.Millisecond))
	}

	opts = append(opts, options...)

	client, err := internal.NewClient(c.Target, opts...)
	if err != nil {
		return nil, err
	}

	return &RpcClient{
		client: client,
	}, nil
}

func NewClientWithTarget(target string, opts ...ClientOption) (Client, error) {
	return internal.NewClient(target, opts...)
}

func (rc *RpcClient) Conn() *grpc.ClientConn {
	return rc.client.Conn()
}

func SetClientSlowThreshold(threshold time.Duration) {
	// TODO
}
