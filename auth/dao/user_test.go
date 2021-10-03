package dao

import (
	config "share/config/database"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func TestUserRepository_FindOrAddUser(t *testing.T) {
	db, err := NewUserRepository(config.DefaultMysqlDSN)
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	defer db.Close()

	type args struct {
		userName string
		password string
	}
	tests := []struct {
		name          string
		m             *UserRepository
		args          args
		wantAccountID string
		wantErr       error
	}{
		{
			"invalid password",
			db,
			args{"newuser1", "newuser0"},
			"",
			bcrypt.ErrMismatchedHashAndPassword,
		},
		{
			"correct password",
			db,
			args{"newuser1", "newuser1"},
			"157",
			nil,
		},
		{
			"new user",
			db,
			args{"newuser2", "newuser2"},
			"158",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAccountID, err := tt.m.FindOrAddUser(tt.args.userName, tt.args.password)

			if errors.Cause(err) != tt.wantErr {
				t.Errorf("UserRepository.FindOrAddUser() error = [%v], wantErr [%v]", errors.Cause(err), errors.Cause(tt.wantErr))
				t.Errorf("stack trace:\n %+v\n", err)
				return
			}

			if gotAccountID != tt.wantAccountID {
				t.Errorf("UserRepository.FindOrAddUser() = %v, want %v", gotAccountID, tt.wantAccountID)
			}
		})
	}
}

func TestNewUserRepository(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			"correct dsn",
			args{"root:root@tcp(127.0.0.1:3306)/comic"},
			nil,
		},
		{
			"invalid dsn",
			args{"root:root@tcp(127.0.0.1:3306)/comic1"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserRepository(tt.args.dsn)
			if errors.Cause(err) != tt.wantErr {
				t.Errorf("NewUserRepository() errorType = [%T], error = [%v], wantErr [%v]",
					errors.Cause(err), errors.Cause(err), tt.wantErr)

				t.Errorf("stack trace:\n%+v\n", err)
				return
			}
			got.Close()
		})
	}
}
