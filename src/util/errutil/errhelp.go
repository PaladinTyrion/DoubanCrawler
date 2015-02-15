package errutil

import (
	"log"
)

func Checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
