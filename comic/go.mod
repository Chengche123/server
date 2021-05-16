module comic-service

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

go 1.16

replace comic => ../

replace crawler => ../../crawler

require (
	comic v0.0.0-00010101000000-000000000000
	crawler v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2
	github.com/google/gofuzz v1.0.0
	github.com/micro/go-micro/v2 v2.9.1
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gorm.io/gorm v1.21.9
)
