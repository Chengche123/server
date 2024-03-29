package dao

import (
	config "share/config/database"

	"testing"
)

func TestGetComicInfos(t *testing.T) {
	ins := newComicRepository()
	defer ins.Close()

	ids := []int64{50758, 33322}
	res, err := ins.FindComicDetails(ids)
	if err != nil {
		panic(err)
	}

	if len(res) != 2 {
		t.Errorf("len(res) want 2,got %v", len(res))
	}
}

func newComicRepository() *ComicRepository {
	db, err := NewComicRepository(config.DefaultMysqlDSN)
	if err != nil {
		panic(err)
	}

	return db
}
