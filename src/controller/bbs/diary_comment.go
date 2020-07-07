package bbs

import (
	"src/abs"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

// DiaryCommentController 评论模块
type DiaryCommentController struct {
	abs.Controller
}

const diaryCommentUri = "http://bbs.hxsapp.com/diary/comment/"

// DiaryComment 构造函数
func DiaryComment() DiaryCommentController {
	c := DiaryCommentController{}
	c.Config = map[string]abs.MethodConfig{
		"Add":      {Rest: "diary/comment/add"},
		"List":     {Rest: "diary/comment/list"},
		"Template": {Rest: "diary/comment/template"},
	}
	return c
}

// Add 添加评论
func (c DiaryCommentController) Add(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryCommentUri+tool.CurrentMethodName())
}

// List 评论列表
func (c DiaryCommentController) List(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryCommentUri+tool.CurrentMethodName())
}

// Template 评论模板
func (c DiaryCommentController) Template(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryCommentUri+tool.CurrentMethodName())
}
