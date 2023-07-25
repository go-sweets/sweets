package migrate

import (
	"database/sql"
	"embed"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

// RunMigration automatically run migration
func RunMigration(dsn string, migrationsDir embed.FS) {
	createSchemaIfNotExists(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logx.Errorf("Failed to connect database for migration: %v", err)
	}

	migrate.SetTable("schema_migrations")
	migrate.SetIgnoreUnknown(true)
	migrationsSource := &migrate.HttpFileSystemMigrationSource{
		FileSystem: http.FS(migrationsDir),
	}

	migrationScripts, err := migrationsSource.FindMigrations()
	if err != nil {
		logx.Errorf("Could not find migration scripts: %v", err)
	}
	if len(migrationScripts) == 0 {
		logx.Infof("No migration script to migrate Database::migration_scripts_len %d", len(migrationScripts))
	} else {
		logx.Infof("Migration scripts count Database::migration_scripts_len %d", len(migrationScripts))
	}

	n, err := migrate.Exec(db, "mysql", migrationsSource, migrate.Up)
	if err != nil {
		logx.Errorf("Failed to run database migration: %v", err)
	}
	logx.Infof("Finish running migration scripts Database::migration_scripts_len %d", n)
}

func createSchemaIfNotExists(dsn string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect database for schema creation")
	}

	logx.Info("Creating database schema if needed, schema_migrations")
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS schema_migrations")
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic("Failed to close database for schema creation")
	}
}
