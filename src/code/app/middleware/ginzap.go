package middleware

import (
	"bytes"
	"code/app/logger"
	"code/app/tool"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
func LoggerWithZap(appLogger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// Some evil middlewares modify this values
		//path := c.Request.URL.Path
		//query := c.Request.URL.RawQuery
		c.Next()

		latency := time.Now().Sub(start)

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, err := range c.Errors.Errors() {
				go appLogger.Error(err)
				go tool.SendDingTalkGroupMessage(c, err)
			}
		} else {
			if latency >= 3*time.Second {
				go logger.Info("morethan3seconds", c.Request.URL.Path,
					zap.Int("status", c.Writer.Status()),
					zap.String("method", c.Request.Method),
					zap.String("query", c.Request.URL.RawQuery),
					zap.Duration("latency", latency),
					zap.String("ip", c.ClientIP()),
					zap.String("agent", c.Request.UserAgent()),
				)
			}
		}
	}
}

// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
func RecoveryWithZap(appLogger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				errMsg := err.(error).Error()
				go tool.SendDingTalkGroupMessage(c, errMsg)

				httpRequest, _ := httputil.DumpRequest(c.Request, true)
				httpReqStr := string(httpRequest)

				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					go appLogger.Error(c.Request.URL.Path, zap.String("error", errMsg), zap.String("request", httpReqStr))
					return
				}

				if stack {
					var bf bytes.Buffer
					bf.WriteString("[" + errMsg + "]\n\t")
					bf.Write(debug.Stack())
					bf.WriteString("[request]: \n\t" + httpReqStr)
					go appLogger.Error(bf.String())

					bf.Reset()
					c.Header("Content-Type", "text/html;charset=utf-8")
					bf.WriteString(`<pre style="font-family:SFMono-Regular,Consolas,Menlo,monospace;line-height:1.5em;font-size:14px"><h1>Error</h1><h2>` + errMsg + "</h2><p>")
					bf.Write(debug.Stack())
					bf.WriteString("</p><h2>Request: </h2><p>" + httpReqStr + "</p></pre>")
					c.String(http.StatusInternalServerError, bf.String())
				} else {
					go appLogger.Error("[Recovery From Panic]", zap.String("error", errMsg), zap.String("request", httpReqStr))
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
