package goutils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 文件类

// 逐行读取文本文件
func ReadTextFile(file string, fn func(text string)) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// 逐行解析文本
	br := bufio.NewReader(f)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if fn != nil {
			fn(string(line))
		}
	}
	return nil
}

// 读取文件所有内容
func ReadFileAll(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// 创建文件
// 自动递归创建父级目录
func CreateFile(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	n := strings.LastIndex(path, "/")
	if n > 0 {
		err := CreateDir(path[:n])
		if err != nil {
			return err
		}
	}
	_, err = os.Create(path)
	return err
}

// 创建目录
// 自动递归创建父级目录
func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModeType|os.ModePerm)
}

// 文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
