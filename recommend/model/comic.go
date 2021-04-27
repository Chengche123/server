package model

type Comic struct {
	Id             int64
	Title          string
	Num            int64 // 热度
	Types          string
	Cover          string
	LastUpdatetime int64
	IsEnd          int8
	Authors        string
	AddTime        int64
	Status         int8
}
