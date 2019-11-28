package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CommonController struct {
	abs.Controller
}

const commonUri = "http://slim.hxsapp.com/base/common/"

func Common() CommonController {
	c := CommonController{}
	return c
}

/**
 * 饮食运动库列表
 */
func (c CommonController) GetRecords(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}
