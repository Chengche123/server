package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type authInterceptor struct {
}

func NewAuthInterceptor() grpc.UnaryServerInterceptor {
	a := &authInterceptor{}

	return a.wrapHandler
}

func (*authInterceptor) wrapHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("call begin")
	return handler(ctx, req)
}
