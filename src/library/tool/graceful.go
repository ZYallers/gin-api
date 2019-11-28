package tool

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"src/config"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func sendMessage(msg string) {
	if !app.DingTalkMsgEnable {
		return
	}
	host, _ := os.Hostname()
	text := []string{
		msg + "\n------------------",
		"AppName: " + app.Name,
		"RunMode: " + gin.Mode(),
		"ListenAddr: " + *app.HttpServerAddr,
		"HostName: " + host,
		"PublicIP: " + PublicIP(),
		"SystemIP: " + SystemIP(),
		"NowTime: " + time.Now().Format("2006/01/02 15:04:05.000000"),
	}
	posts := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": strings.Join(text, "\n") + "\n",
		},
		"at": map[string]interface{}{
			"isAtAll": true,
		},
	}
	_, _ = NewRequest(app.DingTalkGroupRobot).SetHeaders(map[string]string{"Content-Type": "application/json;charset=utf-8"}).SetPostData(posts).Post()
}

func Graceful(srv *http.Server, zl *zap.Logger, timeout time.Duration) {
	pid := strconv.Itoa(syscall.Getpid())
	if gin.IsDebugging() {
		zl.Info("Server is ready to listen and serve", zap.String("pid", pid))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zl.Fatal("Server could not listen and serve", zap.Error(err))
		}
	} else {
		done := make(chan bool, 1)
		quit := make(chan os.Signal, 1)

		// kill（无参数）默认发送是 syscall.SIGTERM
		// kill -2 是 syscall.SIGINT
		// kill -9 是 syscall.SIGKILL 但不能抓住，所以不需要添加它
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

		go func() {
			<-quit
			zl.Info("Server is shutting down...", zap.String("pid", pid))
			sendMessage("Server is shutting down...\nProcess id: " + pid)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			// SetKeepAlivesEnabled 控制是否启用HTTP保持活动，默认情况下始终启用保持活动，只有资源受限的环境或服务器在关闭过程中才应禁用它们
			srv.SetKeepAlivesEnabled(false)
			if err := srv.Shutdown(ctx); err != nil {
				sendMessage("Server could not gracefully shutdown\n" + err.Error())
				zl.Fatal("Server could not gracefully shutdown", zap.Error(err))
			}

			close(done)
		}()

		zl.Info("Server is ready to listen and serve", zap.String("pid", pid))
		sendMessage("Server is ready to listen and serve\nProcess id: " + pid)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sendMessage("Server could not listen and serve\n" + err.Error())
			zl.Fatal("Server could not listen and serve", zap.Error(err))
		}

		<-done
		zl.Info("Server has stopped", zap.String("pid", pid))
		sendMessage("Server has stopped\nProcess id: " + pid)
	}
}
