package server

import (
	"fmt"
	authpb "image-acquisition-service/api/grpc/auth/v1"
	interceptor "share/interceptor/micro/error"
	"share/micro/server"

	"github.com/micro/go-micro/v2"
)

var (
	ServiceName = "go.micro.api.image-acquisition.v1"
)

func NewImageAcquisitionServer(authHandler authpb.AuthServiceHandler) (micro.Service, error) {
	errInter, err := interceptor.NewErrorInterceptor()
	if err != nil {
		return nil, err
	}

	service, err := server.NewMicroServer(ServiceName)
	if err != nil {
		return nil, err
	}

	err = authpb.RegisterAuthServiceHandler(service.Server(), authHandler)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %v", err)
	}

	service.Init(micro.WrapHandler(errInter))

	return service, nil
}
