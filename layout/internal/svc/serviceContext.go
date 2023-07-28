package svc

import (
	"github.com/mix-plus/go-mixplus/layout/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServiceContext)

type ServiceContext struct {
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c *config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DbConf.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.RedisConf.Addr,
		Password: c.RedisConf.Pass,
		DB:       int(c.RedisConf.DataBase),
	})
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  rdb,
	}
}
