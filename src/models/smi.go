package models

import (
	"time"
)

//store simpleMovieInfo
type SimpleMovieInfo struct {
	MovieId   uint64    `sql:"type:bigint(2)" gorm:"primary_key;column:movieId"`
	MovieName string    `sql:"type:varchar(35);not null" gorm:"column:movieName"`
	UpdatedAt time.Time `sql:"type:datetime(0);not null" gorm:"column:updatedAt"`
}

func (smi SimpleMovieInfo) TableName() string {
	return "simpleMovieInfo"
}
