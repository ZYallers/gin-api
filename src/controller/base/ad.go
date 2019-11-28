package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type AdController struct {
	abs.Controller
}

const adUri = "http://base.hxsapp.com/base/ad/"

func Ad() AdController {
	c := AdController{}
	return c
}

/**
 * 取指定类型的广告列表
 */
func (c AdController) GetAdListByType(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adUri+tool.CurrentMethodName())
}
