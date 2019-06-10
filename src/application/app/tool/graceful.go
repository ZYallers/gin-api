package tool

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Graceful(srv *http.Server, logger *zap.Logger, timeout time.Duration) {
	if gin.IsDebugging() {
		logger.Info("\033[32m Server LinstenAndServing HTTP \033[0m", zap.Int("ActualPid", syscall.Getpid()))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("\033[31m Server LinstenAndServe Error \033[0m", zap.Error(err))
			os.Exit(1)
		}
	} else {
		go func() {
			// 服务连接
			fmt.Printf("\033[32m Server LinstenAndServing HTTP, ActualPid: %d\n \033[0m", syscall.Getpid())
			logger.Info("\033[32m Server LinstenAndServing HTTP \033[0m", zap.Int("ActualPid", syscall.Getpid()))
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Printf("\033[31m Server LinstenAndServe Error: %v\n \033[0m", err)
				logger.Error("\033[31m Server LinstenAndServe Error \033[0m", zap.Error(err))
				os.Exit(1)
			}
		}()

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
			fmt.Printf("\033[31m Server Shutdown Error: %v\n \033[0m", err)
			logger.Error("\033[31m Server Shutdown Error \033[0m", zap.Error(err))
			os.Exit(1)
		} else {
			fmt.Printf("\033[32m Server Is Shutdown, ActualPid: %d\n \033[0m", syscall.Getpid())
			logger.Info("\033[32m Server Is Shutdown \033[0m", zap.Int("ActualPid", syscall.Getpid()))
		}
	}
}
