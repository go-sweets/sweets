package main

import (
	"flag"

	"github.com/mix-plus/go-mixplus/core/conf"
	"github.com/mix-plus/go-mixplus/layout/internal/config"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	// data migration
	err := wireMigrate(&c).Migrate()
	if err != nil {
		// TODO
	}
	app, err := initApp(&c)
	if err != nil {
		panic(err)
	}

	app.Run()
}
