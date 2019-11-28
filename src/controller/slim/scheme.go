package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type SchemeController struct {
	abs.Controller
}

const schemeUri = "http://slim.hxsapp.com/slim/Scheme/"

func Scheme() SchemeController {
	c := SchemeController{}
	return c
}

/**
 * 保存开关状态
 */
func (c SchemeController) SaveSwitchState(ctx *gin.Context) {
	c.ServiceRewrite(ctx, schemeUri+tool.CurrentMethodName())
}

/**
 * 获取开关日历
 */
func (c SchemeController) GetSwitchDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, schemeUri+tool.CurrentMethodName())
}
