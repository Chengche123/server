package dao

import (
	"comic-service/model"
	"fmt"
)

func (r *ComicRepository) FindCategoryDetail(types string, sort, offset, limit int) ([]model.CategoryDetail, error) {
	var res []model.CategoryDetail

	types = "%" + types + "%"

	tx := r.Gorm.Limit(limit).Offset(offset).Where("types like ?", types)

	switch sort {
	case 0:
		tx.Order("num DESC")
	case 1:
		tx.Order("last_updatetime DESC")
	case 2:
		tx.Order("RAND()")
	default:
		return nil, fmt.Errorf("invalid sort: %d", sort)
	}

	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
