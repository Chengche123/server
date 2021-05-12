package dao

import (
	"comic-service/model"
	"sync"
)

func (r *ComicRepository) FindComicDetails(ids []int64) ([]*model.ComicDetail, error) {
	var wg sync.WaitGroup
	inch := make(chan int64, len(ids))
	for _, id := range ids {
		inch <- id
	}
	close(inch)
	outch := make(chan *model.ComicDetail)

	concur := 16
	if len(ids) < concur {
		// 节约G
		concur = len(ids)
	}

	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for id := range inch {
				rec, err := r.findComicDetail(id)
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

	recs := make([]*model.ComicDetail, 0, len(ids))

	for rec := range outch {
		recs = append(recs, rec)
	}

	return recs, nil
}

func (r *ComicRepository) findComicDetail(id int64) (*model.ComicDetail, error) {
	var rec model.ComicDetail
	if err := r.Gorm.Where("id = ?", id).Take(&rec).Error; err != nil {
		return nil, err
	}

	return &rec, nil
}
