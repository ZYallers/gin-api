package hardware

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type WeiMiController struct {
	abs.Controller
}

const weimiUri = "http://hardware.hxsapp.com/device/WeiMi/"

func WeiMi() WeiMiController {
	c := WeiMiController{}
	c.Config = map[string]abs.MethodConfig{
		"ReceiveBonus": {ControllerNameFirstUpper: true},
		"GetInfo":      {ControllerNameFirstUpper: true},
		"Exchange":     {ControllerNameFirstUpper: true},
		"Remind":       {ControllerNameFirstUpper: true},
	}
	return c
}

func (c WeiMiController) ReceiveBonus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weimiUri+tool.CurrentMethodName())
}

func (c WeiMiController) GetInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weimiUri+tool.CurrentMethodName())
}

func (c WeiMiController) Exchange(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weimiUri+tool.CurrentMethodName())
}

func (c WeiMiController) Remind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weimiUri+tool.CurrentMethodName())
}
