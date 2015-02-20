package main

import (
	"DoubanCrawler/src/config"
	"DoubanCrawler/src/database"
	. "DoubanCrawler/src/models"
	"DoubanCrawler/src/util/errutil"
	"DoubanCrawler/src/util/regutil"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"time"
)

func CrawlerMInfoFromUrl(url, tagName string) {
	//get webpage
	doc, err := goquery.NewDocument(url)
	errutil.Checkerr(err)

	//get dbhandler
	db, err := database.DatabaseConn(config.DB_TYPE,
		config.DB_USERNAME, config.DB_PASSWORD, config.DB_DBNAME)
	errutil.Checkerr(err)
	defer database.DatabaseClose(&db)

	//get data
	//	db.AutoMigrate(&SimpleMovieInfo{}, &MovieInfo{}, &MovieRatingInfo{})
	db.AutoMigrate(&MovieInfo{}, &MovieRatingInfo{})

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
		movieId, err := strconv.ParseUint(movieIdStr, 10, 64)
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
		movieRating, err := strconv.ParseFloat(mRating, 64)
		errutil.Checkerr(err)

		//for usersNumRating
		usersNumRatingStr := mRatingInfo.Find(".pl").Text()
		usersNumRatingStr = regutil.RegByExp(usersNumRatingStr, "[0-9]")
		usersNumRating, err := strconv.ParseUint(usersNumRatingStr, 10, 64)
		errutil.Checkerr(err)

		//for updatedAt
		updatedAt := time.Now().Local()

		//construct data && update db
		//		//SimpleMovieInfo
		//		smi := SimpleMovieInfo{MovieId: movieId,
		//			MovieName: movieName, UpdatedAt: updatedAt}
		//		if notexist := db.First(&SimpleMovieInfo{MovieId: movieId}).RecordNotFound(); notexist {
		//			db.Create(&smi)
		//		}
		//MovieInfo
		mi := MovieInfo{MovieId: movieId, MovieName: movieName, TagName: tagName,
			MovieUrl: movieUrl, ReleasedAt: releasedAt, UpdatedAt: updatedAt}
		if notexist := db.First(&MovieInfo{MovieId: movieId}).RecordNotFound(); notexist {
			db.Create(&mi)
		}
		//MovieRatingInfo
		mri := MovieRatingInfo{MovieId: movieId, MovieRating: movieRating,
			UsersNumRating: usersNumRating, UpdatedAt: updatedAt}
		if notexist := db.First(&MovieRatingInfo{MovieId: movieId}).RecordNotFound(); notexist {
			db.Create(&mri)
		}
	})
}

func main() {
	url := "http://movie.douban.com/tag/爱情"
	CrawlerMInfoFromUrl(url, "爱情")
}
