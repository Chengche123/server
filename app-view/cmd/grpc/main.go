package main

import (
	"app-view/dao"
	grpcServer "app-view/server/grpc"
	grpcService "app-view/service/grpc"
	comicsrvpb "comic-service/api/grpc/v1"
	comicserver "comic-service/server/grpc"
	"os"
	"os/signal"
	zlog "share/log/zap"
	"syscall"

	"time"
)

var (
	graceTimeout = time.Second * 3
)

func main() {
	service := &grpcService.AppviewService{}

	server, err := grpcServer.NewAppViewServer(service)
	if err != nil {
		zlog.Logger.Error(err.Error())
		return
	}

	service.ComicRepository = &dao.ComicRepository{
		ComicService: comicsrvpb.NewComicService(comicserver.ServiceName, server.Client()),
	}

	go func() {
		err := server.Run()
		if err != nil {
			zlog.Logger.Error(err.Error())
			return
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	<-termChan

	closed := make(chan struct{})
	go func() {
		_ = server.Server().Stop()

		close(closed)
	}()

	select {
	case <-closed:
		zlog.Logger.Info("graceful shutdown")
	case <-time.After(graceTimeout):
		zlog.Logger.Error("shutdown timeout")
	}
}
