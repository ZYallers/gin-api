package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type NbongController struct {
	abs.Controller
}

const nbongUri = "http://hardware.hxsapp.com/device/Nbong/"

func Nbong() NbongController {
	c := NbongController{}
	c.Config = map[string]abs.MethodConfig{
		"GetCity":                 {ControllerNameFirstUpper: true},
		"Weather":                 {ControllerNameFirstUpper: true},
		"GetUserAlarmInfo":        {ControllerNameFirstUpper: true},
		"ConfigUserAlarm":         {ControllerNameFirstUpper: true},
		"GetUserStepByRecordDate": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c NbongController) GetCity(ctx *gin.Context) {
	c.ServiceRewrite(ctx, nbongUri+tool.CurrentMethodName())
}

func (c NbongController) Weather(ctx *gin.Context) {
	c.ServiceRewrite(ctx, nbongUri+tool.CurrentMethodName())
}

func (c NbongController) GetUserAlarmInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, nbongUri+tool.CurrentMethodName())
}

func (c NbongController) ConfigUserAlarm(ctx *gin.Context) {
	c.ServiceRewrite(ctx, nbongUri+tool.CurrentMethodName())
}

func (c NbongController) GetUserStepByRecordDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, nbongUri+tool.CurrentMethodName())
}
