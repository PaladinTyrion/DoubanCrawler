package main

import (
	"DoubanCrawler/src/config"
	"DoubanCrawler/src/database"
	//	. "DoubanCrawler/src/models"
	"DoubanCrawler/src/util/errutil"
	"github.com/PuerkitoBio/goquery"
)

func CrawlerMList() {
	//get webpage
	_, err := goquery.NewDocument(config.DBMLIST_ENTRANCE)
	errutil.Checkerr(err)

	//get dbhandler
	db, err := database.DatabaseConn(config.DB_TYPE, config.DB_USERNAME, config.DB_PASSWORD, config.DB_DBNAME)
	errutil.Checkerr(err)
	defer database.DatabaseClose(&db)
}

func main() {
	CrawlerMList()
}
