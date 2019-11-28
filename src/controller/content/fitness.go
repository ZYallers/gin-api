package content

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type FitnessController struct {
	abs.Controller
}

const fitnessUri = "http://content.hxsapp.com/course/Fitness/"

func Fitness() FitnessController {
	c := FitnessController{}
	c.Config = map[string]abs.MethodConfig{
		"GetSchemeList":           {ControllerNameFirstUpper: true},
		"GetCourseListBySchemeId": {ControllerNameFirstUpper: true},
		"GetDetailByCourseId":     {ControllerNameFirstUpper: true},
		"GetActionByCourseId":     {ControllerNameFirstUpper: true},
	}
	return c
}

func (c FitnessController) GetSchemeList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, fitnessUri+tool.CurrentMethodName())
}

func (c FitnessController) GetCourseListBySchemeId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, fitnessUri+tool.CurrentMethodName())
}

func (c FitnessController) GetDetailByCourseId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, fitnessUri+tool.CurrentMethodName())
}

func (c FitnessController) GetActionByCourseId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, fitnessUri+tool.CurrentMethodName())
}
