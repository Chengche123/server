package dao

import (
	mgorm "share/database/gorm"

	"gorm.io/gorm"
)

type ComicRepository struct {
	Gorm *gorm.DB
	*mgorm.Closer
}

func NewComicRepository(dsn string) (*ComicRepository, error) {
	db, err := mgorm.NewMysqlGormByDSN(dsn)
	if err != nil {
		return nil, err
	}

	return &ComicRepository{
		Gorm: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}
