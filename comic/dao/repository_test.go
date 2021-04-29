package dao

import (
	"comic/share/database/gorm"
	"testing"

	"go.uber.org/zap"
)

func TestGetComicInfos(t *testing.T) {
	ins := newComicRepository()
	defer ins.Close()

	ids := []int64{1, 2, 3, 4, 3826, 3956}
	res, err := ins.GetComicInfos(ids)
	if err != nil {
		panic(err)
	}

	if len(res) != 2 {
		t.Errorf("len(res) want 2,got %v", len(res))
	}
}

func newComicRepository() *ComicRepository {
	db, err := gorm.NewMysqlGormByDSN("root:root@tcp(127.0.0.1:3306)/comic")
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewDevelopment()

	return &ComicRepository{
		Logger: logger,
		Gorm:   db,
	}
}
