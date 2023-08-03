package service

import (
	"context"

	hello "github.com/mix-plus/go-mixplus/layout/api/hello/v1"
	"github.com/mix-plus/go-mixplus/layout/internal/boundedcontexts/hello/application/handlers"
	"github.com/mix-plus/go-mixplus/layout/internal/svc"
)

type HelloService struct {
	hello.UnimplementedHelloServer

	svcCtx       *svc.ServiceContext
	helloHandler *handlers.HelloGrpcHandler
}

func NewHelloServer(ctx *svc.ServiceContext,
	helloHandler *handlers.HelloGrpcHandler,
) *HelloService {
	return &HelloService{
		svcCtx:       ctx,
		helloHandler: helloHandler,
	}
}

func (service *HelloService) SayHello(ctx context.Context, in *hello.HelloReq) (*hello.HelloResp, error) {
	resp, err := service.helloHandler.SayHello(ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
