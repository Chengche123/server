package dao

import (
	pb "app-view/api/grpc/v1"
	comicsrvpb "comic-service/api/grpc/v1"
	comicserver "comic-service/server/grpc"
	"errors"
	newsgrpcpb "news-service/api/grpc/v1"
	newsgrpc "news-service/server/grpc"
	"sync"

	zlog "share/log/zap"

	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"go.uber.org/zap"
)

var (
	ErrComicServer = errors.New("comic service error")
)

func formatErrComicServer(err error) error {
	return fmt.Errorf("%v: %v", ErrComicServer, err)
}

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
	cc, err := r.ComicService.ListComicCategoryFilter(ctx, &comicsrvpb.ListComicCategoryFilterRequest{})
	if err != nil {
		return nil, formatErrComicServer(err)
	}

	res := make([]*pb.CategoryMo, 1)
	res[0] = &pb.CategoryMo{
		Name: "推荐",
	}

	for _, v := range cc.Filters {
		if v.Title == "题材" {
			for _, v1 := range v.Items {
				if v1.TagName != "全部" {
					res = append(res, &pb.CategoryMo{
						Name: v1.TagName,
					})
				}
			}
		}
	}

	return res, nil
}

// TODO 降级
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

	// 依次放新闻 动画(1) 漫画(2) 轻小说(3)和 美图(8)
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

func (r *ComicRepository) ListFeedComicMo(ctx context.Context, categoryName string, pageIndex, pageSize int) ([]*pb.FeedComicMo, error) {
	var srs *comicsrvpb.ListComicCategoryDetailResponse
	limit := pageSize
	if limit == 0 {
		return make([]*pb.FeedComicMo, 0), nil
	}
	offset := pageSize * pageIndex

	var type1 string
	if categoryName != "推荐" {
		type1 = categoryName
	}

	srs, err := r.ComicService.ListComicCategoryDetail(ctx, &comicsrvpb.ListComicCategoryDetailRequest{
		Type:   type1,
		Limit:  int32(limit),
		Offset: int32(offset),
		Sort:   2, // feed
	})
	if err != nil {
		return nil, formatErrComicServer(err)
	}

	res := make([]*pb.FeedComicMo, 0)
	for _, v := range srs.Details {
		res = append(res, convertComicCategoryDetailToFeedComic(v))
	}

	return res, nil
}

func (r *ComicRepository) ListComicDetail(ctx context.Context, comicIds []int32) ([]*pb.ComicDetail, error) {
	type outt struct {
		coms *comicsrvpb.ListCategoryComicDetailResponse
		err  error
	}
	waitChan := make(chan *outt, 1)

	go func() {
		r, err := r.ComicService.ListCategoryComicDetail(ctx, &comicsrvpb.ListCategoryComicDetailRequest{
			ComicIds: comicIds,
		})
		waitChan <- &outt{
			coms: r,
			err:  err,
		}
	}()

	cs, err := r.ComicService.ListComicDetail(ctx, &comicsrvpb.ListComicDetailRequest{
		ComicIds: func() []int64 {
			rs := make([]int64, 0, len(comicIds))
			for _, v := range comicIds {
				rs = append(rs, int64(v))
			}
			return rs
		}(),
	})

	if err != nil {
		return nil, formatErrComicServer(err)
	}

	pcds := make([]*pb.ComicDetail, 0, len(cs.Comics))
	for _, v := range cs.Comics {
		pcds = append(pcds, convertComicDetail(v))
	}

	gr := <-waitChan
	if gr.err != nil {
		return nil, formatErrComicServer(err)
	}

	m1 := map[int32]int{}
	for i, v := range gr.coms.Comics {
		m1[int32(v.Id)] = i
	}

	res := make([]*pb.ComicDetail, 0)
	for _, v := range pcds {
		i, ok := m1[v.Id]
		if ok {
			tmp := gr.coms.Comics[i]
			v.Authors = tmp.Authors
			v.Types = tmp.Types
			v.Status = tmp.Status

			res = append(res, v)
		}
	}

	return res, nil
}
