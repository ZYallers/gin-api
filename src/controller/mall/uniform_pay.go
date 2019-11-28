package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UniformPayController struct {
	abs.Controller
}

const uniformPayUri = "http://mall.hxsapp.com/base/UniformPay/"

func UniformPay() UniformPayController {
	c := UniformPayController{}
	c.Config = map[string]abs.MethodConfig{
		"Pay":        {ControllerNameFirstUpper: true},
		"OrderQuery": {ControllerNameFirstUpper: true},
		"CheckIsPay": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c UniformPayController) Pay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, uniformPayUri+tool.CurrentMethodName())
}

func (c UniformPayController) OrderQuery(ctx *gin.Context) {
	c.ServiceRewrite(ctx, uniformPayUri+tool.CurrentMethodName())
}

func (c UniformPayController) CheckIsPay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, uniformPayUri+tool.CurrentMethodName())
}
