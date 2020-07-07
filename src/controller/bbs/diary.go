package bbs

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type DiaryController struct {
	abs.Controller
}

const diaryUri = "http://bbs.hxsapp.com/diary/"

func Diary() DiaryController {
	return DiaryController{}
}

func (c DiaryController) Add(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryUri+tool.CurrentMethodName())
}

func (c DiaryController) Examine(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryUri+tool.CurrentMethodName())
}

func (c DiaryController) List(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryUri+tool.CurrentMethodName())
}

func (c DiaryController) Detail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryUri+tool.CurrentMethodName())
}

func (c DiaryController) Delete(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryUri+tool.CurrentMethodName())
}

func (c DiaryController) Check(ctx *gin.Context) {
	c.ServiceRewrite(ctx, diaryUri+tool.CurrentMethodName())
}
