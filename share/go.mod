module share

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/docker v20.10.6+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.9.1
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	go.mongodb.org/mongo-driver v1.5.2
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.37.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.9
	gotest.tools/v3 v3.0.3 // indirect
)
