package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

const (
	Name                  = "gin-api"
	HttpServerDefaultAddr = "0.0.0.0:9010"
	HttpServerIdleTimeout = 30 * time.Second
	LogDir                = "/apps/logs/go/gin-api"
	DingTalkMsgEnable     = true
	DingTalkGroupRobot    = "https://oapi.dingtalk.com/robot/send?access_token=0d10a39901249e43fe66065f55449e17df7bc788617edd5dbcaf77396668be7b"
)

var (
	HttpServerAddr *string
	Engine         *gin.Engine
	Logger         *zap.Logger
	DebugStack     bool
)
