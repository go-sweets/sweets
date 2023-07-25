package migrate

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
)

func ApplyMigrations(dir migrate.MigrationDirection, dryrun bool, limit int, version int64) error {
	env, err := GetEnvironment()
	if err != nil {
		return fmt.Errorf("could not parse config: %s", err)
	}

	db, dialect, err := GetConnection(env)
	if err != nil {
		return err
	}
	defer db.Close()

	source := migrate.FileMigrationSource{
		Dir: env.Dir,
	}

	if dryrun {
		var migrations []*migrate.PlannedMigration

		if version >= 0 {
			migrations, _, err = migrate.PlanMigrationToVersion(db, dialect, source, dir, version)
		} else {
			migrations, _, err = migrate.PlanMigration(db, dialect, source, dir, limit)
		}

		if err != nil {
			return fmt.Errorf("cannot plan migration: %s", err)
		}

		for _, m := range migrations {
			PrintMigration(m, dir)
		}
	} else {
		var n int

		if version >= 0 {
			n, err = migrate.ExecVersion(db, dialect, source, dir, version)
		} else {
			n, err = migrate.ExecMax(db, dialect, source, dir, limit)
		}

		if err != nil {
			return fmt.Errorf("migration failed: %s", err)
		}

		if n == 1 {
			fmt.Println("Applied 1 migration")
		} else {
			fmt.Println(fmt.Sprintf("Applied %d migrations", n))
		}
	}

	return nil
}

func PrintMigration(m *migrate.PlannedMigration, dir migrate.MigrationDirection) {
	if dir == migrate.Up {
		fmt.Println(fmt.Sprintf("==> Would apply migration %s (up)", m.Id))
		for _, q := range m.Up {
			fmt.Println(q)
		}
	} else if dir == migrate.Down {
		fmt.Println(fmt.Sprintf("==> Would apply migration %s (down)", m.Id))
		for _, q := range m.Down {
			fmt.Println(q)
		}
	} else {
		panic("Not reached")
	}
}
