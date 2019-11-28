package account

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type AdviserController struct {
	abs.Controller
}

const adviserUri = "http://account.hxsapp.com/user/Adviser/"

func Adviser() AdviserController {
	c := AdviserController{}
	c.Config = map[string]abs.MethodConfig{
		"GetAdviserList":       {ControllerNameFirstUpper: true},
		"GetMyAdviserInfo":     {ControllerNameFirstUpper: true},
		"GetAdviserCard":       {ControllerNameFirstUpper: true},
		"GetCurAdviserRank":    {ControllerNameFirstUpper: true},
		"GetConsultRewardInfo": {ControllerNameFirstUpper: true},
		"GetConsultReward":     {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 获取顾问列表
 */
func (c AdviserController) GetAdviserList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adviserUri+tool.CurrentMethodName())
}

func (c AdviserController) GetMyAdviserInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adviserUri+tool.CurrentMethodName())
}

func (c AdviserController) GetAdviserCard(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adviserUri+tool.CurrentMethodName())
}

func (c AdviserController) GetCurAdviserRank(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adviserUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2019-06-25 17:16:22
 * @Description: 顾问的咨询奖励
 */
func (c AdviserController) GetConsultRewardInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adviserUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2019-06-25 17:36:42
 * @Description: 顾问的咨询奖励 - 第一、二、三次奖励
 */
func (c AdviserController) GetConsultReward(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adviserUri+tool.CurrentMethodName())
}
