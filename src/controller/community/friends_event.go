package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type FriendsEventController struct {
	abs.Controller
}

const friendsEventUri = "http://community.hxsapp.com/user/friendsEvent/"

func FriendsEvent() FriendsEventController {
	c := FriendsEventController{}
	c.Config = map[string]abs.MethodConfig{
		"GetList":          {ControllerNameFirstUpper: true},
		"GetNewFollowTime": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 获取好友动态列表
 */
func (c FriendsEventController) GetList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, friendsEventUri+tool.CurrentMethodName())
}

func (c FriendsEventController) GetNewFollowTime(ctx *gin.Context) {
	c.ServiceRewrite(ctx, friendsEventUri+tool.CurrentMethodName())
}
