package main

import (
	"flag"
	"github.com/mix-plus/go-mixplus/layout/internal/db/migrations"
	"github.com/mix-plus/go-mixplus/pkg/migrate"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/mix-plus/go-mixplus/layout/internal/config"
	"github.com/mix-plus/go-mixplus/pkg/conf"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	// sql migration
	migrate.RunMigration(c.DSN, migrations.Fs)
	// data migration
	err := wireMigrate(&c).Migrate()
	if err != nil {
		logx.Errorf("Exec Migration error:%v", err)
	}
	app, err := initApp(&c)
	if err != nil {
		panic(err)
	}

	app.Run()
}
