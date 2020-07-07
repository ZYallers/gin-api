package app

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"go.uber.org/zap"
	"time"
)

const (
	Name                  = "gin-api"
	HttpServerDefaultAddr = "0.0.0.0:9010"
	HttpServerIdleTimeout = 30 * time.Second
	LogDir                = "/apps/logs/go/gin-api"
	ErrorRobotToken       = ""
	GracefulRobotToken    = ""
)

var (
	HttpServerAddr   *string
	Engine           *gin.Engine
	Logger           *zap.Logger
	RobotEnable      bool
	DebugStack       bool
	ResetXForwardFor bool //将 ctx.ClientIP() 设为 x-forward-for, 以便通过 nginx 日志排查问题
	Json             = jsoniter.ConfigCompatibleWithStandardLibrary
)
