package repository

import (
	"api-catalog-auto/config"
	"api-catalog-auto/internal/core/port/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//Note в рамках пакетов /internal/adapter описываются компоненты которых вызывает бизнес логика(core/service)

// database в рамках компонента заложена логика инициализации БД
type database struct {
	gormDB *gorm.DB
}

func (d database) CloseDB() (err error) {
	s, err := d.gormDB.DB()
	if err != nil {
		return
	}

	err = s.Close()
	return
}

func NewDB(cfg config.DBSettings) (repository.IDB, error) {
	logMode := logger.Error
	if cfg.LogMode {
		logMode = logger.Info
	}

	gormDB, err := gorm.Open(postgres.Open(cfg.ConnectionString), &gorm.Config{
		Logger:      logger.Default.LogMode(logMode),
		PrepareStmt: true,
	})

	if err != nil {
		return nil, err
	}

	postgresDb, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	postgresDb.SetMaxOpenConns(cfg.MaxOpenConns)
	postgresDb.SetMaxIdleConns(cfg.MaxIdleConns)

	return &database{gormDB: gormDB}, nil
}
