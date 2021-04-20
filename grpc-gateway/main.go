package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	authpb "comic/auth/controler/grpc/api/gen/v1"
)

const serviceAddr = ":18080"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create zap logger: %v", err)
	}

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers:  true,
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
		},
	))

	err = authpb.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		":7000",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("failed to register auth service to gateway: %v", err)
	}

	server := http.Server{
		Addr:    serviceAddr,
		Handler: mux,
	}

	go func() {
		logger.Info("gateway service is starting", zap.String("addr", serviceAddr))
		logger.Info("gateway service shutdown", zap.Error(server.ListenAndServe()))
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	<-termChan

	closed := make(chan struct{})
	go func() {
		cancel()
		server.Close()

		time.Sleep(500 * time.Millisecond)
		close(closed)
	}()

	select {
	case <-closed:
		logger.Info("graceful shutdown")
	case <-time.After(2 * time.Second):
		logger.Info("shutdown timeout")
	}
}
