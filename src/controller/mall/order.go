package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type OrderController struct {
	abs.Controller
}

const orderUri = "http://mall.hxsapp.com/base/Order/"

func Order() OrderController {
	c := OrderController{}
	c.Config = map[string]abs.MethodConfig{
		"AddOrder":           {ControllerNameFirstUpper: true},
		"ConfirmOrder":       {ControllerNameFirstUpper: true},
		"CheckStock":         {ControllerNameFirstUpper: true},
		"OrderQuery":         {ControllerNameFirstUpper: true},
		"RePay":              {ControllerNameFirstUpper: true},
		"DelOrder":           {ControllerNameFirstUpper: true},
		"CancelOrder":        {ControllerNameFirstUpper: true},
		"BuildOrder":         {ControllerNameFirstUpper: true},
		"GetOrderList":       {ControllerNameFirstUpper: true},
		"GetExpress":         {ControllerNameFirstUpper: true},
		"GetOrderDetailList": {ControllerNameFirstUpper: true},
		"ConfirmNewOrder":    {ControllerNameFirstUpper: true},
		"AddNewOrder":        {ControllerNameFirstUpper: true},
		"AddOrderLog":        {ControllerNameFirstUpper: true},
		"WebAddNewOrder":     {ControllerNameFirstUpper: true},
		"WebConfirmNewOrder": {ControllerNameFirstUpper: true},
		"WebRePay":           {ControllerNameFirstUpper: true},
		"CancelGroupOrder":   {ControllerNameFirstUpper: true},
		"AlipayNotify":       {ControllerNameFirstUpper: true},
		"WechatNotify":       {ControllerNameFirstUpper: true},
		"WebWxpayNotify":     {ControllerNameFirstUpper: true},
		"GetVirtualCode":     {ControllerNameFirstUpper: true},
	}
	return c
}

func (c OrderController) AddOrder(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) ConfirmOrder(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) CheckStock(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) OrderQuery(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) RePay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) DelOrder(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) CancelOrder(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) BuildOrder(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) GetOrderList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) GetExpress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) GetOrderDetailList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) ConfirmNewOrder(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) AddNewOrder(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) AddOrderLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) WebAddNewOrder(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) WebConfirmNewOrder(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) WebRePay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) CancelGroupOrder(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}

func (c OrderController) AlipayNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}
func (c OrderController) WechatNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}
func (c OrderController) WebWxpayNotify(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}
func (c OrderController) GetVirtualCode(ctx *gin.Context) {
	c.ServiceRewrite(ctx, orderUri+tool.CurrentMethodName())
}
