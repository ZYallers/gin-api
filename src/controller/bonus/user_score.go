package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserScoreController struct {
	abs.Controller
}

const userScoreUri = "http://bonus.hxsapp.com/base/UserScore/"

func UserScore() UserScoreController {
	c := UserScoreController{}
	c.Config = map[string]abs.MethodConfig{
		"GetScoreAndHonor": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 获取积分和勋章
 * @Author   lifeng
 * @DateTime 2017-09-16T11:35:54+0800
 * @return   [type]                   [description]
 */
func (c UserScoreController) GetScoreAndHonor(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userScoreUri+tool.CurrentMethodName())
}
