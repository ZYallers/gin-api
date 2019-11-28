package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type AndroidController struct {
	abs.Controller
}

const androidUri = "http://base.hxsapp.com/base/android/"

func Android() AndroidController {
	c := AndroidController{}
	c.Config = map[string]abs.MethodConfig{
		"AndroidVersion": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c AndroidController) AndroidVersion(ctx *gin.Context) {
	c.ServiceRewrite(ctx, androidUri+tool.CurrentMethodName())
}
