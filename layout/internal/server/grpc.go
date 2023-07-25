package server

import (
	"github.com/mix-plus/go-mixplus/layout/internal/boundedcontexts/hello/application/handlers"
	"github.com/mix-plus/go-mixplus/layout/internal/config"
	"github.com/mix-plus/go-mixplus/layout/internal/service"
	"github.com/mix-plus/go-mixplus/layout/internal/svc"
	"github.com/zeromicro/go-zero/zrpc"

	hello "github.com/mix-plus/go-mixplus/layout/api/hello/v1"

	"google.golang.org/grpc"
)

func NewGrpcServer(c *config.Config,
	svc *svc.ServiceContext,
	handler *handlers.HelloGrpcHandler,
) *zrpc.RpcServer {

	srv := service.NewHelloServer(svc, handler)
	s := zrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		// grpc register
		hello.RegisterHelloServer(g, srv)
	})
	s.AddOptions()

	return s
}
