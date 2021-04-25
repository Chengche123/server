package main

import (
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"go.uber.org/zap"

	"comic/share/os/env"

	controler "rec-service/server/http"

	"os"
	"os/signal"
	"syscall"
)

var (
	registryAddr = env.FormatEnvOrDefault("%s", "COMIC_REGISTRY_ADDR", "127.0.0.1:2379")
)

func main() {
	logger, _ := zap.NewDevelopment()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{registryAddr}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.api.comic.rec.v1"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(5*time.Second),
	)

	service.Init()

	err := service.Server().Handle(
		service.Server().NewHandler(
			&controler.Service{},
		),
	)

	if err != nil {
		logger.Error("cannot register handler", zap.Error(err))
		return
	}

	go func() {
		err = service.Run()
		if err != nil {
			_ = service.Server().Stop()
			logger.Fatal("failed to run service", zap.Error(err))
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	<-termChan

	closed := make(chan struct{})
	go func() {
		err := service.Server().Stop()
		if err != nil {
			logger.Error("failed to stop service", zap.Error(err))
		}

		close(closed)
	}()

	select {
	case <-closed:
		logger.Info("graceful shutdown")
	case <-time.After(5 * time.Second):
		logger.Error("shutdown timeout")
	}
}
