package service

import (
	"context"
	"fmt"
	"io"
	"time"

	"go.uber.org/zap"
)

type UserDomain struct {
	token    string
	expireIn time.Duration
}

func NewUserDomain(accountID string, expireIn time.Duration) (*UserDomain, error) {
	return nil, nil
}

func (u *UserDomain) GetToken() string {
	return u.token
}

func (u *UserDomain) GetExpireIn() time.Duration {
	return u.expireIn
}

type TokenGenerator interface {
	GenerateToken(accountID string, expireIn time.Duration) (string, error)
}

type UserRepository interface {
	FindOrAddUser(userName, password string) (accountID string, err error)
}

type UserService struct {
	UserRepository UserRepository
	TokenGenerator TokenGenerator
	Logger         *zap.Logger
}

func (s *UserService) Close() error {
	if closer, ok := s.UserRepository.(io.Closer); ok {
		err := closer.Close()
		if err != nil {
			s.Logger.Error("failed to close user repository", zap.Error(err))
		}
	}

	return nil
}

func (s *UserService) ResolveUser(ctx context.Context, userName, password string) (userBo interface{}, err error) {
	accountId, err := s.UserRepository.FindOrAddUser(userName, password)
	if err != nil {
		return nil, fmt.Errorf("cannot find or add user: %v", err)
	}

	expireIn := time.Hour * 2
	tkn, err := s.TokenGenerator.GenerateToken(accountId, expireIn)
	if err != nil {
		return nil, fmt.Errorf("cannot generate token: %v", err)
	}

	return &UserDomain{
		token:    tkn,
		expireIn: expireIn,
	}, nil
}
