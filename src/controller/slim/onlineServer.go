package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type HwOnlineServerController struct {
	abs.Controller
}

const hwOnlineServerUri = "http://hardware.hxsapp.com/api/onlineServer/"

func HwOnlineServer() HwOnlineServerController {
	c := HwOnlineServerController{}
	return c
}

func (c HwOnlineServerController) GetHxsServiceSports(ctx *gin.Context) {
	c.ServiceRewrite(ctx, hwOnlineServerUri+tool.CurrentMethodName())
}

func (c HwOnlineServerController) SetVideoCourseClock(ctx *gin.Context) {
	c.ServiceRewrite(ctx, hwOnlineServerUri+tool.CurrentMethodName())
}
