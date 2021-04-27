package dao

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	_ "gorm.io/driver/mysql"
)

type RecRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (r *RecRepository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sqlDB")
	}
	return sqlDB.Close()
}

func NewRecRepository(raw *sql.DB, logger *zap.Logger) (*RecRepository, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

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
		logger: logger,
		db:     db,
	}, nil
}

func NewRecRepositoryByDSN(dsn string, logger *zap.Logger) (*RecRepository, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	rawDB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/comic")
	if err != nil {
		return nil, fmt.Errorf("cannot open sql connect: %v", err)
	}
	err = rawDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %v", err)
	}

	return NewRecRepository(rawDB, logger)
}
