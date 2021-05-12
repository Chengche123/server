package service

import (
	pb "comic-service/api/grpc/v1"
	"comic-service/model"
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ComicRepository interface {
	FindComicDetails(ids []int64) ([]*model.ComicDetail, error)
	FindCategoryDetail(types string, sort, offset, limit int) ([]model.CategoryDetail, error)
}

type ComicService struct {
	Logger     *zap.Logger
	Repository ComicRepository
}

func (s *ComicService) ListComicDetail(ctx context.Context, req *pb.ListComicDetailRequest, res *pb.ListComicDetailResponse) error {
	mos, err := s.Repository.FindComicDetails(req.ComicIds)
	if err != nil {
		return status.Error(codes.NotFound, "")
	}

	res.Comics = newComicDetail(mos)
	return nil
}

func (s *ComicService) ListCategoryDetail(ctx context.Context, req *pb.ListCategoryDetailRequest, res *pb.ListCategoryDetailResponse) error {
	mos, err := s.Repository.FindCategoryDetail(req.Type, int(req.Sort), int(req.Offset), int(req.Limit))
	if err != nil || len(mos) == 0 {
		return status.Error(codes.NotFound, "")
	}

	res.Details = newCategoryDetail(mos)
	return nil
}
