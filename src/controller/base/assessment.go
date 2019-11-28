package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type AssessmentController struct {
	abs.Controller
}

const assessmentUri = "http://base.hxsapp.com/base/Assessment/"

func Assessment() AssessmentController {
	c := AssessmentController{}
	c.Config = map[string]abs.MethodConfig{
		"SaveAssessment": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c AssessmentController) SaveAssessment(ctx *gin.Context) {
	c.ServiceRewrite(ctx, assessmentUri+tool.CurrentMethodName())
}
