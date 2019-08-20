package cons

import (
	"time"
)

const (
	Name                  = "gin-api"
	HttpServerDefaultAddr = "0.0.0.0:9010"
	HttpServerIdleTimeout = 30 * time.Second
	LogDir                = "/apps/logs/go/gin-api"
	DingTalkMsgEnable     = true
	DingTalkGroupRobot    = "https://oapi.dingtalk.com/robot/send?access_token=95e215f5829f7f781f934ba89f5106522cca3c9c29e9a9da2e968fded62f8633"
)

var (
	HttpServerAddr string
)
