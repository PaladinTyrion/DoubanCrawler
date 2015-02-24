package main

import (
	"DoubanCrawler/src/config"
	"DoubanCrawler/src/database"
	. "DoubanCrawler/src/models"
	"DoubanCrawler/src/util/errutil"
	"DoubanCrawler/src/util/regutil"
	"DoubanCrawler/src/util/strutil"
	"DoubanCrawler/src/util/urlutil"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
	"time"
)

func GetPageNumFromTag(tagName string) uint64 {
	//get webpage
	doc, err := goquery.NewDocument(config.DBMLIST_ENTRANCE + tagName)
	errutil.Checkerr(err)

	//get page number
	pageNumStr := doc.Find("div.paginator").ChildrenFiltered("a").Last().Text()
	pageNumUint, err := strutil.ParseUint(pageNumStr)
	errutil.Checkerr(err)

	return pageNumUint
}

func CrawlMInfoFromTag(tagName string) {

	//get dbhandler
	db, err := database.DatabaseConn(config.DB_TYPE,
		config.DB_USERNAME, config.DB_PASSWORD, config.DB_DBNAME)
	errutil.Checkerr(err)
	defer database.DatabaseClose(&db)

	//	db.AutoMigrate(&SimpleMovieInfo{}, &MovieInfo{}, &MovieRatingInfo{})
	dbs := db.AutoMigrate(&MovieInfo{}, &MovieRatingInfo{})

	pageNum := GetPageNumFromTag(tagName)
	log.Println(pageNum)

	quit := make(chan bool, uint64(pageNum))
	for i := uint64(0); i <= pageNum; i++ {
		go func(j uint64) {
			url := urlutil.GetUrlfromTag(j, tagName)
			log.Println(url)
			CrawlMInfoFromUrl(url, tagName, dbs)
			quit <- true
		}(i)
	}

	for i := uint64(0); i <= pageNum; i++ {
		<-quit
	}

}

func CrawlMInfoFromUrl(url, tagName string, db *gorm.DB) {
	//get webpage
	doc, err := goquery.NewDocument(url)
	errutil.Checkerr(err)

	//get data
	doc.Find("div.article table div.pl2").Each(func(i int, s *goquery.Selection) {

		mBaseInfo := s.ChildrenFiltered("a")
		//for movieUrl
		movieUrl, _ := mBaseInfo.Attr("href")

		//for movieId
		movieSp := strings.Split(movieUrl, "/")
		var movieIdStr string
		if order := len(movieSp); order >= 2 {
			movieIdStr = movieSp[order-2]
		}
		movieId, err := strutil.ParseUint(movieIdStr)
		errutil.Checkerr(err)

		//for movieName
		movieName := strings.TrimSpace(mBaseInfo.Text())
		movieName = strings.Replace(movieName, "\n", "", -1)
		movieName = strings.Replace(movieName, " ", "", -1)

		mDetailInfo := s.ChildrenFiltered("p.pl")
		//for releasedAt
		releasedAt := mDetailInfo.Text()
		releasedAt = releasedAt[:10]
		releasedAt = regutil.RegExcepByExp(releasedAt, "[^0-9|-]")

		mRatingInfo := s.ChildrenFiltered("div.star")
		//for rating
		mRating := mRatingInfo.Find(".rating_nums").Text()
		movieRating, err := strutil.ParseFloat(mRating)
		errutil.Checkerr(err)

		//for usersNumRating
		usersNumRatingStr := mRatingInfo.Find(".pl").Text()
		usersNumRatingStr = regutil.RegByExp(usersNumRatingStr, "[0-9]")
		usersNumRating, err := strutil.ParseUint(usersNumRatingStr)
		errutil.Checkerr(err)

		//for updatedAt
		updatedAt := time.Now()

		//MovieInfo
		mi := MovieInfo{MovieId: movieId, MovieName: movieName, TagName: tagName,
			MovieUrl: movieUrl, ReleasedAt: releasedAt, UpdatedAt: updatedAt}
		if notexist := db.First(&MovieInfo{MovieId: movieId}).RecordNotFound(); notexist {
			db.Create(&mi)
		}
		//MovieRatingInfo
		mri := MovieRatingInfo{MovieId: movieId, MovieRating: movieRating,
			UsersNumRating: usersNumRating, TagName: tagName, UpdatedAt: updatedAt}
		if notexist := db.First(&MovieRatingInfo{MovieId: movieId}).RecordNotFound(); notexist {
			db.Create(&mri)
		}
	})
}

func main() {
	tagName := "爱情"

	CrawlMInfoFromTag(tagName)
}
