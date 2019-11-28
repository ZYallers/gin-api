package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserCollectController struct {
	abs.Controller
}

const userCollectUri = "http://community.hxsapp.com/user/userCollect/"

func UserCollect() UserCollectController {
	c := UserCollectController{}
	c.Config = map[string]abs.MethodConfig{
		"CheckIsCollect": {ControllerNameFirstUpper: true},
		"GetCollectNums": {ControllerNameFirstUpper: true},
		"DeleteCollect":  {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 用户收藏
 */
func (c UserCollectController) DoCollect(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCollectUri+tool.CurrentMethodName())
}

/**
 *  是否收藏过，多个用逗号分隔
 */
func (c UserCollectController) CheckIsCollect(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCollectUri+tool.CurrentMethodName())
}

/**
 *  获取收藏数量
 */
func (c UserCollectController) GetCollectNums(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCollectUri+tool.CurrentMethodName())
}

/**
 *  获取用户收藏列表
 */
func (c UserCollectController) GetUserCollectionList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCollectUri+tool.CurrentMethodName())
}

/**
 * 删除收藏
 */
func (c UserCollectController) DeleteCollect(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCollectUri+tool.CurrentMethodName())
}
