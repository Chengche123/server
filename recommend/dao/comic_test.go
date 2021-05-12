package dao

import (
	config "comic/share/config/database"
	"testing"
)

func TestGetComicsByHot(t *testing.T) {
	repo, err := NewRecRepository(config.DefaultMysqlDSN)
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
