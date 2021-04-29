package dao

import (
	"comic-service/model"

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

func (r *ComicRepository) GetComicInfos(ids []int64) ([]*model.Comic, error) {
	recs := make([]*model.Comic, 0, len(ids))

	for _, id := range ids {
		rec, err := r.getComicInfo(id)
		if err != nil {
			continue
		}

		recs = append(recs, rec)
	}

	return recs, nil
}

func (r *ComicRepository) getComicInfo(id int64) (*model.Comic, error) {
	var rec model.Comic
	if err := r.Gorm.Where("id = ?", id).Take(&rec).Error; err != nil {
		return nil, err
	}

	return &rec, nil
}
