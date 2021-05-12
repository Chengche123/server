package dao

import (
	"comic/share/database/gorm"
	"comic/share/database/mysql/dsn"
	"testing"

	"go.uber.org/zap"
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
	db, err := gorm.NewMysqlGormByDSN(dsn.DefaultDSN)
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewDevelopment()

	return &ComicRepository{
		Logger: logger,
		Gorm:   db,
	}
}
