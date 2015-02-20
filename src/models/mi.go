package models

import (
	"time"
)

//store movieInfo
type MovieInfo struct {
	MovieId    uint64    `sql:"type:bigint(2)" gorm:"primary_key;column:movieId"`
	MovieName  string    `sql:"type:varchar(35);not null" gorm:"column:movieName"`
	TagName    string    `sql:"type:varchar(10);not null" gorm:"column:tagName"`
	MovieUrl   string    `sql:"type:varchar(70);not null" gorm:"column:movieUrl"`
	ReleasedAt string    `sql:"type:varchar(10)" gorm:"column:releasedAt"`
	UpdatedAt  time.Time `sql:"type:datetime(0);not null" gorm:"column:updatedAt"`
}

func (mi MovieInfo) TableName() string {
	return "movieInfo"
}

type MovieRatingInfo struct {
	MovieId        uint64    `sql:"type:bigint(2)" gorm:"primary_key;column:movieId"`
	MovieRating    float64   `sql:"type:float(3)" gorm:"column:movieRating"`
	UsersNumRating uint64    `sql:"type:bigint(1)" gorm:"column:usersNumRating"`
	UpdatedAt      time.Time `sql:"type:datetime(0);not null" gorm:"column:updatedAt"`
}

func (mri MovieRatingInfo) TableName() string {
	return "movieRatingInfo"
}
