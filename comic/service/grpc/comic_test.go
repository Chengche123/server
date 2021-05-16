package service

import (
	v1 "comic-service/api/grpc/v1"
	"comic-service/dao"
	"context"
	"encoding/json"
	config "share/config/database"
	"testing"
)

func TestComicService_ListComicCategoryFilter(t *testing.T) {
	service := newComicService()
	defer service.ComicRepository.Close()

	res := &v1.ListComicCategoryFilterResponse{}
	err := service.ListComicCategoryFilter(context.Background(), &v1.ListComicCategoryFilterRequest{}, res)
	if err != nil {
		t.Error(err)
		return
	}

	bs, _ := json.MarshalIndent(&res, "", "  ")
	t.Log(string(bs))
}

func newComicService() *ComicService {
	repo, err := dao.NewComicRepository(config.DefaultMysqlDSN)
	if err != nil {
		panic(err)
	}

	return &ComicService{
		ComicRepository: repo,
	}
}
