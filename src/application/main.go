package main

import (
	"application/app/constant"
	"application/app/logger"
	"application/app/router"
	"application/app/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	gin.DisableConsoleColor()     //禁用控制台颜色
	engine := gin.New()
	logger := logger.RouterLogger()
	router := router.New(engine, logger)
	router.Use().Group()
	srv := &http.Server{
		Addr:        constant.HttpServerAddr,
		IdleTimeout: constant.HttpServerIdleTimeout,
		Handler:     engine,
	}
	tool.Graceful(srv, logger, 5*time.Second)
}
