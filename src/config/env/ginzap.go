package env

import "time"

const (
	LogMaxTimeout  = 3 * time.Second        // 超过3秒的请求记录日志
	SendMaxTimeout = 5 * time.Second        // 超过5秒的请求发送预警消息
	ReqStrKey      = "gin-gonic/gin/reqstr" // http.Request请求内容值的key
)
