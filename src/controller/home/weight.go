package home

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type WeightController struct {
	abs.Controller
}

const weightUri = "http://hardware.hxsapp.com/home/Weight/"

func Weight() WeightController {
	c := WeightController{}
	c.Config = map[string]abs.MethodConfig{
		"GetReduceWeightData": {ControllerNameFirstUpper: true},
		"EvaluateBmi":         {ControllerNameFirstUpper: true},
		"GetSevenWeightRank":  {ControllerNameFirstUpper: true},
		"GetScaleDangerDescr": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 查看今日减肥人数、平均体重、平均减重数据
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17124389
 * @author chs
 */
func (c WeightController) GetReduceWeightData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weightUri+tool.CurrentMethodName())
}

/**
 * 身体评测结果
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17124553
 * @author chs
 */
func (c WeightController) EvaluateBmi(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weightUri+tool.CurrentMethodName())
}

/**
 * 7天减重排行榜
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17128302
 * @author chs
 */
func (c WeightController) GetSevenWeightRank(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weightUri+tool.CurrentMethodName())
}

/**
 * 体脂数据风险项小知识
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17129405
 * @author chs
 */
func (c WeightController) GetScaleDangerDescr(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weightUri+tool.CurrentMethodName())
}
