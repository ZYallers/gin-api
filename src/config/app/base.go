package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	Name                  = "gin-api"
	HttpServerDefaultAddr = "0.0.0.0:9010"
	LogDir                = "/apps/logs/go/gin-api"
)

var (
	HttpServerAddr *string
	Engine         *gin.Engine
	Logger         *zap.Logger
)
