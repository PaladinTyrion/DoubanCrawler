package errutil

import (
	//	log "github.com/cihub/seelog"
	"log"
)

func Checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
