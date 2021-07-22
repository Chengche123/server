module app-view

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace (
	comic-service => ../comic
	crawler => ../../crawler
	news-service => ../news
	share => ../share
)

go 1.16

require (
	comic-service v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	news-service v0.0.0-00010101000000-000000000000
	share v0.0.0-00010101000000-000000000000
)
