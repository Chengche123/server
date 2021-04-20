package main

import (
	grpcInterceptor "comic/share/interceptor/grpc"
	"comic/share/os/env"
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	microInterceptor "interceptor-micro/interceptor"

	authpb "micro/auth/api/gen/v1"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"go.uber.org/zap"

	"comic/auth/dao/mysql/repository"
	userService "comic/auth/service"
	"comic/auth/token"

	controler "micro/auth/grpc"
)

var (
	registryAddr = env.FormatEnvOrDefault("http://%s", "COMIC_REGISTRY_ADDR", "127.0.0.1:2379")
	mysqlDBAddr  = env.FormatEnvOrDefault("root:root@tcp(%s)/comic", "COMIC_MYSQL_ADDR", "127.0.0.1:3306")
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEoQIBAAKCAQBb+HGPGnfVG6l/SSPg20fHl1brmbQ2Nrw880AA4eRFUhp7w+u9
fKDpzcJhfrDLga28hrcgn4nURDZEtXL//ISAAI1+H5py4o3OGwhZ42jj01VHXiCd
BTA9YqY79dMxm1uD401ANaVRpdb0HMBXyopb80630F1NxDMaAIgfoBPCZWdyE2pw
sTBUVoexK+/X+6pDUxJpD7RK5DLEwk5IsymLOIPmxauD9r9z+itaWyFRVi/WYbe6
Pxf9tVK9vgVm0A51rxJeOBHCMSvd6xgUB5tu/IvXQTyh/G+ZASIAObGGyqqqY8Nu
FnWESP3zu3WpsrD579K90HeVgkXaqTBH7+RPAgMBAAECggEAJRBDW2NURYqnTQeB
sP1NuQ6xVbMUoACA7aMt+O0P3CvRRm1XVH6kLnTgUAAJaYnyJRa5OClzFxsCL8Bb
/vOZxify3ZAI4yGP+i6EuAzgWWfyZxd01zKfFS3rRiC0Irq0L7trh2PXIsdNkAUC
Mp8KETJFV+hdoZhe66ypLu2I6P8OJ5VVQsvoephjjxXRCA9IvV/IINcc53IoBZkX
tV9VgjceKyP0ltV/+RP7xGt5Mgl5FmtEsyfn6EQl6stIZI7vK+2rgWITIjwKtFKT
u1wLi39VmQ8pqtpxKoh+IuGtCHJJnKkha4BbChNmV7UNYTBxz4zirJof5bgnf27N
sWHbmQKBgQCsQgLWdFWRhsEg2++DtOpA2ow6NdsKxByRZZ2WphC9wIw0DA0uuA4d
FMvWe4oR+U2np2Kl8M7LFmyb8bImUOLx74yV+c3OkRz+1mDqG8i+Uuvt5n7C3McM
bJ55RdhIsZRWtKMkLX5B39jcW7eFE8s17FKdQ2DXxqK2ABiEFNKf4wKBgQCIrm+8
LlyxjdMk6B+4c1x3T/A6fXEIQoCRI2yWYk/PMnMA7B5GjBbwc9PNEYmUC3tVdiYS
m03cXepQFOWBR2zCn5SHWrxlNf3OIdMDAGZEMslLokw3AVokjqpSd48rsExl8Nwo
rqjyJJzygSfXcJQVK7ukfJb9+fex+hVoC/99pQKBgCVqVSl1rVxls69KdlTaSAN8
NPcz2XcWL8pZEwi023AL0ahAceCS8+XXLYtR3CSZTQe8cM3wZ0pvfXnF3tc5vIGm
cZfl4ZEbrfugXv6auFi2tC5BYYk74TROp4FZ7Wekwr6uj7z88K6oS3dZqJwMN5hn
0237Q566s37quGiACVvXAoGAMPIupzf/D8JXI3dQrK/7I+rnfzqeuLN/8Pm7kBsC
s5NAZcsoiGDwcgk4hs5J808tSox10+Hzvv+OwkDJc0NNqmCVQud2YdQ04JVySDfj
Z3a8aS5klqoQStkgG0ofZijLIxJe9N6fN59u6NThnJ//F5nCp4PKWocicYS6F7l8
M9UCgYBRsz2wQHjrYuTEL2lZqYMJtxubLp8U/SuRNw0zI8eJ/p+Xk20lykJva1Io
SCG5wX5B/vfk/rzSo16ThA+5T/l9JGKko/7UKCyOrUuuTEANEdbn/kF9cWXogC9M
57G38AqlNiTrsGyIp2lhNdvWpQ7OeNQisHkpB65ZzUD81oSxBg==
-----END RSA PRIVATE KEY-----`

func main() {
	ctx := context.Background()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create zap logger: %v", err)
	}

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{registryAddr}
	})

	authWrapper := microInterceptor.NewAuthInterceptor(grpcInterceptor.NewAuthInterceptor())

	service := micro.NewService(
		micro.WrapHandler(authWrapper),
		micro.Registry(reg),
		micro.Name("go.micro.srv.comic.auth.v1"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(10*time.Second),
	)

	service.Init()

	tokenGenerator, err := token.NewJWTTokenGen([]byte(privateKey), "comic/auth")
	if err != nil {
		logger.Fatal("cannot init tokenGenerator", zap.Error(err))
	}

	rawDB, err := sql.Open(
		"mysql", mysqlDBAddr)
	if err != nil {
		logger.Fatal("cannot open sql connect", zap.Error(err))
	}
	defer rawDB.Close()

	userRepository, err := repository.NewMySqlTable(ctx, rawDB)
	if err != nil {
		logger.Fatal("falied to connect mysql", zap.Error(err))
	}

	ctrl := &controler.Service{
		UserResolver: &userService.UserService{
			TokenGenerator: tokenGenerator,
			UserRepository: userRepository,
		},
		Logger: logger,
	}

	err = authpb.RegisterAuthServiceHandler(service.Server(), ctrl)
	if err != nil {
		logger.Fatal("cannot register service handler", zap.Error(err))
	}

	go func() {
		err = service.Run()
		if err != nil {
			logger.Fatal("falied to run service", zap.Error(err))
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
		err = ctrl.Close()
		if err != nil {
			logger.Error("failed to stop controler", zap.Error(err))
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
