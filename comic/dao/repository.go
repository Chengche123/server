package dao

import "comic-service/model"

type ComicRepository struct {
}

func (*ComicRepository) GetComicInfo(ids []int64) ([]model.ComicInfo, error) {

	return nil, nil
}
