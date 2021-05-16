package dao

import (
	"fmt"
	"rec-service/model"
	zlog "share/log/zap"

	"go.uber.org/zap"
)

func (r *RecRepository) GetComicsByHot(offset int, limit int) ([]model.Comic, error) {
	var comics []model.Comic

	if err := r.db.Order("num desc").Limit(limit).Offset(offset).Find(&comics).Error; err != nil {
		zlog.Logger.Error("failed to retrive comics", zap.Error(err))

		return nil, fmt.Errorf("failed to retrive comics: %v", err)
	}

	return comics, nil
}
