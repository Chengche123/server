package server

import (
	"fmt"
	"share/os/env"
	"strings"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

var (
	registerTTL      = 30 * time.Second
	registerInterval = 10 * time.Second
	registryAddr     = env.FormatEnvOrDefault("%s", "COMIC_REGISTRY_ADDR", "127.0.0.1:2379")
)

// NewMicroServer 返回一个micro server实例
func NewMicroServer(name string) (micro.Service, error) {
	if !strings.HasPrefix(name, "go.micro.") {
		return nil, fmt.Errorf("invalid service name: %v", name)
	}

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{registryAddr}
	})

	service := micro.NewService(
		micro.Name(name),
		micro.Registry(reg),
		micro.RegisterTTL(registerTTL),
		micro.RegisterInterval(registerInterval),
	)

	return service, nil
}
