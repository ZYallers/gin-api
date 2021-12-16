package helper

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func logAndPushMsg(zl *zap.Logger, msg string) {
	zl.Info(msg)
	PushSimpleMessage(msg, true)
}

func Graceful(srv *http.Server, zl *zap.Logger, timeout time.Duration) {
	quitChan := make(chan os.Signal, 1)

	// SIGTERM 结束程序(kill pid)(可以被捕获、阻塞或忽略)
	// SIGHUP 终端控制进程结束(终端连接断开)
	// SIGINT 用户发送INTR字符(Ctrl+C)触发
	// SIGQUIT 用户发送QUIT字符(Ctrl+/)触发
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)

	pid := syscall.Getpid()
	doneChan := make(chan bool, 1)

	go func() {
		defer close(doneChan)
		for s := range quitChan {
			logAndPushMsg(zl, fmt.Sprintf("server(%d) is shutting down(%v)...", pid, s))

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			// 控制是否启用HTTP保持活动，默认情况下始终启用保持活动，只有资源受限的环境或服务器在关闭过程中才应禁用它们
			srv.SetKeepAlivesEnabled(false)

			if err := srv.Shutdown(ctx); err != nil {
				// 关闭服务失败则重新恢复
				srv.SetKeepAlivesEnabled(true)
				logAndPushMsg(zl, fmt.Sprintf("server could not gracefully shutdown, %v", err))
				continue
			} else {
				return
			}
		}
	}()

	logAndPushMsg(zl, fmt.Sprintf("server(%d) is ready to listen and serve", pid))
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logAndPushMsg(zl, fmt.Sprintf("server could not listen and serve, %v", err))
		os.Exit(1)
	}

	<-doneChan
	logAndPushMsg(zl, fmt.Sprintf("server(%d) has stopped", pid))
}
