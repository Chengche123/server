package server

import (
	"fmt"
	interceptor "interceptor-micro/auth"
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
	srvName          = "go.micro.api.comic.rec.v1"
	registryAddr     = env.FormatEnvOrDefault("%s", "COMIC_REGISTRY_ADDR", "127.0.0.1:2379")
)

func NewRecServer(srv pb.RecServiceHandler) (micro.Service, error) {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{registryAddr}
	})

	authInter, err := interceptor.NewAuthInterceptor(pubKey)
	if err != nil {
		return nil, fmt.Errorf("cannot init auth interceptor: %v", err)
	}

	service := micro.NewService(
		micro.WrapHandler(authInter),
		micro.Registry(reg),
		micro.Name(srvName),
		micro.RegisterTTL(registerTTL),
		micro.RegisterInterval(registerInterval),
	)

	service.Init()

	err = pb.RegisterRecServiceHandler(service.Server(), srv)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %v", err)
	}

	return service, nil
}

const pubKey = `
-----BEGIN PUBLIC KEY-----
MIIBITANBgkqhkiG9w0BAQEFAAOCAQ4AMIIBCQKCAQBb+HGPGnfVG6l/SSPg20fH
l1brmbQ2Nrw880AA4eRFUhp7w+u9fKDpzcJhfrDLga28hrcgn4nURDZEtXL//ISA
AI1+H5py4o3OGwhZ42jj01VHXiCdBTA9YqY79dMxm1uD401ANaVRpdb0HMBXyopb
80630F1NxDMaAIgfoBPCZWdyE2pwsTBUVoexK+/X+6pDUxJpD7RK5DLEwk5IsymL
OIPmxauD9r9z+itaWyFRVi/WYbe6Pxf9tVK9vgVm0A51rxJeOBHCMSvd6xgUB5tu
/IvXQTyh/G+ZASIAObGGyqqqY8NuFnWESP3zu3WpsrD579K90HeVgkXaqTBH7+RP
AgMBAAE=
-----END PUBLIC KEY-----
`
