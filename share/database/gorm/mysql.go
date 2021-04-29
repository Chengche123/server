package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewMysqlGormByDSN(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		// Logger:                 logger.Default.LogMode(logger.Info),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cannot init gorm object: %v", err)
	}

	return db, nil
}
