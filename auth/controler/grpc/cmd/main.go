package main

import (
	authpb "comic/auth/controler/grpc/api/gen/v1"
	"comic/auth/dao/mongo"
	"comic/auth/dao/mysql/repository"
	"comic/auth/service"
	"comic/auth/token"
	"comic/auth/wechat"
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	controler "comic/auth/controler/grpc"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const serviceAddr = ":7000"
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
	ctx, cancel := context.WithCancel(context.Background())

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	lis, err := net.Listen("tcp", serviceAddr)
	_ = lis
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	gserver := grpc.NewServer()

	openidResolver, err := mongo.NewMongoCol(
		ctx,
		logger,
		"mongodb://localhost:27017",
		"comic")
	if err != nil {
		logger.Fatal("failed to connect mongodb", zap.Error(err))
	}

	tokenGenerator, err := token.NewJWTTokenGen([]byte(privateKey), "comic/auth")
	if err != nil {
		logger.Fatal("cannot init tokenGenerator", zap.Error(err))
	}

	rawDB, err := sql.Open(
		"mysql", "root:root@tcp(127.0.0.1:3306)/comic")
	if err != nil {
		logger.Fatal("cannot open sql connect", zap.Error(err))
	}
	userRepository, err := repository.NewMySqlTable(ctx, rawDB)
	if err != nil {
		logger.Fatal("falied to connect mysql", zap.Error(err))
	}

	serviceBuilder := &controler.ParamSet{
		Logger:         logger,
		CodeResolver:   wechat.NewWechatService("6c1da32eb144746a75425a82e8505973", "wx6ed5600ae559cbd8"),
		OpenIDResolver: openidResolver,
		TokenExpireIn:  time.Hour * 2,
		TokenGenerator: tokenGenerator,
		UserResolver: &service.UserService{
			UserRepository: userRepository,
			TokenGenerator: tokenGenerator,
			Logger:         logger,
		},
	}
	authService, err := serviceBuilder.Build()
	if err != nil {
		logger.Fatal("failed to init auth service", zap.Error(err))
	}
	defer func() {
		err := authService.Close()
		if err != nil {
			logger.Error("falied to close authService", zap.Error(err))
		}
	}()

	authpb.RegisterAuthServiceServer(gserver, authService)

	logger.Info("auth service is starting", zap.String("addr", serviceAddr))

	go func() {
		err = gserver.Serve(lis)
		if err != nil {
			logger.Fatal("auth service disconnected", zap.Error(err))
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	<-termChan

	closed := make(chan struct{})
	go func() {
		cancel()
		gserver.Stop()
		authService.Close()
		lis.Close()

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
