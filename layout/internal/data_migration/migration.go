package data_migration

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/mix-plus/go-mixplus/layout/internal/data_migration/migrations"
	"github.com/mix-plus/go-mixplus/layout/internal/svc"
	"time"
)

var ProviderSet = wire.NewSet(CreateDataMigrations, svc.ProviderSet, NewMigrator)

func CreateDataMigrations() []Migration {
	return []Migration{
		&migrations.UserMigration{},
	}
}

const TableName = "data_migrations"

type Migration interface {
	Name() string // 迁移的名字，必须是唯一的
	Up() error    // 执行迁移
	Down() error  // 回滚迁移
}

type Migrator struct {
	svc        *svc.ServiceContext
	migrations []Migration
}

func NewMigrator(migrations []Migration, svc *svc.ServiceContext) *Migrator {
	return &Migrator{
		svc:        svc,
		migrations: migrations,
	}
}

func (m *Migrator) Migrate() error {
	for _, migration := range m.migrations {
		if err := m.performMigration(migration); err != nil {
			return err
		}
	}
	return nil
}

func (m *Migrator) performMigration(migration Migration) error {
	// 检查迁移是否已经执行过
	var name string
	err := m.svc.DB.Select("SELECT id FROM ? WHERE id = ? LIMIT 1", TableName, migration.Name()).Scan(&name).Error
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if name != "" {
		// 这个迁移已经执行过，所以跳过
		return nil
	}

	// 执行迁移
	if err := migration.Up(); err != nil {
		// 迁移错误 回滚
		upErr := err
		if err = migration.Down(); err != nil {
			// 回滚错误
			return err
		}
		return upErr
	}

	// 记录迁移
	err = m.svc.DB.Exec("INSERT INTO ? (id, applied_at) VALUES (?, ?)", TableName, migration.Name(), time.Now()).Error
	if err != nil {
		return err
	}

	return nil
}
