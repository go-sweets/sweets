package migrate

import (
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var migrateDown = &cobra.Command{
	Use:   "down",
	Short: "Undo a database migration.",
	Long:  "Undo a database migration.",
	Run:   MigrateDownRun,
}

func init() {
	migrateDown.PersistentFlags().IntP("limit", "l", 1, "Limit the number of migrations (0 = unlimited).")
	migrateDown.PersistentFlags().Int64P("version", "v", -1, "Run migrate down to a specific version, eg: the version number of migration 1_initial.sql is 1.")
	migrateDown.PersistentFlags().BoolP("dryrun", "d", false, "Don't apply migrations, just print them.")
}

func MigrateDownRun(cmd *cobra.Command, args []string) {
	limit, _ := cmd.Flags().GetInt("limit")
	version, _ := cmd.Flags().GetInt64("version")
	dryrun, _ := cmd.Flags().GetBool("dryrun")
	ConfigFlags(cmd)
	err := ApplyMigrations(migrate.Down, dryrun, limit, version)
	if err != nil {
		panic(err)
	}
}
