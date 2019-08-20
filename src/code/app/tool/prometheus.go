package tool

import (
	"code/app/cons"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"strings"
)

var (
	ph      = promhttp.Handler()
	appInfo = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "go_app_info",
			Help: "Now running go app information.",
		},
		[]string{"name", "cmdline"},
	)
)

func init() {
	prometheus.MustRegister(appInfo)
	appInfo.WithLabelValues(cons.Name, strings.Join(os.Args, " ")).Inc()
}

func prometheusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ph.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func MetricsHttpHandler(rg *gin.RouterGroup) {
	rg.GET("/metrics", prometheusHandler())
}
