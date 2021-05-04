package dao

import (
	"comic-service/model"
	"sync"

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
	var wg sync.WaitGroup
	inch := make(chan int64, len(ids))
	for _, id := range ids {
		inch <- id
	}
	close(inch)
	outch := make(chan *model.Comic)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for id := range inch {
				rec, err := r.getComicInfo(id)
				if err != nil {
					continue
				}
				outch <- rec
			}

		}()
	}

	go func() {
		wg.Wait()
		close(outch)
	}()

	recs := make([]*model.Comic, 0, len(ids))

	for rec := range outch {
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
