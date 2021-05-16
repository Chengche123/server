module auth-service

go 1.16

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	share => ../share
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/medivhzhan/weapp/v2 v2.4.1
	github.com/micro/go-micro/v2 v2.9.1
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gorm.io/gorm v1.21.9
	share v0.0.0-00010101000000-000000000000
)
