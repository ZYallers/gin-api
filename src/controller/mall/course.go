package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CourseController struct {
	abs.Controller
}

const courseUri = "http://mall.hxsapp.com/base/course/"

func Course() CourseController {
	c := CourseController{}
	return c
}

func (c CourseController) GetClist(ctx *gin.Context) {
	c.ServiceRewrite(ctx, courseUri+tool.CurrentMethodName())
}