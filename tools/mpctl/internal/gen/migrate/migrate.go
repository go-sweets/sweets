package migrate

import "github.com/spf13/cobra"

const (
	DefaultConfig = "./etc/gen.yml"
	DefaultEnv    = "gen"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "migrate: Generate Directory. migrate status, migrate up, migrate down",
	Long:  "migrate: Generate Directory. migrate status, migrate up, migrate down",
}

func init() {
	CmdMigrate.AddCommand(migrateUp)
	CmdMigrate.AddCommand(migrateDown)
	CmdMigrate.AddCommand(migrateRedo)
	CmdMigrate.AddCommand(migrateStatus)
	CmdMigrate.AddCommand(migrateNew)
	CmdMigrate.AddCommand(migrateSkip)

	CmdMigrate.PersistentFlags().StringP("config", "c", DefaultConfig, "Database configuration file.")
	CmdMigrate.PersistentFlags().StringP("env", "e", DefaultEnv, "Environment.")
}
