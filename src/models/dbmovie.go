package models

import (
	"time"
)

//store movieUrlNum
type SimpleMovieInfo struct {
	MovieId   uint64    `sql:"type:bigint(20)" gorm:"primary_key;column:movieId"`
	MovieName string    `sql:"type:varchar(80);not null" gorm:"column:movieName"`
	UpdatedAt time.Time `sql:"type:datetime(6);not null" gorm:"column:updatedAt"`
}

func (smi SimpleMovieInfo) TableName() string {
	return "simpleMovieInfo"
}

//store movieTagList
type MovieTagList struct {
	TagName  string `sql:"type:varchar(10)" gorm:"primary_key;column:tagName"`
	TypeName string `sql:"type:varchar(10);not null" gorm:"column:typeName"`
	NumInTag string `sql:"type:int(10);not null" gorm:"column:numInTag"`
}

func (stl MovieTagList) TableName() string {
	return "movieTagList"
}
