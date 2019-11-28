package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type EvaluateController struct {
	abs.Controller
}

const evaluateUri = "http://slim.hxsapp.com/tool/evaluate/"

func Evaluate() EvaluateController {
	c := EvaluateController{}
	return c
}

/**
 * 瘦身-集合小工具
 */
func (c EvaluateController) GetBodyReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, evaluateUri+tool.CurrentMethodName())
}

/**
 * 瘦身小工具 - 减脂运动心率
 */
func (c EvaluateController) SaveSportHeartRate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, evaluateUri+tool.CurrentMethodName())
}

/**
 * 瘦身小工具 - 标准体重测量
 */
func (c EvaluateController) SaveWeightMeasure(ctx *gin.Context) {
	c.ServiceRewrite(ctx, evaluateUri+tool.CurrentMethodName())
}

/**
 * 瘦身小工具 - 标准身材比例
 */
func (c EvaluateController) SaveStatureProportion(ctx *gin.Context) {
	c.ServiceRewrite(ctx, evaluateUri+tool.CurrentMethodName())
}

/**
 * 瘦身小工具 - 每天最高摄入量
 */
func (c EvaluateController) SaveHighestIntake(ctx *gin.Context) {
	c.ServiceRewrite(ctx, evaluateUri+tool.CurrentMethodName())
}

/**
 * 瘦身小工具 - 评测保存
 */
func (c EvaluateController) SaveReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, evaluateUri+tool.CurrentMethodName())
}

/**
 * 瘦身小工具 - 评测ID查看报告
 */
func (c EvaluateController) GetReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, evaluateUri+tool.CurrentMethodName())
}
