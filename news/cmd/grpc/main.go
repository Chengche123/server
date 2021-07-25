package main

import (
	"news-service/dao"
	server "news-service/server/grpc"
	nservice "news-service/service"
	"os"
	"os/signal"
	config "share/config/database"
	zlog "share/log/zap"
	"share/os/env"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var (
	mysqlDBAddr  = env.FormatEnvOrDefault("%s", "COMIC_MYSQL_DSN", config.DefaultMysqlDSN)
	graceTimeout = time.Second * 3
)

func main() {
	repo, err := dao.NewNewsRepository(mysqlDBAddr)
	if err != nil {
		zlog.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	server, err := server.NewNewsServer(&nservice.NewsService{
		NewsRepository: repo,
	})
	if err != nil {
		zlog.Logger.Error("", zap.Error(err))
		return
	}

	go func() {
		err := server.Run()
		if err != nil {
			zlog.Logger.Error("news grpc server has stoped", zap.Error(err))
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
