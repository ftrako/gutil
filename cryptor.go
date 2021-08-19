package goutils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}
