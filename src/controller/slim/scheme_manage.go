package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type SchemeManageController struct {
	abs.Controller
}

const schemeManageUri = "http://slim.hxsapp.com/lsmember/SchemeManage/"

func SchemeManage() SchemeManageController {
	c := SchemeManageController{}
	return c
}

/**
 * 明日建议
 */
func (c SchemeManageController) TwmSuggest(ctx *gin.Context) {
	c.ServiceRewrite(ctx, schemeManageUri+tool.CurrentMethodName())
}
