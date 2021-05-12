package service

import (
	"context"
	pb "rec-service/api/grpc/v1"
	"rec-service/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecRepository interface {
	GetComicsByHot(offset int, limit int) ([]model.Comic, error)
}

type RecService struct {
	RecRepository RecRepository
}

func (s *RecService) Rec(ctx context.Context, req *pb.RecReq, res *pb.RecResponse) error {
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

func (s *RecService) RelatedRec(ctx context.Context, req *pb.RecReq, res *pb.RecResponse) error {
	return status.Error(codes.Unimplemented, "")
}

func (s *RecService) AuthorRec(ctx context.Context, req *pb.RecReq, res *pb.RecResponse) error {
	return status.Error(codes.Unimplemented, "")
}
