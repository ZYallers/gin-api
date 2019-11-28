package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type SpikeController struct {
	abs.Controller
}

const spikeUri = "http://mall.hxsapp.com/base/Spike/"

func Spike() SpikeController {
	c := SpikeController{}
	c.Config = map[string]abs.MethodConfig{
		"DailySpike": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c SpikeController) DailySpike(ctx *gin.Context) {
	c.ServiceRewrite(ctx, spikeUri+tool.CurrentMethodName())
}
