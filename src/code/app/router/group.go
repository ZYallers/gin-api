package router

import (
	"code/app/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type router struct {
	engine *gin.Engine
	logger *zap.Logger
}

func New(e *gin.Engine, l *zap.Logger) *router {
	return &router{engine: e, logger: l}
}

func (r *router) Use() *router {
	r.engine.Use(
		middleware.RecoveryWithZap(r.logger, gin.IsDebugging()),
		middleware.VersionControl(r.engine, Api),
	)
	return r
}

func (r *router) Group() *router {
	VersionGroup(r.engine, middleware.LoggerWithZap(r.logger))
	return r
}
