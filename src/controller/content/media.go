package content

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type MediaController struct {
	abs.Controller
}

const mediaUri = "http://content.hxsapp.com/article/Media/"

func Media() MediaController {
	c := MediaController{}
	return c
}

func (c MediaController) IncrMediaPlay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, mediaUri+tool.CurrentMethodName())
}
