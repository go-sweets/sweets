package svc

import (
	"github.com/mix-plus/mixplus-layout/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServiceContext)

type ServiceContext struct {
	Config *config.Config
	DB     *gorm.DB
}

func NewServiceContext(c *config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DbConf.DSN))
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
