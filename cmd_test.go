package gutils

import (
	"fmt"
	"testing"
)

func TestExecCmd(t *testing.T) {
	str, _ := ExecCmd("pwd")
	fmt.Println("res:", str)
}
