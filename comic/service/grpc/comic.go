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
	GetComicInfos(ids []int64) ([]*model.Comic, error)
}

type ComicService struct {
	Logger     *zap.Logger
	Repository ComicRepository
}

func (s *ComicService) ListComicInfo(ctx context.Context, req *pb.ListComicInfoRequest, res *pb.ListComicInfoResponse) error {
	entiry, err := s.Repository.GetComicInfos(req.ComicIds)
	if err != nil || len(entiry) == 0 {
		return status.Error(codes.NotFound, "")
	}

	res.Comics = convert(entiry)
	return nil
}

func convert(entity []*model.Comic) []*pb.ComicInfo {
	res := make([]*pb.ComicInfo, 0, len(entity))
	for _, v := range entity {
		c := new(pb.ComicInfo)

		c.AddTime = v.AddTime
		c.Authors = v.Authors
		c.Cover = v.Cover
		c.Id = v.Id
		c.IsEnd = int32(v.IsEnd)
		c.LastUpdatetime = v.LastUpdatetime
		c.Num = v.Num
		c.Status = int32(v.Status)
		c.Title = v.Title
		c.Types = v.Types

		res = append(res, c)
	}

	return res
}
