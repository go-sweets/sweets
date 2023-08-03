package config

import (
	mpConf "github.com/mix-plus/go-mixplus/pkg/conf"
)

type Config struct {
	mpConf.GrpcConf  `mapstructure:"RpcServerConf"`
	mpConf.ApiConf   `mapstructure:"HttpServerConf"`
	mpConf.DbConf    `mapstructure:"DbConf"`
	mpConf.RedisConf `mapstructure:"RedisConf"`
}
