package gutils

import (
	"time"

	"github.com/ftrako/glog"

	"github.com/gin-gonic/gin"
)

// 需要ShouldBindBodyWith绑定参数

// LogStat 记录日志
func LogStat(c *gin.Context, t time.Time) {
	body := ""
	if b, exists := c.Get(gin.BodyBytesKey); exists {
		if cbb, ok := b.([]byte); ok {
			body = string(cbb)
		}
	}
	if body == "" {
		body = c.Request.Form.Encode()
	}

	glog.FuncDepth(1).D("request: url:%v client_ip:%v body:%v cost:%v", c.Request.RequestURI, c.ClientIP(), body, time.Since(t))
}
