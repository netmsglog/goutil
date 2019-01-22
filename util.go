package goutil

import (
	"crypto/md5"
	b64 "encoding/base64"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/nu7hatch/gouuid"
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

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func File2Str(fn string) string {
	f, err := ioutil.ReadFile(fn)
	if err != nil {
		return ""
	}
	return Trim(string(f))
}

func INT(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func UUID() string {
	uu, _ := uuid.NewV4()
	return uu.String()
}

func MD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

func Segment(str string, b string, e string) string {
	pos := strings.Index(str, b)
	if pos == -1 {
		return ""
	}
	npos := pos + len(b)
	if e == "" {
		return str[npos:]
	}
	epos := strings.Index(str[npos:], e)
	if epos == -1 {
		return str[npos:]
	}
	return str[npos:(npos + epos)]
}

func Enc64(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}

func Dec64(s string) string {
	b, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		return s
	}
	return string(b)
}
