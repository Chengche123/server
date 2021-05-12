package main

import (
	config "comic/share/config/database"
	zlog "comic/share/log/zap"
	"time"

	"go.uber.org/zap"

	"rec-service/dao"
	"rec-service/service"

	"comic/share/os/env"

	"os"
	"os/signal"
	server "rec-service/server/grpc"
	"syscall"
)

var (
	mysqlDBAddr = env.FormatEnvOrDefault("root:root@tcp(%s)/comic", "COMIC_MYSQL_ADDR", config.DefaultMySqlAddr)
)

func main() {
	repo, err := dao.NewRecRepository(mysqlDBAddr)
	if err != nil {
		zlog.Logger.Error(err.Error())
		return
	}
	defer repo.Close()

	service, err := server.NewRecServer(&service.RecService{
		RecRepository: repo,
	})
	if err != nil {
		zlog.Logger.Error(err.Error())
		return
	}

	go func() {
		err = service.Run()
		if err != nil {
			zlog.Logger.Error("falied to run service", zap.Error(err))
			return
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	<-termChan

	closed := make(chan struct{})
	go func() {
		_ = repo.Close()
		_ = service.Server().Stop()

		close(closed)
	}()

	select {
	case <-closed:
		zlog.Logger.Info("graceful shutdown")
	case <-time.After(3 * time.Second):
		zlog.Logger.Error("shutdown timeout")
	}
}
