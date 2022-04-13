package gutils

import (
	"strings"
	"testing"
	"time"
)

func TestGetZeroTime(t *testing.T) {
	now := GetZeroTime(time.Now())
	if !strings.Contains(now.String(), "00:00:00") {
		// fmt.Println(now.String())
		t.FailNow()
	}
}

func TestGetMonth(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02", "2022-10-12", time.Local)
	if GetMonth(t1) != 10 {
		t.FailNow()
	}
}
