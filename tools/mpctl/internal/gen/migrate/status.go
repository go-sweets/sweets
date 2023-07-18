package migrate

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var migrateStatus = &cobra.Command{
	Use:   "status",
	Short: "Show migration status.",
	Long:  "Show migration status.",
	Run:   MigrateStatusRun,
}

type statusRow struct {
	Id        string
	Migrated  bool
	AppliedAt time.Time
}

func MigrateStatusRun(cmd *cobra.Command, args []string) {
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
	migrations, err := source.FindMigrations()
	if err != nil {
		panic(err)
	}

	records, err := migrate.GetMigrationRecords(db, dialect)
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Migration", "Applied"})
	table.SetColWidth(60)

	rows := make(map[string]*statusRow)

	for _, m := range migrations {
		rows[m.Id] = &statusRow{
			Id:       m.Id,
			Migrated: false,
		}
	}

	for _, r := range records {
		if rows[r.Id] == nil {
			fmt.Println(fmt.Sprintf("Could not find migration file: %v", r.Id))
			continue
		}

		rows[r.Id].Migrated = true
		rows[r.Id].AppliedAt = r.AppliedAt
	}

	for _, m := range migrations {
		if rows[m.Id] != nil && rows[m.Id].Migrated {
			table.Append([]string{
				m.Id,
				rows[m.Id].AppliedAt.String(),
			})
		} else {
			table.Append([]string{
				m.Id,
				"no",
			})
		}
	}

	table.Render()
}
