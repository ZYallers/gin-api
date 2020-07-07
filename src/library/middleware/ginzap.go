package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	app "src/config"
	"src/library/logger"
	"src/library/tool"
	"strings"
	"time"
)

const (
	logMaxSecond = 3 * time.Second
	msgMaxSecond = 10 * time.Second
)

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
func LoggerWithZap(zl *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		go func(ctx *gin.Context, runtime time.Duration) {
			if len(ctx.Errors) > 0 {
				reqStr := ctx.GetString(app.ReqStrKey)
				for _, err := range ctx.Errors.Errors() {
					zl.Error(err)
					tool.PushContextMessage(ctx, err, reqStr, "", true)
				}
			}
			if runtime >= logMaxSecond {
				logger.Use("timeout").Info(ctx.Request.URL.Path,
					zap.Duration("runtime", runtime),
					zap.String("proto", ctx.Request.Proto),
					zap.String("method", ctx.Request.Method),
					zap.String("host", ctx.Request.Host),
					zap.String("url", ctx.Request.URL.String()),
					zap.String("query", ctx.Request.URL.RawQuery),
					zap.String("clientIP", tool.ClientIP(ctx.ClientIP())),
					zap.Any("header", ctx.Request.Header),
					zap.String("request", ctx.GetString(app.ReqStrKey)),
				)
			}
			if runtime > msgMaxSecond {
				msg := fmt.Sprintf("%s take %s to response, exceeding the maximum %s limit", strings.TrimLeft(ctx.Request.URL.Path, "/"), runtime, msgMaxSecond)
				tool.PushContextMessage(ctx, msg, ctx.GetString(app.ReqStrKey), "", false)
			}
		}(ctx.Copy(), time.Now().Sub(start))
	}
}

// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
func RecoveryWithZap(zl *zap.Logger, stack bool) gin.HandlerFunc {
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
				reqStr := ctx.GetString(app.ReqStrKey)
				stacks := string(debug.Stack())
				tool.PushContextMessage(ctx, errMsg, reqStr, stacks, true)

				if brokenPipe {
					zl.Error(errMsg, zap.String("request", reqStr), zap.String("stack", stacks))
					// If the connection is dead, we can't write a status to it.
					_ = ctx.Error(err.(error)) // nolint: errcheck
					ctx.Abort()
					return
				}

				if stack {
					var buf bytes.Buffer
					ctx.Header("Content-Type", "text/html;charset=utf-8")
					buf.WriteString(`<pre style="font-family:Consolas,Menlo,monospace;line-height:1.5em;font-size:12px">`)
					buf.WriteString("<h1>" + errMsg + "</h1><h2>stack: </h2><p>")
					buf.WriteString(stacks)
					buf.WriteString("</p><h2>request: </h2><p>" + reqStr + "</p></pre>")
					ctx.String(http.StatusInternalServerError, buf.String())
					ctx.Abort()
				} else {
					zl.Error(errMsg, zap.String("request", reqStr), zap.String("stack", stacks))
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "server internal error"})
				}
			}
		}()
		
		reqBytes, _ := httputil.DumpRequest(ctx.Request, true)
		ctx.Set(app.ReqStrKey, string(reqBytes))
		ctx.Next()
	}
}
