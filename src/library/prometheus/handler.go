package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"src/config"
	"strings"
)

var (
	prom    = promhttp.Handler()
	appInfo = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "go_app_info",
			Help: "now running go app information.",
		},
		[]string{"name", "cmdline"},
	)
)

func init() {
	prometheus.MustRegister(appInfo)
	appInfo.WithLabelValues(app.Name, strings.Join(os.Args, " ")).Inc()
}

func ServerHandler(ctx *gin.Context) {
	prom.ServeHTTP(ctx.Writer, ctx.Request)
}
