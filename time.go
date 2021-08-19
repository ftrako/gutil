package goutils

import (
	"strconv"
	"time"
)

// 当天剩余秒数
func LeftSecondsInCurrentDay() int64 {
	t := time.Now()
	zero := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return 24*60*60 - (t.Unix() - zero.Unix())
}

// 日期转成数字，格式：20060102
func Date2Number(t time.Time) int {
	str := t.Format("20060102")
	n, _ := strconv.Atoi(str)
	return n
}
