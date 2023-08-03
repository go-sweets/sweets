package main

import (
	"flag"

	"github.com/mix-plus/go-mixplus/layout/internal/config"
	"github.com/mix-plus/go-mixplus/layout/internal/db/migrations"
	"github.com/mix-plus/go-mixplus/pkg/conf"
	"github.com/mix-plus/go-mixplus/pkg/migrate"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	// sql migration
	err := migrate.RunMigration(c.DSN, migrations.Fs)
	if err != nil {
		logx.Errorf("Exec Sql Migration error:%v", err)
	}
	// data migration
	err = wireMigrate(&c).Migrate()
	if err != nil {
		logx.Errorf("Exec Data Migration error:%v", err)
	}
	app, err := initApp(&c)
	if err != nil {
		panic(err)
	}

	app.Run()
}
