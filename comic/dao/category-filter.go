package dao

import (
	"comic-service/model"
	"fmt"

	"gorm.io/gorm"
)

func (r *ComicRepository) FindComicCategoryFilter() ([]model.ComicCategoryFilter, error) {
	var res []model.ComicCategoryFilter
	if err := r.Gorm.Find(&res).Error; err != nil && err != gorm.ErrEmptySlice {
		return nil, fmt.Errorf("failed to find comic category filter: %v", err)
	}

	return res, nil
}
