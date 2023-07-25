package migrate

import (
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var migrateUp = &cobra.Command{
	Use:   "up",
	Short: "Migrates the database to the most recent version available.",
	Long:  "Migrates the database to the most recent version available.",
	Run:   MigrateUpRun,
}

func init() {
	migrateUp.PersistentFlags().IntP("limit", "l", 0, "Limit the number of migrations (0 = unlimited).")
	migrateUp.PersistentFlags().Int64P("version", "v", -1, "Run migrate up to a specific version, eg: the version number of migration 1_initial.sql is 1.")
	migrateUp.PersistentFlags().BoolP("dryrun", "d", false, "Don't apply migrations, just print them.")
}

func MigrateUpRun(cmd *cobra.Command, args []string) {
	limit, _ := cmd.Flags().GetInt("limit")
	version, _ := cmd.Flags().GetInt64("version")
	dryrun, _ := cmd.Flags().GetBool("dryrun")
	ConfigFlags(cmd)
	err := ApplyMigrations(migrate.Up, dryrun, limit, version)
	if err != nil {
		panic(err)
	}
}
