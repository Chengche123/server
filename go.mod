module comic

go 1.15

require (
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/beego/beego/v2 v2.0.1
	github.com/containerd/containerd v1.4.4 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v20.10.5+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.4.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0
	github.com/jmoiron/sqlx v1.3.1
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/medivhzhan/weapp/v2 v2.4.1
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	go.mongodb.org/mongo-driver v1.5.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	google.golang.org/grpc v1.36.1
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.6
	gotest.tools/v3 v3.0.3 // indirect
)

replace(
	google.golang.org/grpc => google.golang.org/grpc v1.36.1
)
