package main

import (
	"DoubanCrawler/src/config"
	"DoubanCrawler/src/database"
	. "DoubanCrawler/src/models"
	"DoubanCrawler/src/util/errutil"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"time"
)

func ExampleScrape() {
	//get webpage
	doc, err := goquery.NewDocument("http://movie.douban.com/subject/25884822/")
	errutil.Checkerr(err)

	//get dbhandler
	db, err := database.DatabaseConn(config.DB_TYPE,
		config.DB_USERNAME, config.DB_PASSWORD, config.DB_DBNAME)
	errutil.Checkerr(err)
	defer database.DatabaseClose(&db)

	//get data
	db.AutoMigrate(&SimpleMovieInfo{})

	doc.Find("#recommendations .recommendations-bd dl").Each(func(i int, s *goquery.Selection) {

		//for movieId
		link, _ := s.Find("dt a").Attr("href")
		link = link[strings.Index(link, "//")+2 : strings.Index(link, "?")]

		paserUrl := strings.Split(link, "/")
		var movieId string
		var movieIdN uint64
		if len(paserUrl) >= 3 {
			movieId = paserUrl[2]
			movieIdN, err = strconv.ParseUint(movieId, 10, 64)
			errutil.Checkerr(err)
		}

		//for moviename
		title := s.Find("dd").Text()
		title = strings.TrimSpace(title)

		//for updatedAt
		updatedAt := time.Now().Local()

		//construct data && update db
		movie := SimpleMovieInfo{MovieId: movieIdN,
			MovieName: title, UpdatedAt: updatedAt}
		if notexist := db.First(&SimpleMovieInfo{MovieId: movieIdN}).RecordNotFound(); notexist {
			db.Create(&movie)
		}
	})
}

func main() {
	ExampleScrape()
}
