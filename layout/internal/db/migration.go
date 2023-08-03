package db

import (
	"database/sql"
	"time"

	"github.com/google/wire"
	"github.com/mix-plus/go-mixplus/layout/internal/db/migrations"
	"github.com/mix-plus/go-mixplus/layout/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
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
	migrations := m.migrations
	db, err := m.svc.DB.DB()
	if err != nil {
		logx.Errorf("migration get db error:%v", err)
		return err
	}
	logx.Infof("migration start length:%d", len(migrations))

	for _, migration := range migrations {
		logx.Infof("migration perform migration id: %s", migration.Name())
		// 检查迁移是否已经执行过
		var id bool
		err := db.QueryRow("SELECT COUNT(*) FROM data_migrations WHERE id = ?", migration.Name()).Scan(&id)
		if err != nil && err != sql.ErrNoRows {
			logx.Errorf("migration select error:%v", err)
			return err
		}

		if id {
			// 这个迁移已经执行过，所以跳过
			logx.Infof("migration jump over id:%s", migration.Name())
			return nil
		}

		// 执行迁移
		if err := migration.Up(); err != nil {
			logx.Errorf("migration up error:%v", err)
			upErr := err
			if err = migration.Down(); err != nil {
				// 回滚错误
				logx.Errorf("migration down error:%v", err)
				return err
			}
			return upErr
		}

		// 记录迁移
		_, err = db.Exec("INSERT INTO data_migrations (id,applied_at) VALUES (?, ?)", migration.Name(), time.Now())
		if err != nil {
			logx.Errorf("migration exec error:%v", err)
			return err
		}
	}
	logx.Infof("migration end length:%d", len(migrations))
	return nil
}
