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
}

type ComicService struct {
	Logger     *zap.Logger
	Repository ComicRepository
}

func (s *ComicService) ListComicDetail(ctx context.Context, req *pb.ListComicDetailRequest, res *pb.ListComicDetailResponse) error {
	entiry, err := s.Repository.FindComicDetails(req.ComicIds)
	if err != nil || len(entiry) == 0 {
		return status.Error(codes.NotFound, "")
	}

	res.Comics = newComicDetail(entiry)
	return nil
}

func (s *ComicService) ListCategoryDetail(ctx context.Context, req *pb.ListCategoryDetailRequest, res *pb.ListCategoryDetailResponse) error {
	return nil
}
