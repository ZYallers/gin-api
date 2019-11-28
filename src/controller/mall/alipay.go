package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type AlipayController struct {
	abs.Controller
}

const alipayUri = "http://mall.hxsapp.com/mall/Alipay/"

func Alipay() AlipayController {
	c := AlipayController{}
	c.Config = map[string]abs.MethodConfig{
		"AlipayTradeAppPayNotify": {ControllerNameFirstUpper: true},
		"OrderNotify":             {ControllerNameFirstUpper: true},
		"RewardNotify":            {ControllerNameFirstUpper: true},
		"LvOrderNotify":           {ControllerNameFirstUpper: true},
		"KashgarOrderNotify":      {ControllerNameFirstUpper: true},
		"PercentOrderNotify":      {ControllerNameFirstUpper: true},
		"PepopleOrderNotify":        {ControllerNameFirstUpper: true},
	}
	return c
}

func (c AlipayController) AlipayTradeAppPayNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, alipayUri+tool.CurrentMethodName())
}
func (c AlipayController) KashgarOrderNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, alipayUri+tool.CurrentMethodName())
}

func (c AlipayController) LvOrderNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, alipayUri+tool.CurrentMethodName())
}

func (c AlipayController) OrderNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, alipayUri+tool.CurrentMethodName())
}

func (c AlipayController) RewardNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, alipayUri+tool.CurrentMethodName())
}

func (c AlipayController) PercentOrderNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, alipayUri+tool.CurrentMethodName())
}

func (c AlipayController) PepopleOrderNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, alipayUri+tool.CurrentMethodName())
}
