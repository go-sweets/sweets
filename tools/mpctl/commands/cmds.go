package commands

import "github.com/mix-go/xcli"

var Cmds = []*xcli.Command{
	{
		Name:  "new",
		Short: "\tCreate a project",
		RunI:  &NewCommand{},
	},
	{
		Name:  "upgrade",
		Short: "Upgrade mpctl to latest version",
		RunI:  &UpgradeCommand{},
	},
}
