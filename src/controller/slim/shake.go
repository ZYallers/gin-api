package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ShakeController struct {
	abs.Controller
}

const shakeUri = "http://hardware.hxsapp.com/device/Shake/"

func Shake() ShakeController {
	c := ShakeController{}
	return c
}

func (c ShakeController) SaveReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shakeUri+tool.CurrentMethodName())
}

func (c ShakeController) GetReportByDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shakeUri+tool.CurrentMethodName())
}

func (c ShakeController) DeleteReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shakeUri+tool.CurrentMethodName())
}
