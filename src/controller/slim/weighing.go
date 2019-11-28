package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type WeighingController struct {
	abs.Controller
}

const weighingUri = "http://hardware.hxsapp.com/device/Weighing/"

func Weighing() WeighingController {
	c := WeighingController{}
	return c
}

/**
 * 保存身体报告
 * @Author   lifeng
 * @DateTime 2017-09-16T14:06:01+0800
 */
func (c WeighingController) SaveReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

func (c WeighingController) SaveReportToClock(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

/**
 * 删除报告
 * @Author   lifeng
 * @DateTime 2017-09-16T14:09:00+0800
 * @return   [type]                   [description]
 */
func (c WeighingController) DeleteReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

/**
 * 获取身体报告列表
 * @Author   lifeng
 * @DateTime 2017-09-16T14:09:09+0800
 * @return   [type]                   [description]
 */
func (c WeighingController) GetReportList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

/**
 * 通过id获取身体报告
 * @Author   lifeng
 * @DateTime 2017-09-16T14:09:24+0800
 * @return   [type]                   [description]
 */
func (c WeighingController) GetReportByIdOrDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

/**
 * 获取最后一次报告
 * @Author   lifeng
 * @DateTime 2017-09-16T14:09:38+0800
 * @return   [type]                   [description]
 */
func (c WeighingController) GetUserLastReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

/**
 * 获取最近有报告的七天（每天最后一次）身体报告
 * @Author   lifeng
 * @DateTime 2017-09-16T14:09:49+0800
 * @return   [type]                   [description]
 */
func (c WeighingController) GetUserWeightLastXDays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

func (c WeighingController) GetReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

func (c WeighingController) GetInfoByType(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

func (c WeighingController) GetRankTen(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

func (c WeighingController) GetUserFirstAndLastWeight(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

func (c WeighingController) GetUserWeightDataByTypeLastXDays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

func (c WeighingController) CheckAdc(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}

func (c WeighingController) ErrorLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weighingUri+tool.CurrentMethodName())
}
