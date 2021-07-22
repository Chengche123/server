package dao

import (
	pb "app-view/api/grpc/v1"
	comicsrvpb "comic-service/api/grpc/v1"
	comicserver "comic-service/server/grpc"
	newsgrpcpb "news-service/api/grpc/v1"
	newsgrpc "news-service/server/grpc"
	"sync"

	zlog "share/log/zap"

	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"go.uber.org/zap"
)

type ComicRepository struct {
	ComicService comicsrvpb.ComicService
	NewsService  newsgrpcpb.NewsService
}

func NewComicRepository(service micro.Service) (*ComicRepository, error) {
	return &ComicRepository{
		ComicService: comicsrvpb.NewComicService(comicserver.ServiceName, service.Client()),
		NewsService:  newsgrpcpb.NewNewsService(newsgrpc.ServiceName, service.Client()),
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

// TODO 降级
// TODO 并发
func (r *ComicRepository) ListBannerMo(ctx context.Context) ([]*pb.BannerMo, error) {
	// 动漫专题取三张
	comicSpecial, err := r.ComicService.ListComicSpecial(ctx, &comicsrvpb.ListComicSpecialRequest{
		Offset: 0,
		Limit:  1,
	})
	if err != nil {
		zlog.Logger.Info("comic rpc server error", zap.Error(err))
		return nil, fmt.Errorf("comic rpc server error: %v", err)
	}

	res := make([]*pb.BannerMo, 5)

	// 0 放动漫专题
	specials := comicSpecial.ComicSpecials
	res[0] = convertSpecialToBanner(specials[0])

	// 依次放 动画(1) 漫画(2) 轻小说(3)和 美图(8)
	tids := make(chan int, 4)
	for _, i := range []int{1, 2, 3, 8} {
		tids <- i
	}
	close(tids)

	type outStruct struct {
		news *newsgrpcpb.NewsCategoryDetail
		tid  int
	}
	outchan := make(chan *outStruct, 4)
	concur := 4
	var wg sync.WaitGroup

	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for id := range tids {
				ns, err := r.NewsService.ListNewsCategoryDetail(ctx, &newsgrpcpb.ListNewsCategoryDetailRequest{
					TagId:  int32(id),
					Sort:   1,
					Limit:  1,
					Offset: 0,
				})
				if err != nil {
					zlog.Logger.Info("news-service error", zap.Error(err), zap.Int("tag_id", id))
					continue
				}

				outchan <- &outStruct{
					news: ns.Details[0],
					tid:  id,
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outchan)
	}()

	for out := range outchan {
		index := out.tid
		if out.tid == 8 {
			index = 4
		}
		res[index] = convertNewsToBanner(out.news)
	}

	return res, nil
}
