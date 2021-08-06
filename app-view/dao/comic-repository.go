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

	"github.com/go-kratos/kratos/pkg/sync/errgroup"
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
	// 动漫专题取一张
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

	srs, err := r.ComicService.ListComicCategoryDetail(ctx, &comicsrvpb.ListComicCategoryDetailRequest{
		Type:   categoryName,
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
	eg := errgroup.WithCancel(ctx)
	result := make([]struct {
		Ret interface{}
		Err error
	}, 3)

	// 漫画详情1
	eg.Go(func(ctx context.Context) error {
		r, err := r.ComicService.ListCategoryComicDetail(ctx, &comicsrvpb.ListCategoryComicDetailRequest{
			ComicIds: comicIds,
		})
		result[0].Ret = r
		result[0].Err = err
		return err
	})
	// 漫画详情2
	eg.Go(func(ctx context.Context) error {
		r, err := r.ComicService.ListComicDetail(ctx, &comicsrvpb.ListComicDetailRequest{
			ComicIds: func() []int64 {
				rs := make([]int64, 0, len(comicIds))
				for _, v := range comicIds {
					rs = append(rs, int64(v))
				}
				return rs
			}(),
		})
		result[1].Ret = r
		result[1].Err = err
		return err
	})
	// 漫画章节
	eg.Go(func(ctx context.Context) error {
		inchan := make(chan int32, len(comicIds))
		for _, v := range comicIds {
			inchan <- v
		}
		close(inchan)
		outchan := make(chan *comicsrvpb.ListComicChapterResponse, 1)

		var wg sync.WaitGroup
		concur := 16
		if concur > len(comicIds) {
			concur = len(comicIds)
		}

		for i := 0; i < concur; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for v := range inchan {
					r, err := r.ComicService.ListComicChapter(ctx, &comicsrvpb.ListComicChapterRequest{
						ComicId: v,
					})
					if err != nil {
						continue
					}

					outchan <- r
				}
			}()
		}

		go func() {
			wg.Wait()
			close(outchan)
		}()

		rs := make([]*comicsrvpb.ListComicChapterResponse, 0, len(comicIds))
		for v := range outchan {
			rs = append(rs, v)
		}

		result[2].Ret = rs
		result[2].Err = nil
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, formatErrComicServer(err)
	}

	ccd := result[0].Ret.(*comicsrvpb.ListCategoryComicDetailResponse)
	cd := result[1].Ret.(*comicsrvpb.ListComicDetailResponse)
	ccs := result[2].Ret.([]*comicsrvpb.ListComicChapterResponse)

	pcds := make([]*pb.ComicDetail, 0, len(cd.Comics))
	for _, v := range cd.Comics {
		pcds = append(pcds, convertComicDetail(v))
	}

	// 漫画id - 索引映射
	m1 := map[int32]int{}
	for i, v := range ccd.Comics {
		m1[int32(v.Id)] = i
	}
	m2 := map[int32]int{}
	for i, v := range ccs {
		// 注意，这里可能没有章节
		if len(v.Chapters) == 0 {
			continue
		}
		m2[v.Chapters[0].ComicId] = i
	}

	res := make([]*pb.ComicDetail, 0)
	for _, v := range pcds {
		i, ok1 := m1[v.Id]
		k, ok2 := m2[v.Id]

		// 降级，因为没爬取全部漫画的章节信息，所以返回的漫画可能没有章节信息
		if ok2 {
			ccsItem := ccs[k]
			v.Chapters = convertComicChapters(ccsItem.Chapters)
		}

		if ok1 {
			ccdItem := ccd.Comics[i]
			v.Authors = ccdItem.Authors
			v.Types = ccdItem.Types
			v.Status = ccdItem.Status

			res = append(res, v)
		}
	}

	return res, nil
}

func (r *ComicRepository) ListComicChapterDetail(ctx context.Context, comicId, chapterId int32) (*pb.ListComicChapterDetailResponse, error) {
	// TODO 加入redis缓存, 这里获取漫画的全部章节, 而实际上只需要一章, 太亏了
	res, err := r.ComicService.ListComicChapter(ctx, &comicsrvpb.ListComicChapterRequest{
		ComicId: comicId,
	})
	if err != nil {
		return nil, formatErrComicServer(err)
	}

	if len(res.Chapters) == 0 {
		return nil, formatErrComicServer(fmt.Errorf("没有找到漫画章节"))
	}

	mo := new(comicsrvpb.ChapterDetail)
	for _, v := range res.Chapters {
		if v.Chapterid == chapterId {
			mo = v
			break
		}
	}

	return convertComicChapterDetail(mo)
}
