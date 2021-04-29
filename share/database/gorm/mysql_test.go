package gorm

import "testing"

func TestNewMysqlGormByDSN(t *testing.T) {
	db, err := NewMysqlGormByDSN("root:root@tcp(127.0.0.1:3306)/INVALID")
	if err == nil {
		t.Errorf("want error,got no error")
		raw, _ := db.DB()
		raw.Close()
	}

	db, err = NewMysqlGormByDSN("root:root@tcp(127.0.0.1:3306)/comic")
	if err != nil {
		t.Errorf("cannot create gorm object: %v", err)
	} else {
		raw, _ := db.DB()
		raw.Close()
	}
}
