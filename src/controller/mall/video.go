package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type VideoController struct {
	abs.Controller
}

const videoUri = "http://mall.hxsapp.com/api/Video/"

func Video() VideoController {
	c := VideoController{}
	c.Config = map[string]abs.MethodConfig{
		"Index": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c VideoController) Index(ctx *gin.Context) {
	c.ServiceRewrite(ctx, videoUri+tool.CurrentMethodName())
}
