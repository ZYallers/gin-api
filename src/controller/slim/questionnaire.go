package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type QuestionnaireController struct {
	abs.Controller
}

const questionnaireUri = "http://slim.hxsapp.com/other/Questionnaire/"

func Questionnaire() QuestionnaireController {
	c := QuestionnaireController{}
	return c
}

func (c QuestionnaireController) GetAllQuestion(ctx *gin.Context) {
	c.ServiceRewrite(ctx, questionnaireUri+tool.CurrentMethodName())
}

func (c QuestionnaireController) SaveUserAnswer(ctx *gin.Context) {
	c.ServiceRewrite(ctx, questionnaireUri+tool.CurrentMethodName())
}

func (c QuestionnaireController) GetUserLabel(ctx *gin.Context) {
	c.ServiceRewrite(ctx, questionnaireUri+tool.CurrentMethodName())
}

func (c QuestionnaireController) HasUserFinishQuestionnaire(ctx *gin.Context) {
	c.ServiceRewrite(ctx, questionnaireUri+tool.CurrentMethodName())
}
