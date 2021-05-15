package dao

import (
	"testing"
)

func TestComicRepository_FindComicCategoryFilter(t *testing.T) {
	repo := newComicRepository()
	defer repo.Close()

	res, err := repo.FindComicCategoryFilter()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res)
}
