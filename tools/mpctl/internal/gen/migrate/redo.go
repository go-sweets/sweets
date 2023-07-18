package migrate

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var migrateRedo = &cobra.Command{
	Use:   "redo",
	Short: "Reapply the last migration.",
	Long:  "Reapply the last migration.",
	Run:   MigrateRedoRun,
}

func init() {
	migrateRedo.Flags().BoolP("dryrun", "d", false, "Don't apply migrations, just print them.")
}

func MigrateRedoRun(cmd *cobra.Command, args []string) {
	dryrun, _ := cmd.Flags().GetBool("dryrun")
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

	migrations, _, err := migrate.PlanMigration(db, dialect, source, migrate.Down, 1)
	if err != nil {
		panic(fmt.Sprintf("Migration (redo) failed: %v", err))
	} else if len(migrations) == 0 {
		panic("Nothing to do!")
	}

	if dryrun {
		PrintMigration(migrations[0], migrate.Down)
		PrintMigration(migrations[0], migrate.Up)
	} else {
		_, err := migrate.ExecMax(db, dialect, source, migrate.Down, 1)
		if err != nil {
			panic(fmt.Sprintf("Migration (down) failed: %s", err))
		}

		_, err = migrate.ExecMax(db, dialect, source, migrate.Up, 1)
		if err != nil {
			panic(fmt.Sprintf("Migration (up) failed: %s", err))
		}

		fmt.Println(fmt.Sprintf("Reapplied migration %s.", migrations[0].Id))
	}
}
