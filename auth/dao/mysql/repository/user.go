package repository

import (
	"comic/auth/dao/mysql/models"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MySqlTable struct {
	db *gorm.DB
}

func (m *MySqlTable) Close() error {
	rawDb, err := m.db.DB()
	if err != nil {
		return fmt.Errorf("cannot close db connect: %v", err)
	}
	return rawDb.Close()
}

func (m *MySqlTable) FindOrAddUser(userName, password string) (accountID string, err error) {
	var userAccount models.UserAccount
	err = m.db.Transaction(func(tx *gorm.DB) error {
		tx.Where("user_name = ?", userName).First(&userAccount)

		// 找到用户
		if userAccount.Id != 0 {
			// 匹配密码
			err := bcrypt.CompareHashAndPassword([]byte(userAccount.Password), []byte(password))
			if err != nil {
				return fmt.Errorf("invalid password")
			}

			return nil
		}

		// 未找到用户,创建新用户
		userAccount.UserName = userName
		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		userAccount.Password = string(hashedPwd)
		userAccount.AddTime = int(time.Now().Unix())
		userAccount.Status = 1
		if err := tx.Create(&userAccount).Error; err != nil {
			return fmt.Errorf("cannot insert a row: %v", err)
		}

		return nil
	})

	if err != nil {
		return "", fmt.Errorf("cannot resolve user: %v", err)
	}

	return strconv.Itoa(int(userAccount.Id)), nil
}

func NewMySqlTable(ctx context.Context, raw *sql.DB) (*MySqlTable, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: raw,
	}), &gorm.Config{
		// Logger:                 logger.Default.LogMode(logger.Info),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cannot open db connect: %v", err)
	}

	if !db.Migrator().HasTable("user_account") {
		err := db.AutoMigrate(&models.UserAccount{})
		if err != nil {
			return nil, fmt.Errorf("cannot create table: %v", err)
		}
	}

	return &MySqlTable{
		db: db,
	}, nil
}
