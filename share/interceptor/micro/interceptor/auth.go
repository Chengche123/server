package interceptor

import (
	"context"

	"google.golang.org/grpc"

	"github.com/micro/go-micro/v2/server"
)

type authInterceptor struct {
	grpcInterceptor grpc.UnaryServerInterceptor
}

func NewAuthInterceptor(grpcInterceptor grpc.UnaryServerInterceptor) server.HandlerWrapper {
	a := &authInterceptor{
		grpcInterceptor: grpcInterceptor,
	}
	return a.wrapHandler
}

func (i *authInterceptor) wrapHandler(microHandler server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		// 空handler,不在grpc中间件的handler里面调用microHandler
		grpcHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return nil, nil
		}

		// 让microHandler中的ctx经过grpc中间件
		_, err := i.grpcInterceptor(ctx, req.Body(), nil, grpcHandler)
		if err != nil {
			return err
		}

		return microHandler(ctx, req, rsp)
	}
}
