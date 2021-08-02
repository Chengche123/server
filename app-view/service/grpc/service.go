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
	ListFeedComicMo(ctx context.Context, categoryName string, pageIndex, pageSize int) ([]*pb.FeedComicMo, error)
	ListComicDetail(ctx context.Context, comicIds []int32) ([]*pb.ComicDetail, error)
}

type AppviewService struct {
	ComicRepository ComicRepository
}

func (s *AppviewService) ListHomeMo(ctx context.Context, req *pb.ListHomeMoRequest, res *pb.ListHomeMoResponse) error {
	if req.CategoryName == "推荐" {
		banners, err := s.ComicRepository.ListBannerMo(ctx)
		if err != nil {
			return status.Error(codes.NotFound, err.Error())
		}
		res.BannerList = banners

		cates, err := s.ComicRepository.ListCategoryMo(ctx)
		if err != nil {
			return status.Error(codes.NotFound, err.Error())
		}

		res.CategoryList = cates
	}

	fcs, err := s.ComicRepository.ListFeedComicMo(ctx, req.CategoryName, int(req.PageIndex), int(req.PageSize))
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}

	res.ComicList = fcs

	return nil
}

func (s *AppviewService) ListComicDetail(ctx context.Context, req *pb.ListComicDetailRequest, res *pb.ListComicDetailResponse) error {
	rs, err := s.ComicRepository.ListComicDetail(ctx, req.ComicIds)
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}

	res.Comics = rs

	return nil
}
