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
	FindCategoryComicDetail(ids []int32) ([]*model.CategoryDetail, error)
	FindComicChapters(id int32) ([]model.ComicChapter, error)
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

func (s *ComicService) ListComicCategoryDetail(ctx context.Context, req *pb.ListComicCategoryDetailRequest, res *pb.ListComicCategoryDetailResponse) error {
	mos, err := s.ComicRepository.FindCategoryDetail(req.Type, int(req.Sort), int(req.Offset), int(req.Limit))
	if err != nil || len(mos) == 0 {
		return status.Error(codes.NotFound, "")
	}

	res.Details = newComicCategoryDetail(mos)
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
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}

	res.ComicSpecials = newComicSpecial(mos)
	return nil
}

func (s *ComicService) ListCategoryComicDetail(ctx context.Context, req *pb.ListCategoryComicDetailRequest, res *pb.ListCategoryComicDetailResponse) error {
	mos, err := s.ComicRepository.FindCategoryComicDetail(req.ComicIds)
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}
	if len(mos) == 0 {
		return status.Error(codes.NotFound, "未找到任何资源!")
	}

	res.Comics = make([]*pb.ComicCategoryDetail, len(mos))
	for i := 0; i < len(res.Comics); i++ {
		res.Comics[i] = newCategoryComicDetail(mos[i])
	}

	return nil
}

func (s *ComicService) ListComicChapter(ctx context.Context, req *pb.ListComicChapterRequest, res *pb.ListComicChapterResponse) error {
	mos, err := s.ComicRepository.FindComicChapters(req.ComicId)
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}

	res.Chapters = make([]*pb.ChapterDetail, len(mos))
	for i := 0; i < len(mos); i++ {
		res.Chapters[i] = newComicChapterDetail(&mos[i])
	}

	return nil
}
