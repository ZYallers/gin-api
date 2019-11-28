package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserCommendController struct {
	abs.Controller
}

const userCommendUri = "http://community.hxsapp.com/user/userCommend/"

func UserCommend() UserCommendController {
	c := UserCommendController{}
	c.Config = map[string]abs.MethodConfig{
		"CheckIsCommend":        {ControllerNameFirstUpper: true},
		"GetOneCommendUserList": {ControllerNameFirstUpper: true},
		"GetCommendUserLists":   {ControllerNameFirstUpper: true},
		"GetCommendById":        {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 用户点赞
 */
func (c UserCommendController) DoCommend(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommendUri+tool.CurrentMethodName())
}

/**
 * 是否点过赞，多个用逗号分隔
 */
func (c UserCommendController) CheckIsCommend(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommendUri+tool.CurrentMethodName())
}

/**
 * 是否点过赞，取单篇最新点赞人列表
 */
func (c UserCommendController) GetOneCommendUserList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommendUri+tool.CurrentMethodName())
}

/**
 * 取多组最新点赞人列表，多个用逗号分隔
 */
func (c UserCommendController) GetCommendUserLists(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommendUri+tool.CurrentMethodName())
}

/**
 * 获取点赞数量
 * type (文章=0/评论=1/动态=2/课程=3/100=随身听音频)
 */
func (c UserCommendController) GetCommendNums(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommendUri+tool.CurrentMethodName())
}

/**
 * 根据id取点赞信息
 */
func (c UserCommendController) GetCommendById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommendUri+tool.CurrentMethodName())
}
