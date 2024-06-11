package gormDatabase

import (
	"time"

	"github.com/ahsansandiah/dealls-test/packages/config"
	"github.com/cenkalti/backoff"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm interface {
	Connect() (*gorm.DB, error)
}

type Options struct {
	master  string
	maxOpen int
	maxIdle int
}

func NewGorm(cfg *config.Config) Gorm {
	opt := new(Options)
	opt.master = cfg.DatabaseDSN
	opt.maxOpen = cfg.DatabaseMaxOpenConnections
	opt.maxIdle = cfg.DatabaseMaxIdleConnections

	return opt
}

func (o *Options) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(o.master), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(o.maxOpen)
	sqlDB.SetMaxIdleConns(o.maxIdle)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := backoff.Retry(func() error {
		if err := sqlDB.Ping(); err != nil {
			return err
		}

		return nil
	}, backoff.NewExponentialBackOff()); err != nil {
		return nil, err
	}

	return db, nil
}
