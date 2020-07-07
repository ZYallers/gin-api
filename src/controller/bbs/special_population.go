package bbs

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type SpecialPopulationController struct {
	abs.Controller
}

const specialPopulationUri = "http://bbs.hxsapp.com/population/"

func SpecialPopulation() SpecialPopulationController {
	return SpecialPopulationController{}
}

func (c SpecialPopulationController) Detail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, specialPopulationUri+tool.CurrentMethodName())
}
