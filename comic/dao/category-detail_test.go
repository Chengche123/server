package dao

import (
	"encoding/json"
	"testing"
)

func TestComicRepository_FindCategoryDetail(t *testing.T) {
	repo := newComicRepository()
	defer repo.Close()

	res, err := repo.FindCategoryDetail("热血", 0, 0, 5)
	if err != nil {
		t.Error(err)
		return
	}

	jbs, err := json.MarshalIndent(&res, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(jbs))
}
