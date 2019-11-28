package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type SlimPageController struct {
	abs.Controller
}

const slimPageUri = "http://slim.hxsapp.com/base/SlimPage/"

func SlimPage() SlimPageController {
	c := SlimPageController{}
	c.Config = map[string]abs.MethodConfig{
		"GetSlimIndex":      {ControllerNameFirstUpper: true},
		"GetWeightIndex":    {ControllerNameFirstUpper: true},
		"GetNewSlimIndex":   {ControllerNameFirstUpper: true},
		"GetNewWeightIndex": {ControllerNameFirstUpper: true},
		"GetSlimWeekReport": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 3.9.1
 */
func (c SlimPageController) GetSlimIndex(ctx *gin.Context) {
	c.ServiceRewrite(ctx, slimPageUri+tool.CurrentMethodName())
}

func (c SlimPageController) GetWeightIndex(ctx *gin.Context) {
	c.ServiceRewrite(ctx, slimPageUri+tool.CurrentMethodName())
}

/**
 * 3.9.8
 */
func (c SlimPageController) GetNewSlimIndex(ctx *gin.Context) {
	c.ServiceRewrite(ctx, slimPageUri+tool.CurrentMethodName())
}

func (c SlimPageController) GetNewWeightIndex(ctx *gin.Context) {
	c.ServiceRewrite(ctx, slimPageUri+tool.CurrentMethodName())
}

/**
 * 4.0.0
 */
func (c SlimPageController) GetSlimWeekReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, slimPageUri+tool.CurrentMethodName())
}
