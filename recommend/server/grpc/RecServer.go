package server

import (
	"fmt"
	pb "rec-service/api/grpc/v1"
	"share/micro/server"

	"github.com/micro/go-micro/v2"
)

var (
	srvName = "go.micro.api.comic.rec.v1"
)

func NewRecServer(recService pb.RecServiceHandler) (micro.Service, error) {
	service, err := server.NewMicroServer(srvName)
	if err != nil {
		return nil, err
	}

	err = pb.RegisterRecServiceHandler(service.Server(), recService)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %v", err)
	}

	service.Init()

	return service, nil
}
