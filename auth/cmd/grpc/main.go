package main

import (
	server "auth-service/server/grpc"
	"os"
	"os/signal"
	config "share/config/database"
	zlog "share/log/zap"
	"share/os/env"
	"syscall"
	"time"

	"auth-service/dao"

	"go.uber.org/zap"

	authService "auth-service/service"
	"auth-service/token"
)

var (
	mysqlDBAddr = env.FormatEnvOrDefault("root:root@tcp(%s)/comic", "COMIC_MYSQL_ADDR", config.DefaultMySqlAddr)
)

func main() {
	tokenGenerator, err := token.NewJWTTokenGen([]byte(privateKey), "comic/auth")
	if err != nil {
		zlog.Logger.Error("", zap.Error(err))
		return
	}

	userRepository, err := dao.NewUserRepository(mysqlDBAddr)
	if err != nil {
		zlog.Logger.Error("", zap.Error(err))
		return
	}
	defer userRepository.Close()

	handler := &authService.AuthService{
		UserRepository: userRepository,
		TokenGenerator: tokenGenerator,
	}

	service, err := server.NewAuthServer(handler)
	if err != nil {
		zlog.Logger.Error("", zap.Error(err))
		return
	}

	go func() {
		err = service.Run()
		if err != nil {
			zlog.Logger.Error("", zap.Error(err))
			return
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	<-termChan

	closed := make(chan struct{})
	go func() {
		_ = userRepository.Close()
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
