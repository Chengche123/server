package service

import (
	"context"
	pb "news-service/api/grpc/v1"
	"news-service/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NewsRepository interface {
	FindCategoryDetail(tagId, sort, offset, limit int) ([]model.NewsCategoryDetail, error)
}

type NewsService struct {
	NewsRepository NewsRepository
}

func (s *NewsService) ListNewsCategoryDetail(ctx context.Context, req *pb.ListNewsCategoryDetailRequest, res *pb.ListNewsCategoryDetailResponse) error {
	mos, err := s.NewsRepository.FindCategoryDetail(int(req.TagId), int(req.Sort), int(req.Offset), int(req.Limit))
	if err != nil || len(mos) == 0 {
		return status.Error(codes.NotFound, "")
	}

	res.Details = newNewsCategoryDetail(mos)
	return nil
}

// TODO
func (s *NewsService) ListNewsCategory(ctx context.Context, req *pb.ListNewsCategoryRequest, res *pb.ListNewsCategoryResponse) error {
	return status.Error(codes.Unimplemented, "")
}
