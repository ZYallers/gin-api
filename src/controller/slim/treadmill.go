package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type TreadmillController struct {
	abs.Controller
}

const treadmillUri = "http://hardware.hxsapp.com/device/Treadmill/"

func Treadmill() TreadmillController {
	c := TreadmillController{}
	c.Config = map[string]abs.MethodConfig{
		"SaveTreadmillData":      {ControllerNameFirstUpper: true},
		"GetUserReportList":      {ControllerNameFirstUpper: true},
		"GetUserReportLogById":   {ControllerNameFirstUpper: true},
		"GetUserReportTotalInfo": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c TreadmillController) SaveTreadmillData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, treadmillUri+tool.CurrentMethodName())
}

func (c TreadmillController) GetUserReportList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, treadmillUri+tool.CurrentMethodName())
}

func (c TreadmillController) GetUserReportLogById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, treadmillUri+tool.CurrentMethodName())
}

func (c TreadmillController) GetUserReportTotalInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, treadmillUri+tool.CurrentMethodName())
}
