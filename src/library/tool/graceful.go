package tool

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func Graceful(srv *http.Server, zl *zap.Logger, timeout time.Duration) {
	quitChan := make(chan os.Signal, 1)
	// kill（无参数）默认发送是 syscall.SIGTERM
	// kill -2 是 syscall.SIGINT
	// kill -9 是 syscall.SIGKILL 但不能抓住，所以不需要添加它
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	pid := strconv.Itoa(syscall.Getpid())
	doneChan := make(chan bool, 1)
	go func() {
		<-quitChan
		zl.Info("server is shutting down...", zap.String("pid", pid))
		PushSimpleMessage("server is shutting down...\nprocess id: "+pid, true)

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// SetKeepAlivesEnabled 控制是否启用HTTP保持活动，默认情况下始终启用保持活动，只有资源受限的环境或服务器在关闭过程中才应禁用它们
		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			PushSimpleMessage("server could not gracefully shutdown\n"+err.Error(), true)
			zl.Fatal("server could not gracefully shutdown", zap.Error(err))
		}
		close(doneChan)
	}()

	zl.Info("server is ready to listen and serve", zap.String("pid", pid))
	PushSimpleMessage("server is ready to listen and serve\nprocess id: "+pid, true)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		PushSimpleMessage("server could not listen and serve\n"+err.Error(), true)
		zl.Fatal("server could not listen and serve", zap.Error(err))
	}

	<-doneChan
	zl.Info("server has stopped", zap.String("pid", pid))
	PushSimpleMessage("server has stopped\nprocess id: "+pid, true)
}
