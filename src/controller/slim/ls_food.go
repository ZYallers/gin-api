package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type LsFoodController struct {
	abs.Controller
}

const lsFoodUri = "http://slim.hxsapp.com/lsmember/LsFood/"

func LsFood() LsFoodController {
	c := LsFoodController{}
	return c
}

/**
 * 饮食计划
 */
func (c LsFoodController) GetPlan(ctx *gin.Context) {
	c.ServiceRewrite(ctx, lsFoodUri+tool.CurrentMethodName())
}

/**
 * 饮食计划缓一缓
 */
func (c LsFoodController) ChangPlan(ctx *gin.Context) {
	c.ServiceRewrite(ctx, lsFoodUri+tool.CurrentMethodName())
}

/**
 * 饮食计划详情
 */
func (c LsFoodController) GetPlanDetail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, lsFoodUri+tool.CurrentMethodName())
}
