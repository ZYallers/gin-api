package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type GymUniformPayController struct {
	abs.Controller
}

const gymUniformPayUri = "http://mall.hxsapp.com/gym/UniformPay/"

func GymUniformPay() GymUniformPayController {
	c := GymUniformPayController{}
	c.Config = map[string]abs.MethodConfig{
		"Pay":          {ControllerNameFirstUpper: true},
		"OrderQuery":   {ControllerNameFirstUpper: true},
		"GymQuery":     {ControllerNameFirstUpper: true},
		"Gympay":       {ControllerNameFirstUpper: true},
		"GympayQuery":  {ControllerNameFirstUpper: true},
		"GymRefund": {ControllerNameFirstUpper: true},
		"AgymNotify":   {ControllerNameFirstUpper: true},
		"WgymNotify":   {ControllerNameFirstUpper: true},
		"AlipayNotify": {ControllerNameFirstUpper: true},
		"WechatNotify": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c GymUniformPayController) Pay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymUniformPayUri+tool.CurrentMethodName())
}

func (c GymUniformPayController) OrderQuery(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymUniformPayUri+tool.CurrentMethodName())
}

func (c GymUniformPayController) GymQuery(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymUniformPayUri+tool.CurrentMethodName())
}

func (c GymUniformPayController) Gympay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymUniformPayUri+tool.CurrentMethodName())
}

func (c GymUniformPayController) GympayQuery(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymUniformPayUri+tool.CurrentMethodName())
}

func (c GymUniformPayController) GymRefund(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymUniformPayUri+tool.CurrentMethodName())
}

func (c GymUniformPayController) AgymNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymUniformPayUri+tool.CurrentMethodName())
}

func (c GymUniformPayController) WgymNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymUniformPayUri+tool.CurrentMethodName())
}

func (c GymUniformPayController) AlipayNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymUniformPayUri+tool.CurrentMethodName())
}

func (c GymUniformPayController) WechatNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymUniformPayUri+tool.CurrentMethodName())
}
