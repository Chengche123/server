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
	ListComicChapterDetail(ctx context.Context, comicId, chapterId int32) (*pb.ListComicChapterDetailResponse, error)
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

func (s *AppviewService) ListComicChapterDetail(ctx context.Context, req *pb.ListComicChapterDetailRequest, res *pb.ListComicChapterDetailResponse) error {
	res1, err := s.ComicRepository.ListComicChapterDetail(ctx, req.ComicId, req.ChapterId)
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}

	res.ChapterId = res1.ChapterId
	res.ChapterOrder = res1.ChapterOrder
	res.ComicId = res1.ComicId
	res.CommentCount = res1.CommentCount
	res.Direction = res1.Direction
	res.PageUrl = res1.PageUrl
	res.Picnum = res1.Picnum
	res.Title = res1.Title

	return nil
}
