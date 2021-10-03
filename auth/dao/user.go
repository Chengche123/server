package dao

import (
	"auth-service/model"
	mgorm "share/database/gorm"
	"strconv"
	"time"

	xerrors "github.com/pkg/errors"

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
		return nil, xerrors.Wrapf(ErrDatabase, "open mysql, dsn = %s, database: [%v]", dsn, err)
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
				return xerrors.Wrapf(ErrInvalidPassword, "user_name = %s", userName)
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
			return xerrors.Wrapf(ErrDatabase, "database: [%v]", err)
		}

		return nil
	})

	return strconv.Itoa(int(userAccount.Id)), err
}
