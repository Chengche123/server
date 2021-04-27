package dao

import (
	"testing"

	"go.uber.org/zap"
)

func TestGetComicsByHot(t *testing.T) {
	logger, _ := zap.NewDevelopment()

	repo, err := NewRecRepositoryByDSN("root:root@tcp(127.0.0.1:3306)/comic", logger)
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	comics, err := repo.GetComicsByHot(0*5, 5)
	if err != nil {
		panic(err)
	}
	t.Logf("%+v\n", comics)

	comics, err = repo.GetComicsByHot(1*5, 5)
	if err != nil {
		panic(err)
	}
	t.Logf("%+v\n", comics)
}
