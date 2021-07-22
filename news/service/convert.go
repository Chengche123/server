package service

import (
	pb "news-service/api/grpc/v1"
	"news-service/model"
)

func newNewsCategoryDetail(mos []model.NewsCategoryDetail) []*pb.NewsCategoryDetail {
	res := make([]*pb.NewsCategoryDetail, 0, len(mos))

	for _, v := range mos {
		d := new(pb.NewsCategoryDetail)

		d.ArticleId = int64(v.Articleid)
		d.AuthorId = int64(v.Authorid)
		d.AuthorUid = int64(v.Authoruid)
		d.ColPicUrl = v.Colpicurl
		d.CommentAmount = int64(v.Commentamount)
		d.Cover = v.Cover
		d.CreateTime = int64(v.Createtime)
		d.FromName = v.Fromname
		d.Intro = v.Intro
		d.MoodAmount = int64(v.Moodamount)
		d.Nickname = v.Nickname
		d.PageUrl = v.Pageurl
		d.RowPicUrl = v.Rowpicurl
		d.Status = int32(v.Status)
		d.TagId = int64(v.TagId)
		d.Title = v.Title

		res = append(res, d)
	}

	return res
}
