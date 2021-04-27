module micro/auth

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace (
	comic => E:\final\Projects\server
	interceptor-micro => E:\final\Projects\server\share\interceptor\micro
)

require (
	comic v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.9.1
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
)
