package service

import (
	"context"
	"fmt"
	pb "image-acquisition-service/api/grpc/auth/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository interface {
	AddUser(username, password string) (accountID string, err error)
	FindUser(username, password string) (accountID string, err error)
}

type TokenGenerator interface {
	GenerateToken(accountID string, expireIn time.Duration) (string, error)
}

type UserService struct {
	UserRepository UserRepository
	TokenGenerator TokenGenerator
}

func check(username, password string) error {
	if username == "" || password == "" {
		return fmt.Errorf("用户名或密码不能为空")
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest, res *pb.LoginResponse) error {
	err := check(req.UserName, req.Password)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	uid, err := s.UserRepository.FindUser(req.UserName, req.Password)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	var expire time.Duration = time.Hour * 72

	tkn, err := s.TokenGenerator.GenerateToken(uid, expire)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	res.AccessToken = tkn
	res.ExpiresIn = int32(expire.Seconds())

	return nil
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest, res *pb.RegisterResponse) error {
	err := check(req.UserName, req.Password)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	uid, err := s.UserRepository.AddUser(req.UserName, req.Password)
	if err != nil {
		return status.Error(codes.AlreadyExists, err.Error())
	}

	var expire time.Duration = time.Hour * 72

	tkn, err := s.TokenGenerator.GenerateToken(uid, expire)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	res.AccessToken = tkn
	res.ExpiresIn = int32(expire.Seconds())

	return nil
}
