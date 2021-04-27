package models

type UserAccount struct {
	Id       int64  `gorm:"primaryKey;autoIncrement;not null"`
	UserName string `gorm:"uniqueIndex;not null;type:varchar(255)"`
	Password string `gorm:"not null;type:varchar(255)"`
	AddTime  int64
	Status   int8
}
