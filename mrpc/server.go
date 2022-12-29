package mrpc

import (
	"github.com/mix-plus/go-mixplus/mrpc/internal"
	"github.com/mix-plus/go-mixplus/mrpc/internal/serverinterceptors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stat"
	"google.golang.org/grpc"
	"log"
	"time"
)

type RpcServer struct {
	server   internal.Server
	register internal.RegisterFn
}

func MustNewServer(c RpcServerConf, register internal.RegisterFn) *RpcServer {
	server, err := NewServer(c, register)
	if err != nil {
		log.Fatal(err)
	}

	return server
}

func NewServer(c RpcServerConf, register internal.RegisterFn) (*RpcServer, error) {
	var err error
	// TODO ConfValidate

	var server internal.Server
	metrics := stat.NewMetrics(c.Addr)
	serverOptions := []internal.ServerOption{
		internal.WithMetrics(metrics),
	}

	server = internal.NewRpcServer(c.Addr, serverOptions...)

	server.SetName(c.Name)
	if err = setupInterceptors(server, c, metrics); err != nil {
		return nil, err
	}

	rpcServer := &RpcServer{
		server:   server,
		register: register,
	}

	return rpcServer, nil
}

func (rs *RpcServer) AddOptions(options ...grpc.ServerOption) {
	rs.server.AddOptions(options...)
}

func (rs *RpcServer) AddStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) {
	rs.server.AddStreamInterceptors(interceptors...)
}

func (rs *RpcServer) AddUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) {
	rs.server.AddUnaryInterceptors(interceptors...)
}

// Start starts the RpcServer.
// Graceful shutdown is enabled by default.
// Use proc.SetTimeToForceQuit to customize the graceful shutdown period.
func (rs *RpcServer) Start() {
	if err := rs.server.Start(rs.register); err != nil {
		logx.Error(err)
		panic(err)
	}
}

func (rs *RpcServer) Stop() {
	logx.Close()
}

func DontLogContextForMethod(method string) {
	// TODO
}

func SetServerSlowThreshold(threshold time.Duration) {
	// TODO
}

func setupInterceptors(server internal.Server, c RpcServerConf, metrics *stat.Metrics) error {
	if c.Timeout > 0 {
		server.AddUnaryInterceptors(serverinterceptors.UnaryTimeoutInterceptor(
			time.Duration(c.Timeout) * time.Millisecond))
	}
	// TODO Auth
	return nil
}
