package service

import (
	pb "comic-service/api/grpc/v1"
	"comic-service/model"
)

func newComicCategoryDetail(mos []model.CategoryDetail) []*pb.ComicCategoryDetail {
	res := make([]*pb.ComicCategoryDetail, 0, len(mos))
	for _, v := range mos {
		c := new(pb.ComicCategoryDetail)

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

func newComicCategoryFilter(mos []model.ComicCategoryFilter) []*pb.ComicCategoryFilter {
	m := make(map[string][]*pb.ComicCategoryFilterItem)

	for _, v := range mos {
		m[v.Title] = append(m[v.Title], &pb.ComicCategoryFilterItem{
			TagId:   uint32(v.TagID),
			TagName: v.TagName,
		})
	}

	res := make([]*pb.ComicCategoryFilter, 0, len(m))
	for title, items := range m {
		res = append(res, &pb.ComicCategoryFilter{
			Title: title,
			Items: items,
		})
	}

	return res
}

func newComicSpecial(mos []model.ComicSpecial) []*pb.ComicSpecial {
	res := make([]*pb.ComicSpecial, 0, len(mos))
	for _, v := range mos {
		c := new(pb.ComicSpecial)

		c.CreateTime = int64(v.CreateTime)
		c.Id = int64(v.ID)
		c.PageType = int64(v.PageType)
		c.PageUrl = v.PageURL
		c.ShortTitle = v.ShortTitle
		c.SmallCover = v.SmallCover
		c.Sort = int64(v.Sort)
		c.Title = v.Title

		res = append(res, c)
	}

	return res
}
