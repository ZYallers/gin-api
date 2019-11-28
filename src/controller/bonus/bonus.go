package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type BonusController struct {
	abs.Controller
}

const bonusUri = "http://bonus.hxsapp.com/bonus/Bonus/"

func Bonus() BonusController {
	c := BonusController{}
	c.Config = map[string]abs.MethodConfig{
		"ShowPopupBonus":         {ControllerNameFirstUpper: true},
		"OpenPopupBonus":         {ControllerNameFirstUpper: true},
		"ShowActivityPopupBonus": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * [showPopupBonus 显示弹窗红包]
 * @author liangguangzeng 2018-01-25
 * @return [type] [description]
 */
func (c BonusController) ShowPopupBonus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bonusUri+tool.CurrentMethodName())
}

/**
 * [openPopupBonus 打开弹屏红包]
 * @author liangguangzeng 2018-01-25
 * @return [type] [description]
 */
func (c BonusController) OpenPopupBonus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bonusUri+tool.CurrentMethodName())
}

/**
 * bonus域调用act域特例
 * [showActivityPopupBonus 判断是否打开活动广告弹屏]
 * @author liangguangzeng 2018-06-27
 * @return [type] [description]
 */
func (c BonusController) ShowActivityPopupBonus(ctx *gin.Context) {

	c.ServiceRewrite(ctx,"http://act.hxsapp.com/bonus/MilkBonus/getActBonusReceive")
}
