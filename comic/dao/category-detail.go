package dao

import (
	"comic-service/model"
	"fmt"
	zlog "share/log/zap"
	"sync"

	"go.uber.org/zap"
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

func (r *ComicRepository) FindCategoryComicDetail(ids []int32) ([]*model.CategoryDetail, error) {
	res := make([]*model.CategoryDetail, 0, len(ids))

	inch := make(chan int32, len(ids))
	for _, v := range ids {
		inch <- v
	}
	close(inch)
	outch := make(chan *model.CategoryDetail, 3)

	var wg sync.WaitGroup
	concur := 8
	if concur > len(ids) {
		concur = len(ids)
	}

	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for v := range inch {
				rmo, err := r.findCategoryComicDetail(v)
				if err != nil {
					zlog.Logger.Info("cannot find category comic detail", zap.Error(err), zap.Int32("comic_id", v))
					continue
				}
				outch <- rmo
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outch)
	}()

	for v := range outch {
		res = append(res, v)
	}

	return res, nil
}

func (r *ComicRepository) findCategoryComicDetail(id int32) (*model.CategoryDetail, error) {
	var res model.CategoryDetail

	if err := r.DB.Where("id = ?", id).Take(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
