package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type SleepController struct {
	abs.Controller
}

const sleepUri = "http://slim.hxsapp.com/tool/sleep/"

func Sleep() SleepController {
	c := SleepController{}
	c.Config = map[string]abs.MethodConfig{
		"Index":    {ControllerNameFirstUpper: true},
		"SetAlarm": {ControllerNameFirstUpper: true},
		"Clock":    {ControllerNameFirstUpper: true},
	}
	return c
}

func (c SleepController) Index(ctx *gin.Context) {
	c.ServiceRewrite(ctx, sleepUri+tool.CurrentMethodName())
}

func (c SleepController) SetAlarm(ctx *gin.Context) {
	c.ServiceRewrite(ctx, sleepUri+tool.CurrentMethodName())
}

func (c SleepController) Clock(ctx *gin.Context) {
	c.ServiceRewrite(ctx, sleepUri+tool.CurrentMethodName())
}
