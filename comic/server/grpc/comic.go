package server

import (
	pb "comic-service/api/grpc/v1"
	"fmt"

	"comic/share/micro/server"

	key "comic/share/key"
	authInter "interceptor-micro/auth"
	errInter "interceptor-micro/error"

	"github.com/micro/go-micro/v2"
)

var (
	serverName = "go.micro.api.comic.comic.v1"
)

func NewComicServer(comicService pb.ComicServiceHandler) (micro.Service, error) {
	service, err := server.NewMicroServer(serverName)
	if err != nil {
		return nil, fmt.Errorf("cannot init server: %v", err)
	}

	authWrapper, err := authInter.NewAuthInterceptor(key.PubKey)
	if err != nil {
		return nil, fmt.Errorf("cannot init auth interceptor: %v", err)
	}

	errWrapper, err := errInter.NewErrorInterceptor()
	if err != nil {
		return nil, fmt.Errorf("cannot init error interceptor: %v", err)
	}

	service.Init(
		micro.WrapHandler(authWrapper, errWrapper),
	)

	err = pb.RegisterComicServiceHandler(service.Server(), comicService)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %v", err)
	}

	return service, nil
}
