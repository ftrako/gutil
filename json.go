package goutils

import jsoniter "github.com/json-iterator/go"

// 解析对象，返回json字符串
func JsonObject(obj interface{}) string {
	b, _ := jsoniter.Marshal(obj)
	return string(b)
}
