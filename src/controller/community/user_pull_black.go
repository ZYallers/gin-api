package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserPullBlackController struct {
	abs.Controller
}

const userPullBlackUri = "http://community.hxsapp.com/user/userPullBlack/"

func UserPullBlack() UserPullBlackController {
	c := UserPullBlackController{}
	c.Config = map[string]abs.MethodConfig{
		"SavePullBlack": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 获取拉黑名单列表
 */
func (c UserPullBlackController) GetPullBlackList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userPullBlackUri+tool.CurrentMethodName())
}

/**
 * 保存拉黑
 */
func (c UserPullBlackController) SavePullBlack(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userPullBlackUri+tool.CurrentMethodName())
}
