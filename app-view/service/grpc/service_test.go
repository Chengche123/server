package service

import (
	v1 "app-view/api/grpc/v1"
	"app-view/dao"
	"context"
	"encoding/json"
	"share/micro/server"
	"testing"
)

func TestAppviewService_ListHomeMo(t *testing.T) {
	repo, closer := newComicRepository()
	defer func() { _ = closer() }()

	service := &AppviewService{
		ComicRepository: repo,
	}

	var res v1.ListHomeMoResponse
	err := service.ListHomeMo(context.Background(), &v1.ListHomeMoRequest{}, &res)
	if err != nil {
		t.Error(err)
		return
	}

	bs, _ := json.MarshalIndent(&res, "", "  ")
	t.Log(string(bs))
}

func newComicRepository() (*dao.ComicRepository, func() error) {
	service, err := server.NewMicroServer("go.micro.api.test")
	if err != nil {
		panic(err)
	}

	repo, _ := dao.NewComicRepository(service)

	return repo, service.Server().Stop
}
