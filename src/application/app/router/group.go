package router

import (
	"application/app/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type router struct {
	engine *gin.Engine
	logger *zap.Logger
}

func New(engine *gin.Engine, logger *zap.Logger) *router {
	return &router{engine: engine, logger: logger}
}

func (router *router) Use() *router {
	router.engine.Use(
		middleware.RecoveryWithZap(router.logger, false),
		gzip.Gzip(gzip.DefaultCompression),
		middleware.VersionControl(router.engine, Api),
	)
	return router
}

func (router *router) Static() *router {
	router.engine.Static("/assets", "./assets")
	router.engine.StaticFile("/favicon.ico", "./assets/favicon.ico")
	return router
}

func (router *router) Group() *router {
	VersionGroup(router.engine, middleware.LoggerWithZap(router.logger), middleware.Recycle(router.logger))
	return router
}
