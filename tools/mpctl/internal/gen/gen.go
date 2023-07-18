package gen

import (
	"github.com/mix-plus/go-mixplus/tools/mpctl/internal/gen/gorm"
	"github.com/mix-plus/go-mixplus/tools/mpctl/internal/gen/migrate"
	"github.com/spf13/cobra"
)

var CmdGen = &cobra.Command{
	Use:   "gen",
	Short: "gen: Generate Directory. gen gorm ",
	Long:  "gen: Generate Directory. gen gorm ",
}

func init() {
	CmdGen.AddCommand(gorm.CmdGorm)
	CmdGen.AddCommand(migrate.CmdMigrate)
}
