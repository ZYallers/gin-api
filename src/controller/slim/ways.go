package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type WaysController struct {
	abs.Controller
}

const waysUri = "http://slim.hxsapp.com/slim/ways/"

func Ways() WaysController {
	c := WaysController{}
	return c
}

/**
 * 获取我的计划
 */
func (c WaysController) GetMyWays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, waysUri+tool.CurrentMethodName())
}

/**
 * 删除我的计划
 */
func (c WaysController) DeleteMyWays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, waysUri+tool.CurrentMethodName())
}

/**
 * 更新我的计划
 */
func (c WaysController) UpdateMyWays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, waysUri+tool.CurrentMethodName())
}

/**
 *  保存 饮食、运动记录
 */
func (c WaysController) SaveFoodSportLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, waysUri+tool.CurrentMethodName())
}

/**
 * 获取我的减肥报告
 */
func (c WaysController) GetReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, waysUri+tool.CurrentMethodName())
}

/**
 * 获取方法
 */
func (c WaysController) GetWay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, waysUri+tool.CurrentMethodName())
}

/**
 * 获取方法详情
 */
func (c WaysController) GetWayDetails(ctx *gin.Context) {
	c.ServiceRewrite(ctx, waysUri+tool.CurrentMethodName())
}

/**
 * 获取方法天数详情
 */
func (c WaysController) GetWayDay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, waysUri+tool.CurrentMethodName())
}
