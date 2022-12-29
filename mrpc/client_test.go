package mrpc

import (
	"context"
	"fmt"
	"github.com/mix-plus/go-mixplus/mrpc/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
	"time"
)

func init() {
	logx.Disable()
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	mock.RegisterDepositServiceServer(server, &mock.DepositServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestDepositServer_Deposit(t *testing.T) {
	tests := []struct {
		name    string
		amount  float32
		res     *mock.DepositResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"invalid request with negative amount",
			-1.11,
			nil,
			codes.InvalidArgument,
			fmt.Sprintf("cannot deposit %v", -1.11),
		},
		{
			"valid request with non negative amount",
			0.00,
			&mock.DepositResponse{Ok: true},
			codes.OK,
			"",
		},
		{
			"valid request with long handling time",
			2000.00,
			nil,
			codes.DeadlineExceeded,
			"context deadline exceeded",
		},
	}

	directClient := MustNewClient(
		RpcClientConf{
			Timeout: 1000,
		},
		WithDialOption(grpc.WithContextDialer(dialer())),
		WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{},
			cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			return invoker(ctx, method, req, reply, cc, opts...)
		}),
	)
	nonBlockClient := MustNewClient(
		RpcClientConf{
			Timeout:  1000,
			NonBlock: true,
		},
		WithDialOption(grpc.WithContextDialer(dialer())),
		WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{},
			cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			return invoker(ctx, method, req, reply, cc, opts...)
		}),
	)
	tarConfClient := MustNewClient(
		RpcClientConf{
			Target:  "localhost:8080",
			Timeout: 1000,
		},
		WithDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
		WithDialOption(grpc.WithContextDialer(dialer())),
		WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{},
			cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			return invoker(ctx, method, req, reply, cc, opts...)
		}),
	)
	targetClient, err := NewClientWithTarget("foo",
		WithDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
		WithDialOption(grpc.WithContextDialer(dialer())), WithUnaryClientInterceptor(
			func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
				invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				return invoker(ctx, method, req, reply, cc, opts...)
			}), WithTimeout(1000*time.Millisecond))
	assert.Nil(t, err)
	clients := []Client{
		directClient,
		nonBlockClient,
		tarConfClient,
		targetClient,
	}
	SetClientSlowThreshold(time.Second)

	for _, tt := range tests {
		tt := tt
		for _, client := range clients {
			client := client
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				cli := mock.NewDepositServiceClient(client.Conn())
				request := &mock.DepositRequest{Amount: tt.amount}
				response, err := cli.Deposit(context.Background(), request)
				if response != nil {
					assert.True(t, len(response.String()) > 0)
					if response.GetOk() != tt.res.GetOk() {
						t.Error("response: expected", tt.res.GetOk(), "received", response.GetOk())
					}
				}
				if err != nil {
					if e, ok := status.FromError(err); ok {
						if e.Code() != tt.errCode {
							t.Error("error code: expected", codes.InvalidArgument, "received", e.Code())
						}
						if e.Message() != tt.errMsg {
							t.Error("error message: expected", tt.errMsg, "received", e.Message())
						}
					}
				}
			})
		}
	}
}

func TestNewClientWithError(t *testing.T) {
	_, err := NewClient(
		RpcClientConf{
			Timeout: 1000,
		},
		WithDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
		WithDialOption(grpc.WithContextDialer(dialer())),
		WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{},
			cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			return invoker(ctx, method, req, reply, cc, opts...)
		}),
	)
	assert.NotNil(t, err)

	_, err = NewClient(
		RpcClientConf{
			Target:  "localhost:8080",
			Timeout: 1,
		},
		WithDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
		WithDialOption(grpc.WithContextDialer(dialer())),
		WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{},
			cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			return invoker(ctx, method, req, reply, cc, opts...)
		}),
	)
	assert.NotNil(t, err)
}

type depositServer struct{}

func newDepositServer() *depositServer {
	return &depositServer{}
}

func (s *depositServer) Deposit(ctx context.Context, in *mock.DepositRequest) (*mock.DepositResponse, error) {
	fmt.Printf("mrpc client in %+v \n", in)
	return &mock.DepositResponse{
		Ok: true,
	}, nil
}

func TestMrpc(t *testing.T) {
	// start server
	srv := newDepositServer()

	sc := RpcServerConf{
		Addr:    "localhost:8081",
		Timeout: 60,
	}
	s := MustNewServer(sc, func(server *grpc.Server) {
		mock.RegisterDepositServiceServer(server, srv)
	})

	defer s.Stop()

	fmt.Printf("starting rpc server at %s...\n", sc.Addr)

	go s.Start()

	// start client
	cc := RpcClientConf{
		Target:  "localhost:8081",
		Timeout: 60,
	}
	deposit := mock.NewDepositServiceClient(MustNewClient(cc).Conn())

	fmt.Printf("connection mrpc server %s... \n", cc.Target)

	resp, err := deposit.Deposit(context.Background(), &mock.DepositRequest{
		Amount: 10,
	})
	if err != nil {
		fmt.Printf("mrpc deposit err %v \n", err)
		return
	}

	fmt.Printf("mrpc deposit succ %+v \n", resp)
}
