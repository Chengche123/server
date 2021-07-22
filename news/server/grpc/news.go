package server

import (
	"fmt"
	pb "news-service/api/grpc/v1"
	"share/micro/server"

	"github.com/micro/go-micro/v2"
)

var (
	ServiceName = "go.micro.api.comic.news.v1"
)

func NewNewsServer(newsService pb.NewsServiceHandler) (micro.Service, error) {
	service, err := server.NewMicroServer(ServiceName)
	if err != nil {
		return nil, err
	}

	err = pb.RegisterNewsServiceHandler(service.Server(), newsService)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %v", err)
	}

	service.Init()

	return service, nil
}
