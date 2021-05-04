package server

import (
	"fmt"
	pb "rec-service/api/grpc/v1"
	"time"

	"comic/share/os/env"

	"github.com/micro/go-micro/v2"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

var (
	registerTTL      = 30 * time.Second
	registerInterval = 10 * time.Second
	srvName          = "go.micro.srv.comic.rec.v1"
	registryAddr     = env.FormatEnvOrDefault("%s", "COMIC_REGISTRY_ADDR", "127.0.0.1:2379")
)

func NewRecServer(srv pb.RecServiceHandler) (micro.Service, error) {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{registryAddr}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name(srvName),
		micro.RegisterTTL(registerTTL),
		micro.RegisterInterval(registerInterval),
	)

	service.Init()

	err := pb.RegisterRecServiceHandler(service.Server(), srv)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %v", err)
	}

	return service, nil
}
