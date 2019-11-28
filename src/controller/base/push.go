package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type PushController struct {
	abs.Controller
}

const pushUri = "http://base.hxsapp.com/message/Push/"

func Push() PushController {
	c := PushController{}
	c.Config = map[string]abs.MethodConfig{
		"SetAllPushSwitch": {ControllerNameFirstUpper: true},
		"GetAllPushSwitch": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c PushController) SetAllPushSwitch(ctx *gin.Context) {
	c.ServiceRewrite(ctx, pushUri+tool.CurrentMethodName())
}

func (c PushController) GetAllPushSwitch(ctx *gin.Context) {
	c.ServiceRewrite(ctx, pushUri+tool.CurrentMethodName())
}
