package router

import (
	"code/app/middleware"
	"code/app/router/module"
	"code/app/tool"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"reflect"
)

type router struct {
	engine *gin.Engine
	logger *zap.Logger
}

func New(e *gin.Engine, l *zap.Logger) *router {
	return &router{engine: e, logger: l}
}

/**
 * 全局中间件注册
 */
func (r *router) Middleware() *router {
	r.engine.Use(middleware.RecoveryWithZap(r.logger, gin.IsDebugging()))
	return r
}

func (r *router) Module(rg *gin.RouterGroup, modules ...interface{}) *router {
	if len(modules) > 0 {
		for _, module := range modules {
			reflect.ValueOf(module).MethodByName("Group").Call([]reflect.Value{reflect.ValueOf(rg)})
		}
	}
	return r
}

/**
 * NoRoute adds handlers for NoRoute. It return a 404 code by default.
 */
func (r *router) noRouteHandler() {
	r.engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "The page you requested was not found", "data": nil})
	})
}

func (r *router) Group() *router {
	r.noRouteHandler()
	baseRouter := r.engine.Group("/", middleware.LoggerWithZap(r.logger))
	{
		tool.HealthCheckHttpHandler(baseRouter)
		tool.ExpvarHttpHandler(baseRouter)
		tool.MetricsHttpHandler(baseRouter)
		r.Module(baseRouter,
			module.Test(),
			module.Search(),
		)
	}
	return r
}
