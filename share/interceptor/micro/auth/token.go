package interceptor

import (
	"context"
	"fmt"
	"strings"

	"comic/share/auth/token"

	"github.com/micro/go-micro/v2/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	bearerPrefix        = "Bearer "
	authorizationHeader = "authorization"
)

type Verifyer interface {
	VerifyToken(token string) (string, error)
}

type AuthInterceptor struct {
	verifyer Verifyer
}

func NewAuthInterceptor(pubKey string) (server.HandlerWrapper, error) {
	verifyer, err := token.NewJWTVerifier(pubKey)
	if err != nil {
		return nil, fmt.Errorf("invalid pub key: %v", err)
	}

	inter := &AuthInterceptor{
		verifyer: verifyer,
	}

	return inter.warpHandler, nil
}

func (this *AuthInterceptor) warpHandler(handler server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		tkn, err := tokenFromContext(ctx)
		if err != nil {
			return err
		}

		uid, err := this.verifyer.VerifyToken(tkn)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		return handler(contextWithUID(uid, ctx), req, rsp)
	}
}

// 从ctx中的md中拿取jwt
func tokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "null metadata")
	}

	bearers, ok := md[authorizationHeader]
	if !ok {
		return "", status.Error(codes.Unauthenticated, "authorization header?")
	}

	if len(bearers) == 0 {
		return "", status.Error(codes.Unauthenticated, "authorization header dont have value")
	}

	bearer := bearers[0]
	if !strings.HasPrefix(bearer, bearerPrefix) {
		return "", status.Error(codes.Unauthenticated, "invalid bearer prefix")
	}

	token := bearer[len(bearerPrefix):]

	return token, nil
}

type uidKey struct{}

// 将uid注入ctx中
func contextWithUID(uid string, ctx context.Context) context.Context {
	return context.WithValue(ctx, &uidKey{}, uid)
}

func UidFromContext(ctx context.Context) (string, error) {
	r := ctx.Value(&uidKey{})
	uid, ok := r.(string)
	if !ok {
		return "", fmt.Errorf("cannot take uid from context")
	}

	return uid, nil
}
