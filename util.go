package goutil

import (
	"io/ioutil"
	"strings"
)

func Trim(str string) string {
	return strings.Trim(str, "\r\n\t ")
}

func Str2Arr(str string) []string {
	words := make([]string, 0)
	segs := strings.Split(str, ",")
	for _, seg := range segs {
		s := Trim(seg)
		if s != "" {
			words = append(words, s)
		}
	}
	return words
}

func File2Str(fn string) string {
	f, err := ioutil.ReadFile(fn)
	if err != nil {
		return ""
	}
	return Trim(string(f))
}
