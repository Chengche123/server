package dao

import (
	"auth-service/model"
	"fmt"
	mgorm "share/database/gorm"
	zlog "share/log/zap"
	"strconv"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewUserRepository(dsn string) (*UserRepository, error) {
	db, err := mgorm.NewMysqlGormByDSN(dsn)
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (m *UserRepository) FindOrAddUser(userName, password string) (accountID string, err error) {
	var userAccount model.UserAccount

	err = m.db.Transaction(func(tx *gorm.DB) error {
		// 覆盖索引,不回表
		tx.Where("user_name = ?", userName).Select("user_name", "password", "id").Take(&userAccount)

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
		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		userAccount.UserName = userName
		userAccount.Password = string(hashedPwd)
		userAccount.AddTime = time.Now().Unix()
		userAccount.Status = 1
		if err := tx.Create(&userAccount).Error; err != nil {
			return fmt.Errorf("cannot insert a row: %v", err)
		}

		return nil
	})

	if err != nil {
		return "", fmt.Errorf("cannot resolve user: %v", err)
	}

	zlog.Logger.Info("", zap.Int("accountID", int(userAccount.Id)))
	return strconv.Itoa(int(userAccount.Id)), nil
}
