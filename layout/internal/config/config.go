package config

import (
	mpConf "github.com/mix-plus/go-mixplus/pkg/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf `mapstructure:"RpcServerConf"`
	mpConf.ApiConf     `mapstructure:"HttpServerConf"`
	mpConf.DbConf      `mapstructure:"DbConf"`
	mpConf.RedisConf   `mapstructure:"RedisConf"`
}
