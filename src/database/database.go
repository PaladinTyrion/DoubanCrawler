package database

import (
	"DoubanCrawler/src/util/errutil"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/url"
)

func DatabaseConn(dbtype, username, password, dbname string) (gorm.DB, error) {
	//connect database
	db, err := gorm.Open(dbtype, username+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc="+url.QueryEscape("Asia/Shanghai"))
	return db, err
}

func DatabaseClose(db *gorm.DB) {
	//close database
	err := db.Close()
	errutil.Checkerr(err)
}
