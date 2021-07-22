package dao

import (
	"fmt"
	"news-service/model"
)

func (r *NewsRepository) FindCategoryDetail(tagId, sort, offset, limit int) ([]model.NewsCategoryDetail, error) {
	var res []model.NewsCategoryDetail

	tx := r.Gorm.Limit(limit).Offset(offset).Where("tag_id = ?", tagId)

	if sort == 0 {
		tx.Order("(moodamount+commentamount) DESC")
	} else if sort == 1 {
		tx.Order("createtime DESC")
	} else {
		return nil, fmt.Errorf("invalid sort: %d", sort)
	}

	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
