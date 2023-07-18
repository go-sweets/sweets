package config

import (
	mpConf "github.com/mix-plus/go-mixplus/core/conf"
	"github.com/mix-plus/go-mixplus/mrpc"
)

type Config struct {
	mrpc.RpcServerConf `mapstructure:"RpcServerConf"`
	mpConf.ApiConf     `mapstructure:"HttpServerConf"`
}
