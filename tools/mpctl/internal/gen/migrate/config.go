package migrate

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime/debug"

	"github.com/go-gorp/gorp/v3"
	migrate "github.com/rubenv/sql-migrate"
	"gopkg.in/yaml.v3"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DefaultDialect   = "mysql"
	DefaultDir       = "internal/db/migrations"
	DefaultTableName = "schema_migrations"
)

var dialects = map[string]gorp.Dialect{
	"sqlite3":  gorp.SqliteDialect{},
	"postgres": gorp.PostgresDialect{},
	"mysql":    gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"},
}

var ConfigFile string
var ConfigEnvironment string

func init() {

}

func ConfigFlags(f *cobra.Command) {
	ConfigFile, _ = f.Flags().GetString("config")
	ConfigEnvironment, _ = f.Flags().GetString("env")
}

type Environment struct {
	Dialect       string `yaml:"dialect"`
	DSN           string `yaml:"dsn"`
	Dir           string `yaml:"dir"`
	TableName     string `yaml:"table"`
	SchemaName    string `yaml:"schema"`
	IgnoreUnknown bool   `yaml:"ignoreunknown"`
}

func ReadConfig() (map[string]*Environment, error) {
	file, err := os.ReadFile(ConfigFile)
	if err != nil {
		return nil, err
	}

	config := make(map[string]*Environment)
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetEnvironment() (*Environment, error) {
	config, err := ReadConfig()
	if err != nil {
		return nil, err
	}

	env := config[ConfigEnvironment]
	if env == nil {
		return nil, errors.New("no environment: " + ConfigEnvironment)
	}

	if env.Dialect == "" {
		env.Dialect = DefaultDialect
	}

	if env.DSN == "" {
		return nil, errors.New("no data source specified")
	}
	env.DSN = os.ExpandEnv(env.DSN)

	if env.Dir == "" {
		env.Dir = DefaultDir
	}

	if env.TableName == "" {
		env.TableName = DefaultTableName
	}

	if env.TableName != "" {
		migrate.SetTable(env.TableName)
	}

	if env.SchemaName != "" {
		migrate.SetSchema(env.SchemaName)
	}

	migrate.SetIgnoreUnknown(env.IgnoreUnknown)

	return env, nil
}

func GetConnection(env *Environment) (*sql.DB, string, error) {
	db, err := sql.Open(env.Dialect, env.DSN)
	if err != nil {
		return nil, "", fmt.Errorf("cannot connect to database: %s", err)
	}

	// Make sure we only accept dialects that were compiled in.
	_, exists := dialects[env.Dialect]
	if !exists {
		return nil, "", fmt.Errorf("unsupported dialect: %s", env.Dialect)
	}

	return db, env.Dialect, nil
}

// GetVersion returns the version.
func GetVersion() string {
	if buildInfo, ok := debug.ReadBuildInfo(); ok && buildInfo.Main.Version != "(devel)" {
		return buildInfo.Main.Version
	}
	return "dev"
}
