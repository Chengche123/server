package dao

import (
	"encoding/json"
	"testing"
)

func TestComicRepository_FindComicSpecial(t *testing.T) {
	repo := newComicRepository()
	defer repo.Close()

	res, err := repo.FindComicSpecial(0, 5)
	if err != nil {
		t.Error(err)
		return
	}

	bs, _ := json.MarshalIndent(&res, "", "  ")
	t.Log(string(bs))
}
