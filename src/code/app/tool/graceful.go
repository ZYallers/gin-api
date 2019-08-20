package tool

import (
	"code/app/cons"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func pushDingTalkGroupMessage(msg string) {
	if !cons.DingTalkMsgEnable {
		return
	}
	host, _ := os.Hostname()
	text := []string{
		msg + "\n---------------",
		"AppName: " + cons.Name,
		"HostName: " + host,
		"PublicIP: " + PublicIP(),
		"SystemIP: " + SystemIP(),
		"HttpAddr: " + cons.HttpServerAddr,
		"RunMode: " + gin.Mode(),
		"NowTime: " + time.Now().Format("2006/01/02 15:04:05"),
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
	_, _ = NewRequest(cons.DingTalkGroupRobot).SetHeaders(map[string]string{"Content-Type": "application/json;charset=utf-8"}).SetPostData(posts).Post()
}

func Graceful(srv *http.Server, logger *zap.Logger, timeout time.Duration) {
	if gin.IsDebugging() {
		logger.Info("Server Linsten And Serving", zap.Int("ActualPid", syscall.Getpid()))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("[Server LinstenAndServe Error]", zap.Error(err))
			pushDingTalkGroupMessage("[Server LinstenAndServe Error]\n" + err.Error())
			os.Exit(1)
		}
	} else {
		go func() {
			// 服务连接
			logger.Info("Server Linsten And Serving", zap.Int("ActualPid", syscall.Getpid()))
			pushDingTalkGroupMessage("[Server Linsten And Serving]\nActualPid: " + strconv.Itoa(syscall.Getpid()))
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("[Server LinstenAndServe Error]", zap.Error(err))
				pushDingTalkGroupMessage("[Server LinstenAndServe Error]\n" + err.Error())
				os.Exit(1)
			}
		}()

		// 等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
		quit := make(chan os.Signal, 1)
		// kill（无参数）默认发送是 syscanll.SIGTERM
		// kill -2 是 syscall.SIGINT
		// kill -9 是 syscall.SIGKILL 但不能抓住，所以不需要添加它
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		//time.Sleep(3 * time.Second)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Error("[Server Shutdown Error]", zap.Error(err))
			pushDingTalkGroupMessage("[Server Shutdown Error]\n" + err.Error())
			os.Exit(1)
		} else {
			logger.Info("Server Is Shutdown", zap.Int("ActualPid", syscall.Getpid()))
			pushDingTalkGroupMessage("[Server Is Shutdown]\nActualPid: " + strconv.Itoa(syscall.Getpid()))
		}
	}
}
