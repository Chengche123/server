package dao

import (
	pb "app-view/api/grpc/v1"
	comicsrvpb "comic-service/api/grpc/v1"
	newsgrpcpb "news-service/api/grpc/v1"

	"strconv"
)

func convertNewsToBanner(news *newsgrpcpb.NewsCategoryDetail) *pb.BannerMo {
	res := new(pb.BannerMo)
	res.Id = strconv.Itoa(int(news.ArticleId))
	res.Sticky = 0
	res.Type = strconv.Itoa(int(news.TagId))
	res.Title = news.Title
	res.Subtitle = news.Intro
	res.Url = news.PageUrl
	res.Cover = news.RowPicUrl
	res.CreateTime = strconv.Itoa(int(news.CreateTime))

	return res
}

func convertSpecialToBanner(s *comicsrvpb.ComicSpecial) *pb.BannerMo {
	r := new(pb.BannerMo)
	r.Cover = s.SmallCover
	r.Title = s.Title
	r.Id = strconv.Itoa(int(s.Id))
	r.Subtitle = s.ShortTitle
	r.Url = s.PageUrl
	r.CreateTime = strconv.Itoa(int(s.CreateTime))
	r.Type = strconv.Itoa(int(s.PageType))
	r.Sticky = 0

	return r
}

func convertComicCategoryDetailToFeedComic(d *comicsrvpb.ComicCategoryDetail) *pb.FeedComicMo {
	r := new(pb.FeedComicMo)

	r.Authors = d.Authors
	r.Cover = d.Cover
	r.Id = int32(d.Id)
	r.LastUpdatetime = int32(d.LastUpdatetime)
	r.Num = int32(d.Num)
	r.Status = d.Status
	r.Title = d.Title
	r.Types = d.Types

	return r
}
