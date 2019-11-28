package api

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type TimController struct {
	abs.Controller
}

const timUri = "http://im.hxsapp.com/api/Tim/"

func Tim() TimController {
	c := TimController{}
	c.Config = map[string]abs.MethodConfig{
		"RecordLoginLog":         {ControllerNameFirstUpper: true},
		"GetConsultList":         {ControllerNameFirstUpper: true},
		"CallSaleSystemFeedback": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c TimController) PaymentSystemNotice(ctx *gin.Context) {
	c.ServiceRewrite(ctx, timUri+tool.CurrentMethodName())
}

func (c TimController) AfterSaleSystemNotice(ctx *gin.Context) {
	c.ServiceRewrite(ctx, timUri+tool.CurrentMethodName())
}

func (c TimController) CallSaleSystemFeedback(ctx *gin.Context) {
	c.ServiceRewrite(ctx, timUri+tool.CurrentMethodName())
}

func (c TimController) PlaceOrderNotice(ctx *gin.Context) {
	c.ServiceRewrite(ctx, timUri+tool.CurrentMethodName())
}

func (c TimController) RecordLoginLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, timUri+tool.CurrentMethodName())
}

func (c TimController) GetConsultList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, timUri+tool.CurrentMethodName())
}

func (c TimController) FixedUser(ctx *gin.Context) {
	c.ServiceRewrite(ctx, timUri+tool.CurrentMethodName())
}

func (c TimController) BatchReply(ctx *gin.Context) {
	c.ServiceRewrite(ctx, timUri+tool.CurrentMethodName())
}

func (c TimController) UpdateChatInfoByAccountAndImg(ctx *gin.Context) {
	c.ServiceRewrite(ctx, timUri+tool.CurrentMethodName())
}
