package migrate

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var migrateSkip = &cobra.Command{
	Use:   "skip",
	Short: "Set the database level to the most recent version available, without actually running the migrations.",
	Long:  "Set the database level to the most recent version available, without actually running the migrations.",
	Run:   MigrateSkipRun,
}

func init() {
	migrateSkip.Flags().IntP("limit", "l", 0, "Limit the number of migrations (0 = unlimited).")
}
func MigrateSkipRun(cmd *cobra.Command, args []string) {
	limit, _ := cmd.Flags().GetInt("limit")
	ConfigFlags(cmd)
	env, err := GetEnvironment()
	if err != nil {
		panic(err)
	}
	db, dialect, err := GetConnection(env)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	source := migrate.FileMigrationSource{
		Dir: env.Dir,
	}

	n, err := migrate.SkipMax(db, dialect, source, migrate.Up, limit)
	if err != nil {
		panic(fmt.Errorf("migration failed: %s", err))
	}
	switch n {
	case 0:
		fmt.Println("All migrations have already been applied")
	case 1:
		fmt.Println("Skipped 1 migration")
	default:
		fmt.Println(fmt.Sprintf("Skipped %d migrations", n))
	}
}
