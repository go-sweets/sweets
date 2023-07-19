package server

import (
	"github.com/mix-plus/go-mixplus/layout/internal/config"
	"github.com/mix-plus/go-mixplus/layout/internal/service"
	"github.com/mix-plus/go-mixplus/layout/internal/svc"

	hello "github.com/mix-plus/go-mixplus/layout/api/hello/v1"

	"github.com/mix-plus/go-mixplus/mrpc"
	"google.golang.org/grpc"
)

func NewGrpcServer(c *config.Config, svc *svc.ServiceContext) *mrpc.RpcServer {

	srv := service.NewHelloServer(svc)
	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		// grpc register
		hello.RegisterHelloServer(g, srv)
	})
	s.AddOptions()

	return s
}
