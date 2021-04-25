package server

import (
	"context"
	pb "rec-service/api/grpc/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*RecServer) Rec(ctx context.Context, req *pb.RecReq, res *pb.RecResponse) error {
	return status.Error(codes.Unimplemented, "")
}
