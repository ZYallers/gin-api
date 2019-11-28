package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type VisibleUserController struct {
	abs.Controller
}

const visibleUserUri = "http://base.hxsapp.com/base/VisibleUser/"

func VisibleUser() VisibleUserController {
	c := VisibleUserController{}
	c.Config = map[string]abs.MethodConfig{
		"Check": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c VisibleUserController) Check(ctx *gin.Context) {
	c.ServiceRewrite(ctx, visibleUserUri+tool.CurrentMethodName())
}
