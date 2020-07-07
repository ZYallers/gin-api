package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ModController struct {
	abs.Controller
}

const modUri = "http://gin-base.hxsapp.com/mod/"

func ModPush() ModController {
	return ModController{}
}

/**
 * [getNoticeById 根据id获取单条通知]
 * @return [type] [description]
 */
func (c ModController) Push(ctx *gin.Context) {
	c.ServiceRewrite(ctx, modUri+tool.CurrentMethodName())
}
