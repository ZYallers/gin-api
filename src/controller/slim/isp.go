package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type IspController struct {
	abs.Controller
}

const ispUri = "http://slim.hxsapp.com/slim/Isp/"

func Isp() IspController {
	c := IspController{}
	c.Config = map[string]abs.MethodConfig{
		"GetWeightReportList": {ControllerNameFirstUpper: true},
		"GetMultiWeightReportList": {ControllerNameFirstUpper: true},
		"GetGeneReportList": {ControllerNameFirstUpper: true},
		"GetManualData": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 根据手机获取身体报告列表（手动、体脂称、有氧机）
 */
func (c IspController) GetWeightReportList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, ispUri+tool.CurrentMethodName())
}

/**
 * 根据批量手机获取身体报告列表
 */
func (c IspController) GetMultiWeightReportList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, ispUri+tool.CurrentMethodName())
}

/**
 * 获取基因检测报告（FTO、3+1、易健康、伊美颜）
 */
func (c IspController) GetGeneReportList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, ispUri+tool.CurrentMethodName())
}

/**
 * 获取手动记录（包括腰围、腿围等）
 */
func (c IspController) GetManualData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, ispUri+tool.CurrentMethodName())
}
