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
	timeoutRecord  = 3 * time.Second
	timeoutSendMsg = 10 * time.Second
)

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
func LoggerWithZap(zl *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		runtime := time.Now().Sub(start)
		ccp := ctx.Copy()
		go func() {
			reqstr := ctx.GetString(app.ReqStrKey)
			if len(ccp.Errors) > 0 {
				for _, err := range ccp.Errors.Errors() {
					zl.Error(err)
					tool.SendDingTalkGroupMessage(ccp, err, reqstr, "", true)
				}
			}
			if runtime >= timeoutRecord {
				logger.Use("timeout").Info(ccp.Request.URL.Path,
					zap.Duration("runtime", runtime),
					zap.String("proto", ccp.Request.Proto),
					zap.String("method", ccp.Request.Method),
					zap.String("host", ccp.Request.Host),
					zap.String("url", ccp.Request.URL.String()),
					zap.String("query", ccp.Request.URL.RawQuery),
					zap.String("clientIP", ccp.ClientIP()),
					zap.Any("header", ccp.Request.Header),
					zap.String("request", reqstr),
				)
			}
			if runtime > timeoutSendMsg {
				msg := fmt.Sprintf("%s take %v to return, too slow", strings.TrimLeft(ccp.Request.URL.Path, "/"), runtime)
				tool.SendDingTalkGroupMessage(ccp, msg, reqstr, "", false)
			}
		}()
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

				errMsg := fmt.Sprintf("%v", err)
				reqstr := ctx.GetString(app.ReqStrKey)
				debugStack := debug.Stack()
				tool.SendDingTalkGroupMessage(ctx, errMsg, reqstr, string(debugStack), true)

				if brokenPipe {
					zl.Error(ctx.Request.URL.Path, zap.String("error", errMsg),
						zap.String("request", reqstr),
						zap.String("debugStack", string(debugStack)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = ctx.Error(err.(error)) // nolint: errcheck
					ctx.Abort()
					return
				}

				if stack {
					go func() {
						var bf bytes.Buffer
						bf.WriteString(errMsg + "\r\n")
						bf.WriteString("DebugStack: \r\n")
						bf.Write(debugStack)
						bf.WriteString("Request: \r\n" + reqstr)
						zl.Error(bf.String())
					}()
					var buf bytes.Buffer
					ctx.Header("Content-Type", "text/html;charset=utf-8")
					buf.WriteString(`<pre style="font-family:Consolas,Menlo,monospace;line-height:1.5em;font-size:12px">`)
					buf.WriteString("<h1>" + errMsg + "</h1><h2>DebugStack: </h2><p>")
					buf.Write(debugStack)
					buf.WriteString("</p><h2>Request: </h2><p>" + reqstr + "</p></pre>")
					ctx.String(http.StatusInternalServerError, buf.String())
					ctx.Abort()
				} else {
					zl.Error("Recovery from panic", zap.String("error", errMsg),
						zap.String("request", reqstr),
						zap.String("debugStack", string(debugStack)),
					)
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "Server internal error"})
				}
			}
		}()
		reqbyte, _ := httputil.DumpRequest(ctx.Request, true)
		ctx.Set(app.ReqStrKey, string(reqbyte))
		ctx.Next()
	}
}
