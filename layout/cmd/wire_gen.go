// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/mix-plus/mixplus-layout/internal/config"
	"github.com/mix-plus/mixplus-layout/internal/server"
	"github.com/mix-plus/mixplus-layout/internal/service"
	"github.com/mix-plus/mixplus-layout/internal/svc"
)

// Injectors from wire.go:

// initApp init app application.
func initApp(c *config.Config) (*server.AppServer, error) {
	serviceContext := svc.NewServiceContext(c)
	helloService := service.NewHelloServer(serviceContext)
	httpServer := server.NewHttpServer(c, helloService)
	rpcServer := server.NewGrpcServer(c, serviceContext)
	appServer, err := server.NewApp(serviceContext, helloService, httpServer, rpcServer)
	if err != nil {
		return nil, err
	}
	return appServer, nil
}