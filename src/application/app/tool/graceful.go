package tool

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Graceful(srv *http.Server, logger *zap.Logger, timeout time.Duration) {
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//log.Fatalf("[GIN] %v Listen: %v\n", nowTimeFormatFunc(), err)
			logger.Error("Server LinstenAndServe Error", zap.Error(err))
			os.Exit(1)
		}
	}()

	logger.Info("[************************ Server LinstenAndServing... ************************]", zap.Int("actual_pid", syscall.Getpid()))

	// 等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
	quit := make(chan os.Signal, 1)
	// kill（无参数）默认发送是 syscanll.SIGTERM
	// kill -2 是 syscall.SIGINT
	// kill -9 是 syscall.SIGKILL 但不能抓住，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown Error", zap.Error(err))
		os.Exit(1)
	} else {
		logger.Info("[************************ Server Is Shutdown ************************]", zap.Int("actual_pid", syscall.Getpid()))
	}
}
