package dao

import (
	"fmt"
	"image-acquisition-service/model"
	mgorm "share/database/gorm"
	"strconv"

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

func (r *UserRepository) AddUser(username, password string) (accountID string, err error) {
	var user model.User

	r.db.Where("user_name = ?", username).Take(&user)

	if user.Id != 0 {
		return "", fmt.Errorf("user alreay existed")
	}

	user.UserName = username
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashedPwd)
	if err := r.db.Create(&user).Error; err != nil {
		return "", fmt.Errorf("failed to insert user: %v", err)
	}

	return strconv.Itoa(int(user.Id)), nil
}

func (r *UserRepository) FindUser(username, password string) (accountID string, err error) {
	var user model.User

	r.db.Where("user_name = ?", username).Take(&user)

	if user.Id == 0 {
		return "", fmt.Errorf("user not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	return strconv.Itoa(int(user.Id)), nil
}
