package gorm

import (
	"github.com/spf13/cast"
	xgorm "gorm.io/gorm"
)

const (
	// defaultPoolMaxIdleConns .
	defaultPoolMaxIdleConns = 10
	// defaultPoolMaxIdleConns .
	defaultPoolMaxOpenConns = 100
	// defaultPoolMaxLifetime .
	defaultPoolMaxLifetime = 3600
)

// NewDB instance db client with default pool setting
func NewDB(dial xgorm.Dialector, opts ...xgorm.Option) (*xgorm.DB, error) {
	db, err := xgorm.Open(dial, opts...)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(defaultPoolMaxIdleConns)
	sqlDB.SetMaxOpenConns(defaultPoolMaxOpenConns)
	sqlDB.SetConnMaxLifetime(cast.ToDuration(defaultPoolMaxLifetime))

	return db, nil
}
