package models

import (
	"time"
)

//store movieTagList
type MovieTagList struct {
	TagName      string    `sql:"type:varchar(10)" gorm:"primary_key;column:tagName"`
	TagUrl       string    `sql:"type:varchar(70);not null" gorm:"column:tagUrl"`
	TypeName     string    `sql:"type:varchar(10);not null" gorm:"column:typeName"`
	NumInTag     uint64    `sql:"type:int(2);not null" gorm:"column:numInTag"`
	TagUpdatedAt time.Time `sql:"type:datetime(0);not null" gorm:"column:tagUpdatedAt"`
}

func (stl MovieTagList) TableName() string {
	return "movieTagList"
}
