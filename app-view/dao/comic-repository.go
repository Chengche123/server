package dao

import (
	pb "app-view/api/grpc/v1"
	comicsrvpb "comic-service/api/grpc/v1"
	comicserver "comic-service/server/grpc"
	zlog "share/log/zap"

	"context"
	"fmt"
	"strconv"

	"github.com/micro/go-micro/v2"
	"go.uber.org/zap"
)

type ComicRepository struct {
	ComicService comicsrvpb.ComicService
}

func NewComicRepository(service micro.Service) (*ComicRepository, error) {
	return &ComicRepository{
		ComicService: comicsrvpb.NewComicService(comicserver.ServiceName, service.Client()),
	}, nil
}

func (r *ComicRepository) ListCategoryMo(ctx context.Context) ([]*pb.CategoryMo, error) {
	res := []*pb.CategoryMo{
		{
			Name: "推荐",
		},
		{
			Name: "更新",
		},
		{
			Name: "分类",
		},
		{
			Name: "排行",
		},
		{
			Name: "专题",
		},
	}

	return res, nil
}

func (r *ComicRepository) ListBannerMo(ctx context.Context) ([]*pb.BannerMo, error) {
	comicSpecial, err := r.ComicService.ListComicSpecial(ctx, &comicsrvpb.ListComicSpecialRequest{
		Offset: 0,
		Limit:  5,
	})
	if err != nil {
		zlog.Logger.Info("comic rpc server error", zap.Error(err))
		return nil, fmt.Errorf("comic rpc server error: %v", err)
	}

	specials := comicSpecial.ComicSpecials
	res := make([]*pb.BannerMo, len(specials))
	for i := 0; i < len(specials); i++ {
		res[i] = &pb.BannerMo{
			Cover:      specials[i].SmallCover,
			Title:      specials[i].Title,
			Id:         strconv.Itoa(int(specials[i].Id)),
			Subtitle:   specials[i].ShortTitle,
			Url:        specials[i].PageUrl,
			CreateTime: strconv.Itoa(int(specials[i].CreateTime)),
			Type:       strconv.Itoa(int(specials[i].PageType)),
			Sticky:     0, // ?
		}
	}

	return res, nil
}
