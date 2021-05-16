package dao

import (
	"comic-service/model"
	"fmt"
)

func (r *ComicRepository) FindComicSpecial(offset, limit int) ([]model.ComicSpecial, error) {
	var res []model.ComicSpecial

	if err := r.Gorm.Offset(offset).Limit(limit).Order("create_time DESC").Find(&res).Error; err != nil {
		return nil, fmt.Errorf("dao error: %v", err)
	}

	return res, nil
}
