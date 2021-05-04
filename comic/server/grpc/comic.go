package server

import (
	pb "comic-service/api/grpc/v1"
	"fmt"

	"comic/share/micro/server"

	"github.com/micro/go-micro/v2"
)

var (
	serverName = "go.micro.srv.comic.comic.v1"
)

func NewComicServer(comicService pb.ComicServiceHandler) (micro.Service, error) {
	service, err := server.NewMicroServer(serverName)
	if err != nil {
		return nil, fmt.Errorf("cannot init server: %v", err)
	}

	service.Init()

	err = pb.RegisterComicServiceHandler(service.Server(), comicService)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %v", err)
	}

	return service, nil
}
