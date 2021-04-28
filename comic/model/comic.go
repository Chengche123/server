package model

type ComicInfo struct {
	Id             int64
	Title          string
	Num            int64
	Types          string
	Cover          string
	LastUpdatetime int64
	IsEnd          int8
	Authors        string
	AddTime        int64
	Status         int8
}
