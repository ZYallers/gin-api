package live

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type PresentController struct {
	abs.Controller
}

const presentUri = "http://live.hxsapp.com/present/"

func Present() PresentController {
	c := PresentController{}
	return c
}

func (c PresentController) List(ctx *gin.Context) {
	c.ServiceRewrite(ctx, presentUri+tool.CurrentMethodName())
}

func (c PresentController) Send(ctx *gin.Context) {
	c.ServiceRewrite(ctx, presentUri+tool.CurrentMethodName())
}

func (c PresentController) Gains(ctx *gin.Context) {
	c.ServiceRewrite(ctx, presentUri+tool.CurrentMethodName())
}