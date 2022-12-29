package main

import (
	"github.com/mix-go/dotenv"
	"github.com/mix-go/xcli"
	"github.com/mix-plus/go-mixplus/tools/mpctl/commands"
)

func main() {
	xcli.SetName("mpctl").
		SetVersion(commands.CLIVersion).
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Cmds...).Run()
}
