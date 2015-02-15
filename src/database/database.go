package database

import (
	"DoubanCrawler/src/util/errutil"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DatabaseConn(dbtype, username, password, dbname string) (gorm.DB, error) {
	//connect database
	db, err := gorm.Open(dbtype, username+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True")
	return db, err
}

func DatabaseClose(db *gorm.DB) {
	//close database
	err := db.Close()
	errutil.Checkerr(err)
}
