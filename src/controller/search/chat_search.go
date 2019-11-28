package search

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ChatSearchController struct {
	abs.Controller
}

const chatSearchUri = "http://search.hxsapp.com/api/ChatSearch/"

func ChatSearch() ChatSearchController {
	c := ChatSearchController{}
	c.Config = map[string]abs.MethodConfig{
		"DoSearch":                {ControllerNameFirstUpper: true},
		"GetBrmServiceIdByUserId": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * @Author:      lingjijian
 * @DateTime:    2018-08-1 15:10:32
 * @Description: 搜索聊天记录
 */
func (c ChatSearchController) DoSearch(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatSearchUri+tool.CurrentMethodName())
}

/**
 * @Author:      lingjijian
 * @DateTime:    2018-08-1 15:10:32
 * @Description: 根据客户id获取聊天过的 顾问id数组
 */
func (c ChatSearchController) GetBrmServiceIdByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatSearchUri+tool.CurrentMethodName())
}
