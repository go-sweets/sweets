//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/mix-plus/mixplus-layout/internal/config"
	"github.com/mix-plus/mixplus-layout/internal/data_migration"
	"github.com/mix-plus/mixplus-layout/internal/server"
	"github.com/mix-plus/mixplus-layout/internal/service"
	"github.com/mix-plus/mixplus-layout/internal/svc"

	"github.com/google/wire"
)

// initApp init app application.
func initApp(c *config.Config) (*server.AppServer, error) {
	panic(wire.Build(svc.ProviderSet, service.ProviderSet, server.ProviderSet, server.NewApp))
}

func wireMigrate(c *config.Config) *data_migration.Migrator {
	panic(wire.Build(data_migration.ProviderSet))
}
