package dao

import (
	"database/sql"
	"testing"

	_ "gorm.io/driver/mysql"

	"rec-service/model"
)

func TestRepository(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/comic")
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewRecRepository(db)
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	var rec model.Comic
	rec.Id = 3826
	repo.db.First(&rec)
	if rec.AddTime != 1618585088 ||
		rec.Authors != "远山绘麻" ||
		rec.Cover != "https://images.dmzj1.com/webpic/1/160202yaoniduiwofml.jpg" ||
		rec.Id != 3826 ||
		rec.IsEnd != 1 ||
		rec.LastUpdatetime != 1440042922 ||
		rec.Num != 28424093 ||
		rec.Status != 1 ||
		rec.Title != "绝对恋爱命令" ||
		rec.Types != "爱情" {
		panic(err)
	}
	t.Logf("%+v\n", rec)
}
