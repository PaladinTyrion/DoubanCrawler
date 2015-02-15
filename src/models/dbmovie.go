package models

import (
	"time"
)

type SimpleMovieInfo struct {
	MovieId   uint64    `sql:"type:bigint(20)" gorm:"primary_key;column:movieId"`
	MovieName string    `sql:"type:varchar(80);not null" gorm:"column:movieName"`
	UpdatedAt time.Time `sql:"type:datetime(6);not null" gorm:"column:updatedAt"`
}

func (smi SimpleMovieInfo) TableName() string {
	return "simpleMovieInfo"
}
