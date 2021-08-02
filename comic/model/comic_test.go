package model

import (
	"encoding/json"
	"testing"

	fuzz "github.com/google/gofuzz"
)

func TestFoo(t *testing.T) {
	var obj ComicChapter

	fuzz.New().Fuzz(&obj)

	bs, _ := json.MarshalIndent(&obj, "", "  ")
	t.Log(string(bs))
}
