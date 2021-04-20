package controler

import (
	authpb "comic/auth/controler/grpc/api/gen/v1"
	"context"
	"errors"
	"io"
	"time"

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

type OpenIDResolver interface {
	ResolveOpenID(ctx context.Context, openID string) (string, error)
}

type CodeResolver interface {
	ResolveCode(code string) (openID string, err error)
}

type TokenGenerator interface {
	GenerateToken(accountID string, expireIn time.Duration) (string, error)
}

// Service implements auth service
type service struct {
	logger *zap.Logger

	codeResolver   CodeResolver
	openIDResolver OpenIDResolver

	tokenExpireIn  time.Duration
	tokenGenerator TokenGenerator

	userResolver UserResolver

	authpb.UnimplementedAuthServiceServer
}

func (s *service) Close() error {
	if closer, ok := s.userResolver.(io.Closer); ok {
		err := closer.Close()
		if err != nil {
			s.logger.Error("cannot close userResolver", zap.Error(err))
		}

	}
	return nil
}

// Login receive front request of login
func (s *service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	openid, err := s.codeResolver.ResolveCode(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolver openid: %v", err)
	}

	s.logger.Info("receive code", zap.String("code", req.Code))
	accountID, err := s.openIDResolver.ResolveOpenID(c, openid)
	if err != nil {
		s.logger.Error("failed to ResolveOpenID", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	token, err := s.tokenGenerator.GenerateToken(accountID, s.tokenExpireIn)
	if err != nil {
		s.logger.Error("failed to generate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: token,
		ExpiresIn:   int32(s.tokenExpireIn.Seconds()),
	}, nil
}

func (s *service) UserLogin(ctx context.Context, req *authpb.UserLoginRequest) (*authpb.LoginResponse, error) {
	userBo, err := s.userResolver.ResolveUser(ctx, req.UserName, req.Password)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "密码错误")
	}

	IUserBo, ok := userBo.(UserBo)
	if !ok {
		s.logger.Error("invalid assert", zap.Error(errors.New("userBo.(UserBo)")))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: IUserBo.GetToken(),
		ExpiresIn:   int32(IUserBo.GetExpireIn().Seconds()),
	}, nil
}
