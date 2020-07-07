package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"src/config"
	"src/library/logger"
	"src/library/router"
	"src/library/tool"
	"time"
)

func main() {
	app.HttpServerAddr = flag.String("http.addr", app.HttpServerDefaultAddr, "服务监控地址")
	flag.Parse()

	app.RobotEnable = true
	if os.Getenv("hxsenv") == "development" {
		app.DebugStack = true
		app.ResetXForwardFor = true
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DisableConsoleColor()
	app.Engine = gin.New()
	app.Logger = logger.AppLogger()

	rou := router.New(app.Engine, app.Logger, app.DebugStack)
	rou.GlobalMiddleware()
	rou.GlobalHandlerRegister()

	srv := &http.Server{
		Addr:        *app.HttpServerAddr,
		IdleTimeout: app.HttpServerIdleTimeout,
		Handler:     app.Engine,
	}

	tool.Graceful(srv, app.Logger, 10*time.Second)
}
