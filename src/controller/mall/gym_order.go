package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type GymOrderController struct {
	abs.Controller
}

const gymOrderUri = "http://mall.hxsapp.com/gym/Order/"

func GymOrder() GymOrderController {
	c := GymOrderController{}
	c.Config = map[string]abs.MethodConfig{
		"AddPay":        {ControllerNameFirstUpper: true},
		"OrderQuery":    {ControllerNameFirstUpper: true},
		"GetAccountLog": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c GymOrderController) AddPay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymOrderUri+tool.CurrentMethodName())
}

func (c GymOrderController) OrderQuery(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymOrderUri+tool.CurrentMethodName())
}

func (c GymOrderController) GetAccountLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymOrderUri+tool.CurrentMethodName())
}
