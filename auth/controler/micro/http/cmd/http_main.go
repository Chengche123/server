package main

import (
	controler "micro/auth/http"
	"time"

	authpb "micro/auth/api/gen/v1"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"go.uber.org/zap"

	"os"
	"os/signal"
	"syscall"
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
		logger.Fatal("cannot register handler", zap.Error(err))
	}

	go func() {
		err = service.Run()
		if err != nil {
			service.Server().Stop()
			logger.Fatal("failed to run service", zap.Error(err))
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	<-termChan

	closed := make(chan struct{})
	go func() {
		service.Server().Stop()

		close(closed)
	}()

	select {
	case <-closed:
		logger.Info("graceful shutdown")
	case <-time.After(3 * time.Second):
		logger.Error("shutdown timeout")
	}
}
