package model

type ComicDetail struct {
	ID                    int    `json:"Id" gorm:"primaryKey;autoIncrement;not null"`
	Title                 string `json:"Title"`
	Direction             int    `json:"Direction"`
	Islong                int    `json:"Islong"`
	Cover                 string `json:"Cover"`
	Description           string `json:"Description"`
	Lastupdatetime        int    `json:"LastUpdatetime"`
	Lastupdatechaptername string `json:"LastUpdateChapterName"`
	Firstletter           string `json:"FirstLetter"`
	Comicpy               string `json:"ComicPy"`
	Hotnum                int    `json:"HotNum"`
	Hitnum                int    `json:"HitNum"`
	Lastupdatechapterid   int    `json:"LastUpdateChapterId"`
	Subscribenum          int    `json:"SubscribeNum"`
}

type CategoryDetail struct {
	ID             int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title          string `json:"title"`
	Authors        string `json:"authors"`
	Status         string `json:"status"`
	Cover          string `json:"cover"`
	Types          string `json:"types"`
	LastUpdatetime int    `json:"last_updatetime"`
	Num            int    `json:"num"`
}
