package controler

import (
	"context"
	"errors"
	"io"
	"time"

	authpb "micro/auth/api/gen/v1"

	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserBo interface {
	GetToken() string
	GetExpireIn() time.Duration
}

type UserResolver interface {
	ResolveUser(ctx context.Context, userName, password string) (userBo interface{}, err error)
}

type Service struct {
	UserResolver UserResolver
	Logger       *zap.Logger
}

func (s *Service) Close() error {
	if closer, ok := s.UserResolver.(io.Closer); ok {
		err := closer.Close()
		if err != nil {
			s.Logger.Error("failed to close UserResolver", zap.Error(err))
		}
	}

	return nil
}

func (*Service) Login(ctx context.Context, req *authpb.LoginRequest, res *authpb.LoginResponse) error {
	res.AccessToken = req.Code
	res.ExpiresIn = 7200
	return nil
}

func (s *Service) UserLogin(ctx context.Context, req *authpb.UserLoginRequest, res *authpb.LoginResponse) error {
	s.Logger.Info("loggin", zap.String("user_name", req.UserName), zap.String("password", req.Password))

	userBo, err := s.UserResolver.ResolveUser(ctx, req.UserName, req.Password)
	if err != nil {
		return status.Error(codes.Unauthenticated, "密码错误")
	}

	IUserBo, ok := userBo.(UserBo)
	if !ok {
		s.Logger.Error("invalid assert", zap.Error(errors.New("userBo.(UserBo)")))
		return status.Error(codes.Internal, "")
	}

	res.AccessToken = IUserBo.GetToken()
	res.ExpiresIn = int32(IUserBo.GetExpireIn().Seconds())

	return nil
}
