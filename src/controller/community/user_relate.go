package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserRelateController struct {
	abs.Controller
}

const userRelateUri = "http://community.hxsapp.com/user/userRelate/"

func UserRelate() UserRelateController {
	c := UserRelateController{}
	return c
}

/**
 * 关注
 */
func (c UserRelateController) Follow(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userRelateUri+tool.CurrentMethodName())
}

/**
 * 检测是否关注过
 */
func (c UserRelateController) CheckIsFollow(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userRelateUri+tool.CurrentMethodName())
}

/**
 * 取消关注
 */
func (c UserRelateController) CancelFollow(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userRelateUri+tool.CurrentMethodName())
}

/**
 * 移除粉丝
 */
func (c UserRelateController) RemoveFan(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userRelateUri+tool.CurrentMethodName())
}

/**
 * 获取关注列表
 */
func (c UserRelateController) GetFollowsList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userRelateUri+tool.CurrentMethodName())
}

/**
 * 获取粉丝列表
 */
func (c UserRelateController) GetFansList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userRelateUri+tool.CurrentMethodName())
}

/**
 * 获取用户粉丝数
 */
func (c UserRelateController) GetZFansCount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userRelateUri+tool.CurrentMethodName())
}

/**
 * 获取用户关注数
 */
func (c UserRelateController) GetZFollowsCount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userRelateUri+tool.CurrentMethodName())
}

/**
 * 总粉丝排行榜
 */
func (c UserRelateController) GetFansRanking(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userRelateUri+tool.CurrentMethodName())
}
