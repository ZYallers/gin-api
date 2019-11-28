package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type PrepaymentController struct {
	abs.Controller
}

const prepaymentUri = "http://mall.hxsapp.com/base/Prepayment/"

func Prepayment() PrepaymentController {
	c := PrepaymentController{}
	c.Config = map[string]abs.MethodConfig{
		"AddPay":           {ControllerNameFirstUpper: true},
		"GetOrderDetail":   {ControllerNameFirstUpper: true},
		"GetOrderList":     {ControllerNameFirstUpper: true},
		"OrderQuery":       {ControllerNameFirstUpper: true},
		"GetOrderListByIm": {ControllerNameFirstUpper: true},
		"OrderInfo":        {ControllerNameFirstUpper: true},
		"GetOrderCount":    {ControllerNameFirstUpper: true},
		"GetOrderOnline":   {ControllerNameFirstUpper: true},
		"PrepaymentToBrm":  {ControllerNameFirstUpper: true},
	}
	return c
}

func (c PrepaymentController) AddPay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, prepaymentUri+tool.CurrentMethodName())
}

func (c PrepaymentController) GetOrderDetail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, prepaymentUri+tool.CurrentMethodName())
}

func (c PrepaymentController) GetOrderList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, prepaymentUri+tool.CurrentMethodName())
}

func (c PrepaymentController) OrderQuery(ctx *gin.Context) {
	c.ServiceRewrite(ctx, prepaymentUri+tool.CurrentMethodName())
}

func (c PrepaymentController) GetOrderListByIm(ctx *gin.Context) {
	c.ServiceRewrite(ctx, prepaymentUri+tool.CurrentMethodName())
}

func (c PrepaymentController) OrderInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, prepaymentUri+tool.CurrentMethodName())
}

func (c PrepaymentController) GetOrderCount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, prepaymentUri+tool.CurrentMethodName())
}

func (c PrepaymentController) GetOrderOnline(ctx *gin.Context) {
	c.ServiceRewrite(ctx, prepaymentUri+tool.CurrentMethodName())
}

func (c PrepaymentController) PrepaymentToBrm(ctx *gin.Context) {
	c.ServiceRewrite(ctx, prepaymentUri+tool.CurrentMethodName())
}
