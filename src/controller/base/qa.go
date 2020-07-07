package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type QAController struct {
	abs.Controller
}

const qaUri = "http://base.hxsapp.com/base/QA/"

func QA() QAController {
	c := QAController{}
	return c
}


func (c QAController) GetQAList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, qaUri+tool.CurrentMethodName())
}

func (c QAController) SetUsefulState(ctx *gin.Context) {
	c.ServiceRewrite(ctx, qaUri+tool.CurrentMethodName())
}

func (c QAController) GetRuleConfig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, qaUri+tool.CurrentMethodName())
}
