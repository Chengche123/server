package service

import (
	pb "app-view/api/grpc/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ComicRepository interface {
	ListBannerMo(ctx context.Context) ([]*pb.BannerMo, error)
	ListCategoryMo(ctx context.Context) ([]*pb.CategoryMo, error)
}

type AppviewService struct {
	ComicRepository ComicRepository
}

func (s *AppviewService) ListHomeMo(ctx context.Context, req *pb.ListHomeMoRequest, res *pb.ListHomeMoResponse) error {
	banners, err := s.ComicRepository.ListBannerMo(ctx)
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}

	cates, err := s.ComicRepository.ListCategoryMo(ctx)
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}

	res.BannerList = banners
	res.CategoryList = cates

	return nil
}
