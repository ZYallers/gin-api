package constant

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	Name                  = "api.gin.com"
	Version               = "1.1.0"
	HttpServerAddr        = "127.0.0.1:8087"
	HttpServerIdleTimeout = 30 * time.Second
	RunMode               = gin.DebugMode
	LogDir                = "./log"
	DbDebug               = true
)
