package dao

import (
	"errors"
)

var (
	ErrInvalidPassword = errors.New("auth-service/dao: invalid password")
	ErrDatabase        = errors.New("auth-service/dao: database error")
)
