package urlutil

import (
	"DoubanCrawler/src/config"
	"strconv"
)

func GetUrlfromTag(startPage uint64, typeName string) string {
	start := strconv.FormatUint(startPage*config.PAGESIZE, 10)
	url := config.DBMLIST_ENTRANCE + typeName + "?type=T&start=" + start
	return url
}
