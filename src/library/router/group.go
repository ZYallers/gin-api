package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"src/config"
	"src/library/expvar"
	"src/library/logger"
	"src/library/middleware"
	"src/library/prometheus"
	"src/library/tool"
	"src/module"
	"strings"
)

type router struct {
	engine     *gin.Engine
	logger     *zap.Logger
	debugStack bool
}

func New(engine *gin.Engine, logger *zap.Logger, debugStack bool) *router {
	return &router{engine: engine, logger: logger, debugStack: debugStack}
}

// 全局中间件注册
func (r *router) GlobalMiddleware() {
	r.engine.Use(middleware.RecoveryWithZap(r.logger, r.debugStack), middleware.LoggerWithZap(r.logger))
}

func (r *router) Module(eg *gin.Engine, modules ...interface{}) *router {
	if len(modules) > 0 {
		for _, mode := range modules {
			reflect.ValueOf(mode).MethodByName("Group").Call([]reflect.Value{reflect.ValueOf(eg)})
		}
	}
	return r
}

// adds handlers for NoRoute. It return a 404 code by default.
func (r *router) noRouteHandlerRegister() {
	r.engine.NoRoute(func(ctx *gin.Context) {
		go func(ctx *gin.Context) {
			reqStr := ctx.GetString(app.ReqStrKey)
			path := ctx.Request.URL.Path
			logger.Use("404").Info(path,
				zap.String("proto", ctx.Request.Proto),
				zap.String("method", ctx.Request.Method),
				zap.String("host", ctx.Request.Host),
				zap.String("url", ctx.Request.URL.String()),
				zap.String("query", ctx.Request.URL.RawQuery),
				zap.String("clientIP", tool.ClientIP(ctx.ClientIP())),
				zap.Any("header", ctx.Request.Header),
				zap.String("request", reqStr),
			)
			tool.PushContextMessage(ctx, strings.TrimLeft(path, "/")+" page not found", reqStr, "", false)
		}(ctx.Copy())
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "page not found"})
	})
}

func (r *router) GlobalHandlerRegister() {
	// Module handler
	r.Module(r.engine,
		module.Api(),
		module.Account(),
		module.Base(),
		module.Search(),
		module.Mall(),
		module.Hardware(),
		module.Bonus(),
		module.Community(),
		module.Content(),
		module.Home(),
		module.Im(),
		module.Slim(),
		module.Test(),
		module.Live(),
		module.Bbs(),
	)

	// 服务健康检查
	r.engine.GET("/HealthCheck", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, `"ok"`)
	})

	// expvar 统计
	r.engine.GET("/expvar", expvar.RunningStatsHandler)

	// prometheus 统计，为granfa监控提供接口
	r.engine.GET("/metrics", prometheus.ServerHandler)

	// 404 route handler
	r.noRouteHandlerRegister()

}
