package bbs

import (
	"src/abs"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

// DiaryLikeController 点赞模块
type DiaryLikeController struct {
	abs.Controller
}

const diaryLikeUri = "http://bbs.hxsapp.com/diary/like/"

// DiaryLike 构造函数
func DiaryLike() DiaryLikeController {
	c := DiaryLikeController{}
	c.Config = map[string]abs.MethodConfig{
		"Save": {Rest: "diary/like/save"},
		"List": {Rest: "diary/like/list"},
	}
	return c
}

// Save 点赞 / 取消点赞
func (c DiaryLikeController) Save(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryLikeUri+tool.CurrentMethodName())
}

// List 指定动态的点赞列表
func (c DiaryLikeController) List(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryLikeUri+tool.CurrentMethodName())
}
