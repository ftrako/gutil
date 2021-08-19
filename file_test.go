package goutils

import (
	"fmt"
	"testing"
)

func TestCreateDir(t *testing.T) {
	err := CreateDir("abc/abc2")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestCreateFile(t *testing.T) {
	err := CreateFile("abc/abc.txt")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
