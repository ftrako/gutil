package gutils

import (
	"fmt"
	"testing"
)

func TestCreateDir(t *testing.T) {
	err := CreateDir("/Users/chendajian/data/tmp/abc2")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestCreateFile(t *testing.T) {
	err := CreateFile("/Users/chendajian/data/tmp/abc.txt")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
