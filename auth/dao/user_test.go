package dao

import (
	config "comic/share/config/database"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var testDsn string

func TestFindOrAddUser(t *testing.T) {
	db, err := NewUserRepository(config.DefaultMysqlDSN)
	if err != nil {
		panic(err.Error() + " dsn:" + testDsn)
	}

	defer func() {
		db.Close()
		time.Sleep(1 * time.Second)
	}()

	cases := []struct {
		name     string
		userName string
		password string
		want     string
		wantErr  bool
	}{
		{
			name:     "insert new user",
			userName: "chengche",
			password: "qwe1234",
			want:     "1",
			wantErr:  false,
		},
		{
			name:     "invalid password",
			userName: "chengche",
			password: "xxxxxx",
			want:     "",
			wantErr:  true,
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			accountId, err := db.FindOrAddUser(cs.userName, cs.password)

			if err != nil && !cs.wantErr {
				t.Error(err)
			}

			if err == nil && cs.wantErr {
				t.Errorf("want error,got no error")
			}

			if accountId != cs.want {
				t.Errorf("resolve user %q:%q,want %q,got %q", cs.userName, cs.password, cs.want, accountId)
			}
		})
	}
}

func TestMain(m *testing.M) {
	testDsn = "root:root@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	os.Exit(m.Run())
	// os.Exit(test.RunWithMysqlInDocker(m, &testDsn))
}
