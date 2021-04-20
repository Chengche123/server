package main

import (
	controler "micro/auth/http"
	"time"

	authpb "micro/auth/api/gen/v1"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://127.0.0.1:2379"}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.api.comic.auth.v1"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(10*time.Second),
	)

	service.Init()

	err := service.Server().Handle(
		service.Server().NewHandler(
			&controler.Service{
				Grpc:   authpb.NewAuthService("go.micro.srv.comic.auth.v1", service.Client()),
				Looger: logger,
			},
		),
	)
	if err != nil {
		panic(err)
	}

	err = service.Run()
	if err != nil {
		panic(err)
	}
}
