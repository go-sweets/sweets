package service

import (
	"github.com/google/wire"
	"github.com/mix-plus/go-mixplus/layout/internal/boundedcontexts/hello"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(hello.InjectHelloRepository, hello.InjectHelloGrpcHandler, NewHelloServer)
