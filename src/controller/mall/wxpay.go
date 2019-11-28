package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type WxpayController struct {
	abs.Controller
}

const wxpayUri = "http://mall.hxsapp.com/mall/Wxpay/"

func Wxpay() WxpayController {
	c := WxpayController{}
	c.Config = map[string]abs.MethodConfig{
		"KashgarOrderNotify": {ControllerNameFirstUpper: true},
		"LvOrderNotify":      {ControllerNameFirstUpper: true},
		"OrderNotify":        {ControllerNameFirstUpper: true},
		"RewardNotify":       {ControllerNameFirstUpper: true},
		"PercentOrderNotify": {ControllerNameFirstUpper: true},
		"PeoplOrderNotify":   {ControllerNameFirstUpper: true},
	}
	return c
}

func (c WxpayController) KashgarOrderNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, wxpayUri+tool.CurrentMethodName())
}

func (c WxpayController) LvOrderNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, wxpayUri+tool.CurrentMethodName())
}

func (c WxpayController) OrderNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, wxpayUri+tool.CurrentMethodName())
}

func (c WxpayController) RewardNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, wxpayUri+tool.CurrentMethodName())
}

func (c WxpayController) PercentOrderNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, wxpayUri+tool.CurrentMethodName())
}

func (c WxpayController) PeoplOrderNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, wxpayUri+tool.CurrentMethodName())
}
