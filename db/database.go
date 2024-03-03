package database

import (
	config "akozadaev/go_reports/pkg"
	"gorm.io/driver/sqlite"
	"time"

	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	for i := 0; i <= cfg.DBConfig.Pool.MaxOpen; i++ {
		db, err = gorm.Open(sqlite.Open(cfg.DBConfig.DataSourceName), &gorm.Config{TranslateError: true})
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		return nil, err
	}

	rawDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	rawDB.SetMaxOpenConns(cfg.DBConfig.Pool.MaxOpen)
	rawDB.SetMaxIdleConns(cfg.DBConfig.Pool.MaxIdle)
	rawDB.SetConnMaxLifetime(cfg.DBConfig.Pool.MaxLifetime)

	return db, nil
}
