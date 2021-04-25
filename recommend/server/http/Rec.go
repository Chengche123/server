package server

import (
	"context"
	"net/http"
	pb "rec-service/api/grpc/v1"

	api "github.com/micro/go-micro/v2/api/proto"
	"go.uber.org/zap"
)

type Service struct {
	Grpc   pb.RecService
	Logger *zap.Logger
}

func (*Service) Rec(ctx context.Context, req *api.Request, res *api.Response) error {
	res.StatusCode = http.StatusOK
	res.Body = "TODO"
	return nil
}
