package dao

import (
	"testing"
)

func TestComicRepository_FindComicChapters(t *testing.T) {
	r := newComicRepository()
	defer r.Close()

	rs, err := r.FindComicChapters(33322)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(rs))
}
