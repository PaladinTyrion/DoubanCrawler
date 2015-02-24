package strutil

import (
	"strconv"
)

func ParseUint(str string) (num uint64, err error) {
	if str == "" {
		num = uint64(0)
	} else {
		num, err = strconv.ParseUint(str, 10, 64)
	}
	return num, err
}

func ParseFloat(str string) (num float64, err error) {
	if str == "" {
		num = float64(0)
	} else {
		num, err = strconv.ParseFloat(str, 64)
	}
	return num, err
}
