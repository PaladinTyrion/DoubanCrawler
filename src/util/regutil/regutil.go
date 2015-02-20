package regutil

import (
	"regexp"
)

func RegExcepByExp(src, exp string) (dst string) {
	reg := regexp.MustCompile(exp)
	byt := reg.ReplaceAll([]byte(src), []byte(""))
	dst = string(byt)
	return dst
}

func RegByExp(src, exp string) (dst string) {
	reg := regexp.MustCompile(exp)
	byt := reg.FindAll([]byte(src), -1)
	var bs []byte
	for _, b := range byt {
		bs = append(bs, b[0])
	}
	dst = string(bs)
	return dst
}
