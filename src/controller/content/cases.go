package content

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CasesController struct {
	abs.Controller
}

const casesUri = "http://content.hxsapp.com/case/Cases/"

func Cases() CasesController {
	c := CasesController{}
	c.Config = map[string]abs.MethodConfig{
		"HomeList":      {ControllerNameFirstUpper: true},
		"Labellist":     {ControllerNameFirstUpper: true},
		"Detail":        {ControllerNameFirstUpper: true},
		"RecommendCase": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c CasesController) HomeList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, casesUri+tool.CurrentMethodName())
}

func (c CasesController) Labellist(ctx *gin.Context) {
	c.ServiceRewrite(ctx, casesUri+tool.CurrentMethodName())
}

func (c CasesController) Detail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, casesUri+tool.CurrentMethodName())
}

func (c CasesController) RecommendCase(ctx *gin.Context) {
	c.ServiceRewrite(ctx, casesUri+tool.CurrentMethodName())
}
