package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type MedalController struct {
	abs.Controller
}

const medalUri = "http://bonus.hxsapp.com/adviser/Medal/"

func Medal() MedalController {
	c := MedalController{}
	c.Config = map[string]abs.MethodConfig{
		"GetMedal":     {ControllerNameFirstUpper: true},
		"GetMedalWall": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * @Author:      caius
 * @DateTime:    2018-05-25 10:11:59
 * @Description: 获取顾问勋章
 */
func (c MedalController) GetMedal(ctx *gin.Context) {
	c.ServiceRewrite(ctx, medalUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-05-25 10:11:59
 * @Description: 获取顾问勋章墙
 */
func (c MedalController) GetMedalWall(ctx *gin.Context) {
	c.ServiceRewrite(ctx, medalUri+tool.CurrentMethodName())
}
