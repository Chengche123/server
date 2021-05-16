package model

import (
	comicModel "crawler/model/comic"
	specialModel "crawler/model/special"
)

type ComicDetail struct {
	comicModel.ComicDetail
}

type CategoryDetail struct {
	comicModel.CategoryDetail
}

type ComicCategoryFilter struct {
	comicModel.ComicCategoryFilter
}

type ComicSpecial struct {
	specialModel.ComicSpecial
}
