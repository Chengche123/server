package dao

import (
	mgorm "share/database/gorm"

	"gorm.io/gorm"
)

type RecRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewRecRepository(dsn string) (*RecRepository, error) {
	db, err := mgorm.NewMysqlGormByDSN(dsn)
	if err != nil {
		return nil, err
	}

	return &RecRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}
