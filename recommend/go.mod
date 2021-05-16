module rec-service

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace share => ../share

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.9
	share v0.0.0-00010101000000-000000000000
)
