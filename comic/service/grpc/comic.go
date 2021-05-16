package service

import (
	pb "comic-service/api/grpc/v1"
	"comic-service/model"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ComicRepository interface {
	FindComicDetails(ids []int64) ([]*model.ComicDetail, error)
	FindCategoryDetail(types string, sort, offset, limit int) ([]model.CategoryDetail, error)
	FindComicCategoryFilter() ([]model.ComicCategoryFilter, error)
	FindComicSpecial(offset, limit int) ([]model.ComicSpecial, error)
	Close() error
}

type ComicService struct {
	ComicRepository ComicRepository
}

func (s *ComicService) ListComicDetail(ctx context.Context, req *pb.ListComicDetailRequest, res *pb.ListComicDetailResponse) error {
	mos, err := s.ComicRepository.FindComicDetails(req.ComicIds)
	if err != nil {
		return status.Error(codes.NotFound, "")
	}

	res.Comics = newComicDetail(mos)
	return nil
}

func (s *ComicService) ListCategoryDetail(ctx context.Context, req *pb.ListCategoryDetailRequest, res *pb.ListCategoryDetailResponse) error {
	mos, err := s.ComicRepository.FindCategoryDetail(req.Type, int(req.Sort), int(req.Offset), int(req.Limit))
	if err != nil || len(mos) == 0 {
		return status.Error(codes.NotFound, "")
	}

	res.Details = newCategoryDetail(mos)
	return nil
}

func (s *ComicService) ListComicCategoryFilter(ctx context.Context, req *pb.ListComicCategoryFilterRequest, res *pb.ListComicCategoryFilterResponse) error {
	mos, err := s.ComicRepository.FindComicCategoryFilter()
	if err != nil || len(mos) == 0 {
		return status.Error(codes.NotFound, "")
	}

	res.Filters = newComicCategoryFilter(mos)
	return nil
}

func (s *ComicService) ListComicSpecial(ctx context.Context, req *pb.ListComicSpecialRequest, res *pb.ListComicSpecialResponse) error {
	mos, err := s.ComicRepository.FindComicSpecial(int(req.Offset), int(req.Limit))
	if err != nil || len(mos) == 0 {
		return status.Error(codes.NotFound, "")
	}

	res.ComicSpecials = newComicSpecial(mos)
	return nil
}
