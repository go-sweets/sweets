package server

import (
	"net/http"

	"github.com/google/wire"
	"github.com/mix-plus/go-mixplus/mrpc"
	"github.com/mix-plus/mixplus-layout/internal/service"
	"github.com/mix-plus/mixplus-layout/internal/svc"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)

type AppServer struct {
	SvcCtx     *svc.ServiceContext
	HttpServer *http.Server
	GrpcServer *mrpc.RpcServer

	HelloService *service.HelloService
}

func NewApp(svcCtx *svc.ServiceContext, helloService *service.HelloService, hs *http.Server, gs *mrpc.RpcServer) (*AppServer, error) {
	return &AppServer{
		SvcCtx:       svcCtx,
		HelloService: helloService,
		HttpServer:   hs,
		GrpcServer:   gs,
	}, nil
}

func (a *AppServer) Run() {

	go func() {
		err := a.HttpServer.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	a.GrpcServer.Start()

	defer a.GrpcServer.Stop()
}
