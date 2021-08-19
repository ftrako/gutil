package goutils

import (
	"os/exec"
)

// 执行脚本

// 执行shell命令
func ExecCmd(cmd string) (string, error) {
	c := exec.Command("/bin/bash", "-c", cmd)
	bytes, err := c.Output()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
