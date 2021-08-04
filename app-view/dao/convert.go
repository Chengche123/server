package dao

import (
	pb "app-view/api/grpc/v1"
	comicsrvpb "comic-service/api/grpc/v1"
	newsgrpcpb "news-service/api/grpc/v1"
	static "share/static/url"

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
	res.Cover = static.ConverURL(news.RowPicUrl) // 转本服务器
	res.CreateTime = strconv.Itoa(int(news.CreateTime))

	return res
}

func convertSpecialToBanner(s *comicsrvpb.ComicSpecial) *pb.BannerMo {
	r := new(pb.BannerMo)
	r.Cover = static.ConverURL(s.SmallCover) // 转本服务器
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
	r.Cover = static.ConverURL(d.Cover) // 转本服务器
	r.Id = int32(d.Id)
	r.LastUpdatetime = int32(d.LastUpdatetime)
	r.Num = int32(d.Num)
	r.Status = d.Status
	r.Title = d.Title
	r.Types = d.Types

	return r
}

func convertComicDetail(d *comicsrvpb.ComicDetail) *pb.ComicDetail {
	r := new(pb.ComicDetail)

	r.ComicPy = d.ComicPy
	r.Cover = static.ConverURL(d.Cover) // 转本服务器
	r.Description = d.Description
	r.Direction = int32(d.Direction)
	r.FirstLetter = d.FirstLetter
	r.HitNum = int32(d.HitNum)
	r.HotNum = int32(d.HotNum)
	r.Id = int32(d.Id)
	r.Islong = int32(d.Islong)
	r.LastUpdateChapterId = int32(d.LastUpdateChapterId)
	r.LastUpdateChapterName = d.LastUpdateChapterName
	r.LastUpdatetime = int32(d.LastUpdatetime)
	r.SubscribeNum = int32(d.SubscribeNum)
	r.Title = d.Title

	return r
}

func convertComicChapters(ss []*comicsrvpb.ChapterDetail) []*pb.Chapters {
	m := map[string][]*comicsrvpb.ChapterDetail{}
	for _, v := range ss {
		m[v.Title] = append(m[v.Title], v)
	}

	ds := make([]*pb.Chapters, 0, len(m))
	for k, v := range m {
		d := new(pb.Chapters)
		d.Title = k
		for _, v1 := range v {
			d1 := new(pb.Chapter)
			d1.ChapterId = v1.Chapterid
			d1.ChapterOrder = v1.Chapterorder
			d1.ChapterTitle = v1.Chaptertitle
			d1.Filesize = v1.Filesize
			d1.Updatetime = v1.Updatetime

			d.Data = append(d.Data, d1)
		}

		ds = append(ds, d)
	}

	return ds
}
