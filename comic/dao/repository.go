package dao

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ComicRepository struct {
	Gorm   *gorm.DB
	Logger *zap.Logger
}

func (r *ComicRepository) Close() error {
	raw, _ := r.Gorm.DB()
	return raw.Close()
}
