package server

import (
	pb "auth-service/api/grpc/v1"
	interceptor "comic/share/interceptor/micro/error"
	"fmt"

	"comic/share/micro/server"

	"github.com/micro/go-micro/v2"
)

var (
	serverName = "go.micro.api.comic.auth.v1"
)

func NewAuthServer(authService pb.AuthServiceHandler) (micro.Service, error) {
	errInter, err := interceptor.NewErrorInterceptor()
	if err != nil {
		return nil, err
	}

	service, err := server.NewMicroServer(serverName)
	if err != nil {
		return nil, err
	}

	err = pb.RegisterAuthServiceHandler(service.Server(), authService)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %v", err)
	}

	service.Init(micro.WrapHandler(errInter))

	return service, nil
}
