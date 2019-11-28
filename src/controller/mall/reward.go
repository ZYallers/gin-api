package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type RewardController struct {
	abs.Controller
}

const rewardUri = "http://mall.hxsapp.com/base/Reward/"

func Reward() RewardController {
	c := RewardController{}
	c.Config = map[string]abs.MethodConfig{
		"AddReward":    {ControllerNameFirstUpper: true},
		"OrderQuery":   {ControllerNameFirstUpper: true},
		"GetRewardLog": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c RewardController) AddReward(ctx *gin.Context) {
	c.ServiceRewrite(ctx, rewardUri+tool.CurrentMethodName())
}

func (c RewardController) OrderQuery(ctx *gin.Context) {
	c.ServiceRewrite(ctx, rewardUri+tool.CurrentMethodName())
}

func (c RewardController) GetRewardLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, rewardUri+tool.CurrentMethodName())
}
