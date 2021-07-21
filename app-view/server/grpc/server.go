package server

import (
	pb "app-view/api/grpc/v1"
	"fmt"
	interceptor "share/interceptor/micro/error"

	"share/micro/server"

	"github.com/micro/go-micro/v2"
)

var (
	ServiceName = "go.micro.api.comic.appview.v1"
)

func NewAppViewServer(appviewService pb.AppviewServiceHandler) (micro.Service, error) {
	service, err := server.NewMicroServer(ServiceName)
	if err != nil {
		return nil, err
	}

	err = pb.RegisterAppviewServiceHandler(service.Server(), appviewService)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %v", err)
	}

	errInter, err := interceptor.NewErrorInterceptor()
	if err != nil {
		return nil, err
	}

	service.Init(micro.WrapHandler(errInter))

	return service, nil
}
