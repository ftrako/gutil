package gutils

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestGetZeroTime(t *testing.T) {
	now := GetZeroTime(time.Now())
	if !strings.Contains(now.String(), "00:00:00") {
		fmt.Println(now.String())
		t.FailNow()
	}
}
