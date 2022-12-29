package mrpc

import (
	"fmt"
	"github.com/mix-plus/go-mixplus/mrpc/internal"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/stat"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestServer_setupInterceptors(t *testing.T) {
	server := new(mockedServer)
	err := setupInterceptors(server, RpcServerConf{
		Timeout: 100,
	}, new(stat.Metrics))
	fmt.Println(err)
	fmt.Println(server)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(server.unaryInterceptors))
	assert.Equal(t, 1, len(server.streamInterceptors))
}

func TestServer(t *testing.T) {
	//DontLogContentForMethod("foo")
	SetServerSlowThreshold(time.Second)
	svr := MustNewServer(RpcServerConf{

		Addr:    "localhost:8081",
		Timeout: 0,
	}, func(server *grpc.Server) {
	})
	svr.AddOptions(grpc.ConnectionTimeout(time.Hour))
	//svr.AddUnaryInterceptors(serverinterceptors.UnaryCrashInterceptor)
	//svr.AddStreamInterceptors(serverinterceptors.StreamCrashInterceptor)
	go svr.Start()
	svr.Stop()
}

func TestServerError(t *testing.T) {
	_, err := NewServer(RpcServerConf{
		Addr: "localhost:8080",
	}, func(server *grpc.Server) {
	})
	assert.NotNil(t, err)
}

func TestServer_HasEtcd(t *testing.T) {
	svr := MustNewServer(RpcServerConf{
		Addr: "localhost:8080",
	}, func(server *grpc.Server) {
	})
	svr.AddOptions(grpc.ConnectionTimeout(time.Hour))
	//svr.AddUnaryInterceptors(serverinterceptors.UnaryCrashInterceptor)
	//svr.AddStreamInterceptors(serverinterceptors.StreamCrashInterceptor)
	go svr.Start()
	svr.Stop()
}

func TestServer_StartFailed(t *testing.T) {
	svr := MustNewServer(RpcServerConf{
		Addr: "localhost:aaa",
	}, func(server *grpc.Server) {
	})

	assert.Panics(t, svr.Start)
}

type mockedServer struct {
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
}

func (m *mockedServer) AddOptions(_ ...grpc.ServerOption) {
}

func (m *mockedServer) AddStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) {
	m.streamInterceptors = append(m.streamInterceptors, interceptors...)
}

func (m *mockedServer) AddUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) {
	m.unaryInterceptors = append(m.unaryInterceptors, interceptors...)
}

func (m *mockedServer) SetName(_ string) {
}

func (m *mockedServer) Start(_ internal.RegisterFn) error {
	return nil
}
