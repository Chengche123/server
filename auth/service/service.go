package service

import (
	"auth-service/dao"
	"context"
	"log"
	"time"

	authpb "auth-service/api/grpc/v1"

	"github.com/pkg/errors"
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
	uid, err := s.UserRepository.FindOrAddUser(req.UserName, req.Password)

	if err != nil {

		switch errors.Cause(err) {

		case dao.ErrInvalidPassword:
			return status.Error(codes.Unauthenticated, "invalid password")

		case dao.ErrDatabase:
			log.Printf("stack trace:\n%+v\n", err) // 数据库错误跟踪
			return status.Error(codes.Internal, "database error")
		default:
			log.Printf("stack trace:\n%+v\n", err) // 未知错误跟踪
			return status.Error(codes.Internal, "unknown error")

		}

	}

	var expire time.Duration = time.Hour * 2

	tkn, err := s.TokenGenerator.GenerateToken(uid, expire)
	if err != nil {

		log.Printf("stack trace:\n%+v\n", err)
		return status.Error(codes.Unauthenticated, "token generate error")

	}

	res.AccessToken = tkn
	res.ExpiresIn = int32(expire.Seconds())

	return nil
}
