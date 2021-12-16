package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"src/config/env"
	"src/libraries/core"
	"src/libraries/handler"
	"src/libraries/helper"
	"src/libraries/logger"
	"src/libraries/middleware"
	"src/module"
)

type router struct {
	engine *gin.Engine
	logger *zap.Logger
}

func NewRouter(engine *gin.Engine, logger *zap.Logger) *router {
	return &router{engine: engine, logger: logger}
}

// 全局中间件注册
func (r *router) CoreMiddleware() {
	r.engine.Use(
		middleware.RecoveryWithZap(r.logger),
		middleware.LoggerWithZap(r.logger),
	)
}

// 404 route handler
func (r *router) NoRouteHandler() {
	r.engine.NoRoute(func(ctx *gin.Context) {
		go push404Handler(ctx.Copy())
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusNotFound, "msg": "page not found"})
	})
}

func (r *router) ModuleHandler() {
	// Module handler
	modules := []core.IModule{
		&module.Test{},
	}
	for _, mod := range modules {
		mod.Group(r.engine)
	}
}

// 服务健康检查接口
func (r *router) HealthCheckHandler() {
	r.engine.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, `ok`)
	})
}

// expvar 服务状态信息
func (r *router) ExpVarHandler() {
	r.engine.GET("/expvar", handler.ExpHandler)
}

// prometheus 为grafana监控提供数据
func (r *router) PrometheusHandler() {
	r.engine.GET("/metrics", handler.PromHandler)
}

// statsviz 服务状态实时监控
func (r *router) StatsVizHandler() {
	r.engine.Group("/statsviz",
		gin.BasicAuth(gin.Accounts{"admin": "123456"})).
		GET("/*filepath", handler.StatsHandler)
}

// push404Handler
func push404Handler(ctx *gin.Context) {
	defer helper.SafeDefer()
	reqStr := ctx.GetString(env.ReqStrKey)
	path := ctx.Request.URL.Path
	logger.Use("404").Info(path,
		zap.String("proto", ctx.Request.Proto),
		zap.String("method", ctx.Request.Method),
		zap.String("host", ctx.Request.Host),
		zap.String("url", ctx.Request.URL.String()),
		zap.String("query", ctx.Request.URL.RawQuery),
		zap.String("clientIP", helper.ClientIP(ctx.ClientIP())),
		zap.Any("header", ctx.Request.Header),
		zap.String("request", reqStr),
	)
	helper.PushContextMessage(ctx, path+" page not found", reqStr, "", false)
}
