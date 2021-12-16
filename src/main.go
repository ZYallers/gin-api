package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
	"src/config/app"
	"src/config/env"
	"src/config/router"
	"src/libraries/core"
	"src/libraries/helper"
	"src/libraries/logger"
	"time"
)

func main() {
	app.HttpServerAddr = flag.String("http.addr", app.HttpServerDefaultAddr, "服务监控地址，如：0.0.0.0:9010")
	flag.Parse()

	if env.App.Env == env.DevMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DisableConsoleColor()
	app.Engine = gin.New()
	logger.SetLogDir(app.LogDir)
	app.Logger = logger.AppLogger(app.Name)

	ru := router.NewRouter(app.Engine, app.Logger)
	ru.NoRouteHandler()
	ru.CoreMiddleware()
	ru.ModuleHandler()
	ru.HealthCheckHandler()
	ru.ExpVarHandler()
	ru.PrometheusHandler()
	ru.StatsVizHandler()

	env.Session.GetClientFunc = new(core.Redis).GetSession

	srv := &http.Server{
		Addr:         *app.HttpServerAddr,
		Handler:      app.Engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	helper.Graceful(srv, app.Logger, 15*time.Second)
}
