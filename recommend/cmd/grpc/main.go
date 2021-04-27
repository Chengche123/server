package main

import (
	"time"

	"go.uber.org/zap"

	"rec-service/dao"
	grpc "rec-service/service"

	"comic/share/os/env"

	"os"
	"os/signal"
	server "rec-service/server/grpc"
	"syscall"
)

var (
	mysqlDBAddr = env.FormatEnvOrDefault("root:root@tcp(%s)/comic", "COMIC_MYSQL_ADDR", "127.0.0.1:3306")
)

func main() {
	logger, _ := zap.NewDevelopment()

	repo, err := dao.NewRecRepositoryByDSN(mysqlDBAddr, logger)
	if err != nil {
		logger.Fatal("failed to create repository", zap.Error(err))
	}

	service, err := server.NewRecServer(&grpc.Service{
		Logger:        logger,
		RecRepository: repo,
	})
	if err != nil {
		logger.Error("failed to create server", zap.Error(err))
		return
	}

	go func() {
		err = service.Run()
		if err != nil {
			logger.Error("falied to run service", zap.Error(err))
			return
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
	case <-time.After(3 * time.Second):
		logger.Error("shutdown timeout")
	}
}
