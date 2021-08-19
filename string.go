package goutils

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Len(str string) int {
	r := []rune(str)
	return len(r)
}

// end 为-1时表示自动计算字符串长度
func SubStr(str string, start, end int) string {
	r := []rune(str)
	l := len(r)
	if start < 0 || start >= l {
		return ""
	}
	if end == -1 {
		return string(r[start:])
	}

	if end < 0 || end > l || start > end {
		return ""
	}
	return string(r[start:end])
}

// 丢掉转义词
// func DiscardEscape(str string) string {
//     str = strings.Replace(str, "'", "", -1)
//     str = strings.Replace(str, "\"", "", -1)
//     str = strings.Replace(str, "\t", "", -1)
//     str = strings.Replace(str, "\r", "", -1)
//     str = strings.Replace(str, "\\", "", -1)
//     return str
// }

// 特殊字符编码，以支持入库
func EscapeEncodeString(str string) string {
	str = strings.Replace(str, "\\", "\\\\", -1)
	str = strings.Replace(str, "'", "\\'", -1)
	return str
}

// 特殊字符还原
func EscapeDecodeString(str string) string {
	str = strings.Replace(str, "\\'", "'", -1)
	str = strings.Replace(str, "\\\\", "\\", -1)
	return str
}

// 长度设限
func MaxLen(str string, max int) string {
	if len(str) > max {
		return str[:max]
	}
	return str
}

func ConvertGBK2UTF8(b []byte) (string, error) {
	reader := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

func ConvertUTF82GBK(str string) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewEncoder())
	return ioutil.ReadAll(reader)
}

func Str2Int(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func Str2Int64(str string) int64 {
	n, _ := strconv.ParseInt(str, 10, 64)
	return n
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			return false
		}
	}
	return true
}
