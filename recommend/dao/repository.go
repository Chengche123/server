package dao

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type RecRepository struct {
	db *gorm.DB
}

func (r *RecRepository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sqlDB")
	}
	return sqlDB.Close()
}

func NewRecRepository(raw *sql.DB) (*RecRepository, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: raw,
	}), &gorm.Config{
		// Logger:                 logger.Default.LogMode(logger.Info),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cannot open db connect: %v", err)
	}

	if !db.Migrator().HasTable("comic") {
		return nil, fmt.Errorf("cannot find comic table in database")
	}

	return &RecRepository{
		db: db,
	}, nil
}
