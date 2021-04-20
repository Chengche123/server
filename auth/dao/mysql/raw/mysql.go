package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mySqlTable struct {
	name string
	db   *sqlx.DB
}

func (m *mySqlTable) Close() error {
	return m.db.Close()
}

func (m *mySqlTable) ResolveUser(ctx context.Context, userName, password string) (string, error) {
	tx, _ := m.db.Begin()
	row := tx.QueryRow(fmt.Sprintf("SELECT id FROM %s WHERE user_name = ? AND password = ?;", m.name), userName, password)
	var accountID int
	err := row.Scan(&accountID)
	if err == sql.ErrNoRows {
		if err := tx.QueryRow(fmt.Sprintf("SELECT id FROM %s WHERE user_name = ?", m.name), userName).Err(); err != sql.ErrNoRows {
			_ = tx.Rollback()
			return "", fmt.Errorf("user alreay exist")
		}

		res, err := tx.Exec(fmt.Sprintf("INSERT INTO %s(user_name,password) VALUES(?,?);", m.name), userName, password)
		if err != nil {
			_ = tx.Rollback()
			return "", fmt.Errorf("cannot insert a new user: %v", err)
		}

		increaseID, _ := res.LastInsertId()
		accountID = int(increaseID)
	}

	err = tx.Commit()
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return strconv.Itoa(accountID), nil
}

func NewMysqlTable(ctx context.Context, cstr string, tableName string) (*mySqlTable, error) {
	db, err := sqlx.Open("mysql", cstr)
	if err != nil {
		return nil, fmt.Errorf("cannot open mysql connect: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot connect mysql database: %v", err)
	}

	_, err = db.Exec(fmt.Sprintf("SELECT EXISTS(SELECT * FROM %s)", tableName))
	if err != nil {
		return nil, fmt.Errorf("invalid table: %v", err)
	}

	return &mySqlTable{
		db:   db,
		name: tableName,
	}, nil
}
