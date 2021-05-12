package dao

import (
	"comic-service/model"
	"fmt"
)

func (r *ComicRepository) FindCategoryDetail(types string, sort, offset, limit int) ([]model.CategoryDetail, error) {
	var res []model.CategoryDetail

	types = "%" + types + "%"

	tx := r.Gorm.Limit(limit).Offset(offset).Where("types like ?", types)

	if sort == 0 {
		tx.Order("num DESC")
	} else if sort == 1 {
		tx.Order("last_updatetime DESC")
	} else {
		return nil, fmt.Errorf("invalid sort: %d", sort)
	}

	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
