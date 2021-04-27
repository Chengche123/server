package service

import (
	"context"
	interceptor "interceptor-micro/auth"
	pb "rec-service/api/grpc/v1"
	"rec-service/model"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecRepository interface {
	GetComicsByHot(offset int, limit int) ([]model.Comic, error)
}

type Service struct {
	Logger        *zap.Logger
	RecRepository RecRepository
}

func (s *Service) Rec(ctx context.Context, req *pb.RecReq, res *pb.RecResponse) error {
	uid, err := interceptor.UidFromContext(ctx)
	if err != nil {
		return status.Error(codes.Unauthenticated, "")
	}
	s.Logger.Info("context", zap.String("uid", uid))

	comics, err := s.RecRepository.GetComicsByHot(int(req.Offset), int(req.Limit))
	if err != nil {
		return status.Error(codes.Internal, "")
	}
	if len(comics) == 0 {
		return status.Error(codes.NotFound, "failed to find recommend comics")
	}

	res.List = make([]*pb.RecRecord, len(comics))
	for i := 0; i < len(comics); i++ {
		rec := new(pb.RecRecord)
		rec.ComicId = comics[i].Id
		rec.Score = float64(comics[i].Num)
		res.List[i] = rec
	}

	return nil
}

func (s *Service) RelatedRec(ctx context.Context, req *pb.RecReq, res *pb.RecResponse) error {
	return nil
}

func (s *Service) AuthorRec(ctx context.Context, req *pb.RecReq, res *pb.RecResponse) error {
	return nil
}
