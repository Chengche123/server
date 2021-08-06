package dao

import (
	"comic-service/model"
	"fmt"
)

func (r *ComicRepository) FindComicChapters(id int32) ([]model.ComicChapter, error) {
	var rs []model.ComicChapter

	if err := r.DB.Where("comic_id = ?", id).Order("chapterorder DESC").Find(&rs).Error; err != nil {
		return nil, fmt.Errorf("cannot find comic chapter: %v", err)
	}

	return rs, nil
}
