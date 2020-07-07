package hardware

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ManageController struct {
	abs.Controller
}

const manageUri = "http://hardware.hxsapp.com/api/Manage/"

func Manage() ManageController {
	c := ManageController{}
	c.Config = map[string]abs.MethodConfig{
		"GetWeightRecord":   {ControllerNameFirstUpper: true},
		"GetBodyScaleQuota": {ControllerNameFirstUpper: true},
		"CheckDevBind":      {ControllerNameFirstUpper: true},
	}
	return c
}

func (c ManageController) GetWeightRecord(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manageUri+tool.CurrentMethodName())
}

func (c ManageController) GetBodyScaleQuota(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manageUri+tool.CurrentMethodName())
}

func (c ManageController) CheckDevBind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manageUri+tool.CurrentMethodName())
}
