package dao

import (
	mgorm "share/database/gorm"

	"gorm.io/gorm"
)

type NewsRepository struct {
	Gorm *gorm.DB
	*mgorm.Closer
}

func NewNewsRepository(dsn string) (*NewsRepository, error) {
	db, err := mgorm.NewMysqlGormByDSN(dsn)
	if err != nil {
		return nil, err
	}

	return &NewsRepository{
		Gorm: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}
