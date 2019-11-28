package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ManualController struct {
	abs.Controller
}

const manualUri = "http://slim.hxsapp.com/other/Manual/"

func Manual() ManualController {
	c := ManualController{}
	return c
}

/**
 * 保存围度类型
 */
func (c ManualController) SaveManualTypes(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manualUri+tool.CurrentMethodName())
}

/**
 * 获取围度类型
 */
func (c ManualController) GetManualTypes(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manualUri+tool.CurrentMethodName())
}

func (c ManualController) GetManualData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manualUri+tool.CurrentMethodName())
}

func (c ManualController) SaveManualData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manualUri+tool.CurrentMethodName())
}

func (c ManualController) SaveManualDataToClock(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manualUri+tool.CurrentMethodName())
}

func (c ManualController) GetManualDataByType(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manualUri+tool.CurrentMethodName())
}

func (c ManualController) GetBloodData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manualUri+tool.CurrentMethodName())
}

/**
 * 保存血压血脂血糖数据
 */
func (c ManualController) SaveBloodData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, manualUri+tool.CurrentMethodName())
}
