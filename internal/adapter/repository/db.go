package repository

import (
	"api-catalog-auto/config"
	"api-catalog-auto/internal/core/port/repository"
	"api-catalog-auto/internal/entity"
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

// AutoMigrate автомиграция данных
func (d database) AutoMigrate() {
	d.gormDB.AutoMigrate(
		&entity.Catalog{},
	)
}

// AutoMigrateData заполнение типов доуступов
func (d database) AutoMigrateData() {
	var count int64
	d.gormDB.Find(&entity.Catalog{}).Count(&count)
	if count == 0 {
		d.gormDB.Create(&entity.Catalog{ID: "70a69975-6db6-4ad7-b966-b29f383114f7", Mark: "Audi", Model: "A4", Year: 2021})
		d.gormDB.Create(&entity.Catalog{ID: "9fa70a07-7266-41fa-93de-7709b8eb0fca", Mark: "BMW", Model: "X5", Year: 2020})
		d.gormDB.Create(&entity.Catalog{ID: "53db4674-26a5-4e54-83b5-582e7b8062a1", Mark: "Mercedes", Model: "E-class", Year: 2019})
		d.gormDB.Create(&entity.Catalog{ID: "870c2cf7-0141-46c9-82ab-5ac0655909f9", Mark: "Toyota", Model: "Camry", Year: 2018})
	}
}
