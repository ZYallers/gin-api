package im

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"src/abs"
	"src/library/logger"
)

type ErrorController struct {
	abs.Controller
}

func Error() ErrorController {
	c := ErrorController{}
	c.Config = map[string]abs.MethodConfig{
		"Record": {HttpMethods: []string{http.MethodPost}},
	}
	return c
}

func (c ErrorController) Record(ctx *gin.Context) {
	_ = ctx.Request.ParseForm()
	logger.Use("im_error").Info("record", zap.Any("PostForm", ctx.Request.PostForm))
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "ok"})
}
