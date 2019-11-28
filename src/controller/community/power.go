package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type PowerController struct {
	abs.Controller
}

const powerUri = "http://community.hxsapp.com/common/power/"

func Power() PowerController {
	c := PowerController{}
	c.Config = map[string]abs.MethodConfig{
		"SaveUserGagLog": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 *  3.3.0 版本 禁言保存
 */
func (c PowerController) SaveUserGagLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, powerUri+tool.CurrentMethodName())
}
