package server

import (
	pb "comic-service/api/grpc/v1"
	"fmt"

	"share/micro/server"

	"github.com/micro/go-micro/v2"
)

var (
	ServiceName = "go.micro.api.comic.comic.v1"
)

func NewComicServer(comicService pb.ComicServiceHandler) (micro.Service, error) {
	service, err := server.NewMicroServer(ServiceName)
	if err != nil {
		return nil, err
	}

	err = pb.RegisterComicServiceHandler(service.Server(), comicService)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %v", err)
	}

	service.Init()

	return service, nil
}
