package service

import (
	pb "comic-service/api/grpc/v1"
	"comic-service/model"
)

func newCategoryDetail(mos []model.CategoryDetail) []*pb.CategoryDetail {
	res := make([]*pb.CategoryDetail, 0, len(mos))
	for _, v := range mos {
		c := new(pb.CategoryDetail)

		c.Authors = v.Authors
		c.Cover = v.Cover
		c.Id = int64(v.ID)
		c.LastUpdatetime = int64(v.LastUpdatetime)
		c.Num = int64(v.Num)
		c.Status = v.Status
		c.Title = v.Title
		c.Types = v.Types

		res = append(res, c)
	}

	return res
}

func newComicDetail(mos []*model.ComicDetail) []*pb.ComicDetail {
	res := make([]*pb.ComicDetail, 0, len(mos))
	for _, v := range mos {
		c := new(pb.ComicDetail)

		c.ComicPy = v.Comicpy
		c.Cover = v.Cover
		c.Description = v.Description
		c.Direction = int64(v.Direction)
		c.FirstLetter = v.Firstletter
		c.HitNum = int64(v.Hitnum)
		c.HotNum = int64(v.Hotnum)
		c.Id = int64(v.ID)
		c.Islong = int64(v.Islong)
		c.LastUpdateChapterId = int64(v.Lastupdatechapterid)
		c.LastUpdateChapterName = v.Lastupdatechaptername
		c.LastUpdatetime = int64(v.Lastupdatetime)
		c.SubscribeNum = int64(v.Subscribenum)
		c.Title = v.Title

		res = append(res, c)
	}

	return res
}
