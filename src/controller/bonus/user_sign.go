package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserSignController struct {
	abs.Controller
}

const userSignUri = "http://bonus.hxsapp.com/base/UserSign/"

func UserSign() UserSignController {
	c := UserSignController{}
	c.Config = map[string]abs.MethodConfig{
		"SetUserSign":     {ControllerNameFirstUpper: true},
		"GetSignState":    {ControllerNameFirstUpper: true},
		"GetSignDaysData": {ControllerNameFirstUpper: true},
		"UserSignInfo":    {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 签到
 * @Author   lifeng
 * @DateTime 2017-09-16T11:35:54+0800
 * @return   [type]                   [description]
 */
func (c UserSignController) SetUserSign(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userSignUri+tool.CurrentMethodName())
}

/**
 * 获取签到状态
 * @Author   lifeng
 * @DateTime 2017-09-16T11:42:53+0800
 * @return   [type]                   [description]
 */
func (c UserSignController) GetSignState(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userSignUri+tool.CurrentMethodName())
}

/**
 * 获取用户签到天数内容
 * @Author   lifeng
 * @DateTime 2017-09-16T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserSignController) GetSignDaysData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userSignUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2019-02-20 11:03:18
 * @Description: 用户签到天数卡片 3.8.7
 */
func (c UserSignController) UserSignInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userSignUri+tool.CurrentMethodName())
}
