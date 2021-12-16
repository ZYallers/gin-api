package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"src/config/env"
	"src/libraries/helper"
	"src/libraries/logger"
	"strings"
	"time"
)

// LoggerWithZap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
func LoggerWithZap(zl *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		go func(ctx *gin.Context, runtime time.Duration) {
			if len(ctx.Errors) > 0 {
				reqStr := ctx.GetString(env.ReqStrKey)
				for _, err := range ctx.Errors.Errors() {
					zl.Error(err)
					helper.PushContextMessage(ctx, err, reqStr, "", true)
				}
			}
			// WebSocket 不做超时日志记录
			if ctx.GetHeader("Sec-Websocket-Key") != "" {
				return
			}
			if runtime >= env.LogMaxTimeout {
				logger.Use("timeout").Info(ctx.Request.URL.Path,
					zap.Duration("runtime", runtime),
					zap.String("proto", ctx.Request.Proto),
					zap.String("method", ctx.Request.Method),
					zap.String("host", ctx.Request.Host),
					zap.String("url", ctx.Request.URL.String()),
					zap.String("query", ctx.Request.URL.RawQuery),
					zap.String("clientIP", helper.ClientIP(ctx.ClientIP())),
					zap.Any("header", ctx.Request.Header),
					zap.String("request", ctx.GetString(env.ReqStrKey)),
				)
			}
			if runtime > env.SendMaxTimeout {
				msg := fmt.Sprintf("%s take %s to response, exceeding the maximum %s limit",
					ctx.Request.URL.Path, runtime, env.SendMaxTimeout)
				helper.PushContextMessage(ctx, msg, ctx.GetString(env.ReqStrKey), "", false)
			}
		}(ctx.Copy(), time.Now().Sub(start))
	}
}

// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
func RecoveryWithZap(zl *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				errMsg := fmt.Sprintf("recovery from panic: %v", err)
				reqStr := ctx.GetString(env.ReqStrKey)
				stacks := string(debug.Stack())

				helper.PushContextMessage(ctx, errMsg, reqStr, stacks, true)

				if brokenPipe {
					zl.Error(errMsg, zap.String("request", reqStr), zap.String("stack", stacks))
					// If the connection is dead, we can't write a status to it.
					_ = ctx.Error(err.(error)) // nolint: errcheck
					ctx.Abort()
					return
				}

				if gin.IsDebugging() {
					ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
						"code":    http.StatusInternalServerError,
						"msg":     errMsg,
						"request": strings.Split(reqStr, "\r\n"),
						"stack":   strings.Split(strings.ReplaceAll(stacks, "\t", ""), "\n"),
					})
				} else {
					zl.Error(errMsg, zap.String("request", reqStr), zap.String("stack", stacks))
					ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": errMsg})
				}
			}
		}()

		reqBytes, _ := httputil.DumpRequest(ctx.Request, true)
		ctx.Set(env.ReqStrKey, string(reqBytes))
		ctx.Next()
	}
}
