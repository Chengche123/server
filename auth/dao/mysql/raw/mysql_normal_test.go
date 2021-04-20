package mysql

import (
	"context"
	"testing"
)

const mySqlTestURL = "root:root@tcp(127.0.0.1:3306)/comic"

func TestNewMysqlTable(t *testing.T) {
	tableName := "user_account"

	cases := []struct {
		name      string
		tableName string
		wantErr   bool
	}{
		{
			name:      "valid_table",
			tableName: tableName,
			wantErr:   false,
		},
		{
			name:      "invalid_table",
			tableName: tableName + tableName,
			wantErr:   true,
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			table, err := NewMysqlTable(context.Background(), mySqlTestURL, cs.tableName)
			defer func() {
				_ = table.Close()
			}()
			if err != nil && !cs.wantErr {
				t.Errorf("invalid table %s: %v", cs.tableName, err)
			}

			if err == nil && cs.wantErr {
				t.Errorf("want error,got no error; table: %s", cs.tableName)
			}
		})
	}
}

func TestResolveUser(t *testing.T) {
	tableName := "user_account"

	ctx := context.Background()

	table, err := NewMysqlTable(ctx, mySqlTestURL, tableName)
	if err != nil {
		t.Fatalf("failed to open mysql table: %v", err)
	}
	defer table.Close()

	accountID, err := table.ResolveUser(ctx, "chengche123", "qwe1234567")
	if err != nil {
		t.Fatalf("cannot resolve user: %v", err)
	}

	t.Log(accountID)
}
