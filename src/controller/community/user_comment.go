package community

import (
	"src/abs"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

type UserCommentController struct {
	abs.Controller
}

const userCommentUri = "http://community.hxsapp.com/user/userComment/"

func UserComment() UserCommentController {
	c := UserCommentController{}
	c.Config = map[string]abs.MethodConfig{
		"DeleteComment":          {ControllerNameFirstUpper: true},
		"GetCommentNums":         {ControllerNameFirstUpper: true},
		"GetCommentChildrenList": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 用户评论
 */
func (c UserCommentController) DoComment(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommentUri+tool.CurrentMethodName())
}

/**
 * 用户删除评论
 */
func (c UserCommentController) DeleteComment(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommentUri+tool.CurrentMethodName())
}

/**
 * 根据评论id取评论信息
 */
func (c UserCommentController) GetCommentById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommentUri+tool.CurrentMethodName())
}

/**
 * 取评论列表
 */
func (c UserCommentController) GetCommentList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommentUri+tool.CurrentMethodName())
}

/**
 * 获取评论数量
 */
func (c UserCommentController) GetCommentNums(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommentUri+tool.CurrentMethodName())
}

/**
 * 取评论的子评论列表
 */
func (c UserCommentController) GetCommentChildrenList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userCommentUri+tool.CurrentMethodName())
}
