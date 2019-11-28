package account

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
		"GetRolePowerList": {ControllerNameFirstUpper: true},
		"SavePowerLog":     {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 获取角色权限列表
 */
func (c PowerController) GetRolePowerList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, powerUri+tool.CurrentMethodName())
}

/**
 * 保存操作权限日志
 */
func (c PowerController) SavePowerLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, powerUri+tool.CurrentMethodName())
}

/**
 * 获取用户最后一条操作记录
 */
func (c PowerController) GetLastPowerLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, powerUri+tool.CurrentMethodName())
}
