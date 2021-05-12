package dao

import (
	mgorm "comic/share/database/gorm"
	log "comic/share/log/zap"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ComicRepository struct {
	Gorm   *gorm.DB
	Logger *zap.Logger
	*mgorm.Closer
}

func NewComicRepository(dsn string) (*ComicRepository, error) {
	db, err := mgorm.NewMysqlGormByDSN(dsn)
	if err != nil {
		return nil, fmt.Errorf("failto init gorm db: %v", err)
	}

	return &ComicRepository{
		Gorm:   db,
		Logger: log.Logger,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}
