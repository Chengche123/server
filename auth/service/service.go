package service

import (
	"context"
	zlog "share/log/zap"
	"time"

	authpb "auth-service/api/grpc/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository interface {
	FindOrAddUser(userName, password string) (accountID string, err error)
}

type TokenGenerator interface {
	GenerateToken(accountID string, expireIn time.Duration) (string, error)
}

type AuthService struct {
	TokenGenerator TokenGenerator
	UserRepository UserRepository
}

func (*AuthService) Login(ctx context.Context, req *authpb.LoginRequest, res *authpb.LoginResponse) error {
	return status.Error(codes.Unimplemented, "")
}

func (s *AuthService) UserLogin(ctx context.Context, req *authpb.UserLoginRequest, res *authpb.LoginResponse) error {
	zlog.Logger.Info("loggin", zap.String("user_name", req.UserName), zap.String("password", req.Password))

	uid, err := s.UserRepository.FindOrAddUser(req.UserName, req.Password)
	if err != nil {
		return status.Error(codes.Unauthenticated, "")
	}

	var expire time.Duration = time.Hour * 2

	tkn, err := s.TokenGenerator.GenerateToken(uid, expire)
	if err != nil {
		return status.Error(codes.Unauthenticated, "")
	}

	res.AccessToken = tkn
	res.ExpiresIn = int32(expire.Seconds())

	return nil
}
