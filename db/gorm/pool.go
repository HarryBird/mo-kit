package gorm

import (
	"github.com/spf13/cast"
	xgorm "gorm.io/gorm"
)

// WithPool set gorm.db pool
func WithPool(db *xgorm.DB, pool Pooler) error {
	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	if pool.GetMaxIdle() > 0 {
		sqlDB.SetMaxIdleConns(cast.ToInt(pool.GetMaxIdle()))
	}

	if pool.GetMaxOpen() > 0 {
		sqlDB.SetMaxOpenConns(cast.ToInt(pool.GetMaxOpen()))
	}

	if pool.GetMaxLifeTime() > 0 {
		sqlDB.SetConnMaxLifetime(cast.ToDuration(pool.GetMaxLifeTime()))
	}

	return nil
}
