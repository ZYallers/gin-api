package bbs

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CircleController struct {
	abs.Controller
}

const circleUri = "http://bbs.hxsapp.com/circle/"

func Circle() CircleController {
	c := CircleController{}
	return c
}

func (c CircleController) Icon(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleUri+tool.CurrentMethodName())
}

func (c CircleController) Detail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleUri+tool.CurrentMethodName())
}
