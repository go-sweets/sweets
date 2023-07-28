package migrate

import (
	"database/sql"
	"embed"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

// RunMigration automatically run migration
func RunMigration(dsn string, migrationsDir embed.FS) error {
	err := createSchemaIfNotExists(dsn)
	if err != nil {
		return err
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logx.Errorf("Failed to connect database for migration: %v", err)
		return err
	}

	migrate.SetTable("schema_migrations")
	migrate.SetIgnoreUnknown(true)
	migrationsSource := &migrate.HttpFileSystemMigrationSource{
		FileSystem: http.FS(migrationsDir),
	}

	migrationScripts, err := migrationsSource.FindMigrations()
	if err != nil {
		logx.Errorf("Could not find migration scripts: %v", err)
		return err
	}
	if len(migrationScripts) == 0 {
		logx.Infof("No migration script to migrate Database::migration_scripts_len %d", len(migrationScripts))
	} else {
		logx.Infof("Migration scripts count Database::migration_scripts_len %d", len(migrationScripts))
	}

	n, err := migrate.Exec(db, "mysql", migrationsSource, migrate.Up)
	if err != nil {
		logx.Errorf("Failed to run database migration: %v", err)
		return err
	}
	logx.Infof("Finish running migration scripts Database::migration_scripts_len %d", n)

	return nil
}

func createSchemaIfNotExists(dsn string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logx.Errorf("Failed to connect database for schema creation: %v", err)
		return err
	}

	logx.Info("Creating database schema if needed, schema_migrations")
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS schema_migrations")
	if err != nil {
		logx.Errorf("Failed to create database schema: %v", err)
		return err
	}

	err = db.Close()
	if err != nil {
		logx.Errorf("Failed to close database connection for schema creation: %v", err)
		return err
	}

	return nil
}
