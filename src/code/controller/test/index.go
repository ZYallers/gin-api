package test

import (
	"code/app/abst"
	"code/app/logger"
	"code/app/tool"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	abst.Controller
}

func Index() IndexController {
	i := IndexController{}
	i.Config = map[string]abst.MethodConfig{
		"Isok": {HttpMethods: []string{http.MethodGet}},
	}
	return i
}

func (i IndexController) Isok(ctx *gin.Context) {
	logger.Debug("", ctx.Request.RequestURI, zap.String("client_ip", ctx.ClientIP()))
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": gin.H{
			"mode":      gin.Mode(),
			"public_ip": tool.PublicIP(),
			"system_ip": tool.SystemIP(),
			"client_ip": ctx.ClientIP(),
		},
	})
}
