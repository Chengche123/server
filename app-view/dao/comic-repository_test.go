package dao

import (
	"context"
	"encoding/json"
	"share/micro/server"
	"testing"
)

func newComicRepository() (*ComicRepository, func() error) {
	service, err := server.NewMicroServer("test")
	if err != nil {
		panic(err)
	}

	repo, _ := NewComicRepository(service)

	return repo, service.Server().Stop
}

func TestComicRepository_ListBannerMo(t *testing.T) {
	repo, closer := newComicRepository()
	defer func() {
		_ = closer()
	}()

	mos, err := repo.ListBannerMo(context.TODO())
	if err != nil {
		t.Error(err)
		return
	}

	bs, _ := json.MarshalIndent(&mos, "", "  ")
	t.Log(string(bs))
}
