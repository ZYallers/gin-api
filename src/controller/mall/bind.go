package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type BindController struct {
	abs.Controller
}

const bindUri = "http://mall.hxsapp.com/base/Bind/"

func Bind() BindController {
	c := BindController{}
	c.Config = map[string]abs.MethodConfig{
		"ToBind":  {ControllerNameFirstUpper: true},
		"UnBind":  {ControllerNameFirstUpper: true},
		"AliBind": {ControllerNameFirstUpper: true},
		"AddCash": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c BindController) ToBind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bindUri+tool.CurrentMethodName())
}

func (c BindController) UnBind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bindUri+tool.CurrentMethodName())
}

func (c BindController) AliBind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bindUri+tool.CurrentMethodName())
}

func (c BindController) AddCash(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bindUri+tool.CurrentMethodName())
}
