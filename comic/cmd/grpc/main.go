package main

import (
	"comic-service/dao"
	server "comic-service/server/grpc"
	grpc "comic-service/service/grpc"
	"comic/share/database/gorm"
	"comic/share/os/env"

	"go.uber.org/zap"

	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	mysqlDBAddr  = env.FormatEnvOrDefault("root:root@tcp(%s)/comic", "COMIC_MYSQL_ADDR", "127.0.0.1:3306")
	graceTimeout = time.Second * 3
)

func main() {
	logger, _ := zap.NewDevelopment()

	db, err := gorm.NewMysqlGormByDSN(mysqlDBAddr)
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}

	repo := &dao.ComicRepository{
		Gorm:   db,
		Logger: logger,
	}

	service := &grpc.ComicService{
		Logger:     logger,
		Repository: repo,
	}

	server, err := server.NewComicServer(service)
	if err != nil {
		logger.Error("", zap.Error(err))
	}

	go func() {
		err := server.Run()
		if err != nil {
			logger.Error("service has stoped", zap.Error(err))
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	<-termChan

	closed := make(chan struct{})
	go func() {
		err := server.Server().Stop()
		if err != nil {
			logger.Error("failed to stop service", zap.Error(err))
		}

		close(closed)
	}()

	select {
	case <-closed:
		logger.Info("graceful shutdown")
	case <-time.After(graceTimeout):
		logger.Error("shutdown timeout")
	}
}
