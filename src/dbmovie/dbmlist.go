package main

import (
	"DoubanCrawler/src/config"
	"DoubanCrawler/src/database"
	. "DoubanCrawler/src/models"
	"DoubanCrawler/src/util/errutil"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"time"
)

const (
	entranceUrl = config.DBMLIST_ENTRANCE
)

func CrawlerMList() {
	//get webpage
	doc, err := goquery.NewDocument(entranceUrl)
	errutil.Checkerr(err)

	//get dbhandler
	db, err := database.DatabaseConn(config.DB_TYPE, config.DB_USERNAME, config.DB_PASSWORD, config.DB_DBNAME)
	errutil.Checkerr(err)
	defer database.DatabaseClose(&db)

	//get data
	db.AutoMigrate(&MovieTagList{})

	//parse webpage
	var typeName string
	selec := doc.Find("div.article")
	selecA := selec.ChildrenFiltered("a")

	selecA.Each(func(i int, s *goquery.Selection) {
		//typeName
		typeName, _ = s.Attr("name")

		if typeName == "艺术家" {

		} else {

			selecTable := s.Next()
			selecTable.Find("td").Each(func(j int, ts *goquery.Selection) {

				tagA := ts.ChildrenFiltered("a")

				//tagName
				tagName := tagA.Text()

				//tagUrl
				tagUrl, _ := tagA.Attr("href")
				if strings.HasPrefix(tagUrl, "./") {
					tagUrl = strings.Replace(tagUrl, "./", "", 1)
				}
				tagUrl = strings.Join([]string{entranceUrl, tagUrl}, "")

				//tagInNum
				tagB := ts.ChildrenFiltered("b")
				nInTag := tagB.Text()
				nInTag = nInTag[1 : len(nInTag)-1]
				numInTag, err := strconv.ParseUint(nInTag, 10, 64)
				errutil.Checkerr(err)

				//TagUpdatedAt
				tagUpdatedAt := time.Now()

				//construct data && update db
				tag4Seach := MovieTagList{TagName: tagName, TagUrl: tagUrl,
					TypeName: typeName, NumInTag: numInTag, TagUpdatedAt: tagUpdatedAt}
				if notexist := db.First(&MovieTagList{TagName: tagName, TagUrl: tagUrl}).RecordNotFound(); notexist {
					db.Create(&tag4Seach)
				}
			})
		}
	})
}

func main() {
	CrawlerMList()
}
