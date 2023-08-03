package server

import (
	"net/http"

	"github.com/google/wire"
	"github.com/mix-plus/go-mixplus/layout/internal/service"
	"github.com/mix-plus/go-mixplus/layout/internal/svc"
	"github.com/zeromicro/go-zero/zrpc"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)

type AppServer struct {
	SvcCtx     *svc.ServiceContext
	HTTPServer *http.Server
	GrpcServer *zrpc.RpcServer

	HelloService *service.HelloService
}

func NewApp(svcCtx *svc.ServiceContext, helloService *service.HelloService, hs *http.Server, gs *zrpc.RpcServer) (*AppServer, error) {
	return &AppServer{
		SvcCtx:       svcCtx,
		HelloService: helloService,
		HTTPServer:   hs,
		GrpcServer:   gs,
	}, nil
}

func (a *AppServer) Run() {

	go func() {
		err := a.HTTPServer.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	a.GrpcServer.Start()

	defer a.GrpcServer.Stop()
}
