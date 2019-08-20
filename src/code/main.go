package main

import (
	"code/app/cons"
	"code/app/logger"
	"code/app/router"
	"code/app/tool"
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	cons.HttpServerAddr = *flag.String("http.addr", cons.HttpServerDefaultAddr, "Web服务监控地址")
	flag.Parse()

	gin.DisableConsoleColor()
	engine := gin.New()

	logger := logger.RouterLogger()
	router := router.New(engine, logger)
	router.Middleware().Group()

	srv := &http.Server{
		Addr:        cons.HttpServerAddr,
		IdleTimeout: cons.HttpServerIdleTimeout,
		Handler:     engine,
	}

	tool.Graceful(srv, logger, 5*time.Second)
}
