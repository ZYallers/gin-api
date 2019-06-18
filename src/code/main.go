package main

import (
	"code/app/cons"
	"code/app/logger"
	"code/app/router"
	"code/app/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	gin.DisableConsoleColor()
	engine := gin.New()
	logger := logger.RouterLogger()
	router := router.New(engine, logger)
	router.Use().Group()
	srv := &http.Server{
		Addr:        cons.HttpServerAddr,
		IdleTimeout: cons.HttpServerIdleTimeout,
		Handler:     engine,
	}
	tool.Graceful(srv, logger, 5*time.Second)
}
