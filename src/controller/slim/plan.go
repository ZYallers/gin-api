package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type PlanController struct {
	abs.Controller
}

const planUri = "http://slim.hxsapp.com/slim/plan/"

func Plan() PlanController {
	c := PlanController{}
	return c
}

/**
 * 设置目标
 */
func (c PlanController) SaveTarget(ctx *gin.Context) {
	c.ServiceRewrite(ctx, planUri+tool.CurrentMethodName())
}

/**
 * 获取目标
 */
func (c PlanController) GetTarget(ctx *gin.Context) {
	c.ServiceRewrite(ctx, planUri+tool.CurrentMethodName())
}

/**
 * 获取我的最后目标
 */
func (c PlanController) GetMyEndTarget(ctx *gin.Context) {
	c.ServiceRewrite(ctx, planUri+tool.CurrentMethodName())
}

/**
 *  获取日历
 */
func (c PlanController) GetCalendar(ctx *gin.Context) {
	c.ServiceRewrite(ctx, planUri+tool.CurrentMethodName())
}

/**
 * 我的计划
 */
func (c PlanController) GetMyPlan(ctx *gin.Context) {
	c.ServiceRewrite(ctx, planUri+tool.CurrentMethodName())
}
