package main

import (
	. "DoubanCrawler/src/models"
	"DoubanCrawler/src/util/errutil"
	"DoubanCrawler/src/util/timeutil"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
)

func ExampleScrape() {
	//get webpage
	doc, err := goquery.NewDocument("http://movie.douban.com/subject/25884822/")
	errutil.Checkerr(err)

	//connect database
	db, err := gorm.Open("mysql", "root:32167@/dbMovie?charset=utf8&parseTime=True")
	errutil.Checkerr(err)
	db.DB()
	db.AutoMigrate(&SimpleMovieInfo{})

	//defer close db
	defer func() {
		err := db.Close()
		errutil.Checkerr(err)
	}()

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

		//for name
		title := s.Find("dd").Text()
		title = strings.TrimSpace(title)

		//		log.Printf("Remm %d: %s - %s\n", i, link, title)

		//for updatedAt
		updatedAt := time.Now().Format(timeutil.STANDARDTIME)

		//construct data && update db
		movie := SimpleMovieInfo{MovieId: movieIdN, MovieName: title, UpdatedAt: updatedAt}

		if notexist := db.First(&SimpleMovieInfo{MovieId: movieIdN}).RecordNotFound(); notexist {
			db.Create(&movie)
		}
	})
}

func main() {
	ExampleScrape()
}
