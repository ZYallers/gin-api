package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CommentController struct {
	abs.Controller
}

const commentUri = "http://slim.hxsapp.com/base/Comment/"

func Comment() CommentController {
	c := CommentController{}
	c.Config = map[string]abs.MethodConfig{
		"Impress": {ControllerNameFirstUpper: true},
		"Voucher": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 顾问口碑 基本信息
 */
func (c CommentController) GetInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commentUri+tool.CurrentMethodName())
}

/**
 * 顾问口碑 印象标签
 */
func (c CommentController) Impress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commentUri+tool.CurrentMethodName())
}

/**
 * 顾问口碑 口碑列表
 */
func (c CommentController) GetList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commentUri+tool.CurrentMethodName())
}

/**
 * 顾问口碑 新增口碑
 */
func (c CommentController) Voucher(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commentUri+tool.CurrentMethodName())
}
