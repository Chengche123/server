package main

import (
	"comic-service/dao"
	server "comic-service/server/grpc"
	grpc "comic-service/service/grpc"
	config "share/config/database"
	zlog "share/log/zap"

	"share/os/env"

	"go.uber.org/zap"

	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	mysqlDBAddr  = env.FormatEnvOrDefault("root:root@tcp(%s)/comic", "COMIC_MYSQL_ADDR", config.DefaultMySqlAddr)
	graceTimeout = time.Second * 3
)

func main() {
	repo, err := dao.NewComicRepository(mysqlDBAddr)
	if err != nil {
		zlog.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	server, err := server.NewComicServer(&grpc.ComicService{
		ComicRepository: repo,
	})
	if err != nil {
		zlog.Logger.Error("", zap.Error(err))
		return
	}

	go func() {
		err := server.Run()
		if err != nil {
			zlog.Logger.Error("service has stoped", zap.Error(err))
			return
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	<-termChan

	closed := make(chan struct{})
	go func() {
		_ = repo.Close()
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
