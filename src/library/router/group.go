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
		ccp := ctx.Copy()
		go func() {
			reqstr := ctx.GetString(app.ReqStrKey)
			path := ccp.Request.URL.Path
			logger.Use("404").Info(path,
				zap.String("proto", ccp.Request.Proto),
				zap.String("method", ccp.Request.Method),
				zap.String("host", ccp.Request.Host),
				zap.String("url", ccp.Request.URL.String()),
				zap.String("query", ccp.Request.URL.RawQuery),
				zap.String("clientIP", ccp.ClientIP()),
				zap.Any("header", ccp.Request.Header),
				zap.String("request", reqstr),
			)
			tool.SendDingTalkGroupMessage(ccp, strings.TrimLeft(path, "/")+" page not found", reqstr, "", false)
		}()
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "Page not found"})
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
