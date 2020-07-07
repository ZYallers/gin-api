package mall

import (
	"src/abs"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

type CouponController struct {
	abs.Controller
}

const couponUri = "http://mall.hxsapp.com/base/Coupon/"

func Coupon() CouponController {
	c := CouponController{}
	c.Config = map[string]abs.MethodConfig{
		"CollectCoupons":      {ControllerNameFirstUpper: true},
		"CouponNum":           {ControllerNameFirstUpper: true},
		"CouponStatistics":    {ControllerNameFirstUpper: true},
		"GetCouponsList":      {ControllerNameFirstUpper: true},
		"GetCouponsListById":  {ControllerNameFirstUpper: true},
		"List":                {ControllerNameFirstUpper: true},
		"WebCollectCoupons":   {ControllerNameFirstUpper: true},
		"WebCouponStatistics": {ControllerNameFirstUpper: true},
		"Weblist":             {ControllerNameFirstUpper: true},
		"GetCouponGoods":      {ControllerNameFirstUpper: true},
	}
	return c
}

func (c CouponController) CollectCoupons(ctx *gin.Context) {
	c.ServiceRewrite(ctx, couponUri+tool.CurrentMethodName())
}

func (c CouponController) CouponNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, couponUri+tool.CurrentMethodName())
}

func (c CouponController) CouponStatistics(ctx *gin.Context) {
	c.ServiceRewrite(ctx, couponUri+tool.CurrentMethodName())
}

func (c CouponController) GetCouponsList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, couponUri+tool.CurrentMethodName())
}

func (c CouponController) GetCouponsListById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, couponUri+tool.CurrentMethodName())
}

func (c CouponController) List(ctx *gin.Context) {
	c.ServiceRewrite(ctx, couponUri+tool.CurrentMethodName())
}

func (c CouponController) WebCollectCoupons(ctx *gin.Context) {
	c.ServiceRewrite(ctx, couponUri+tool.CurrentMethodName())
}

func (c CouponController) WebCouponStatistics(ctx *gin.Context) {
	c.ServiceRewrite(ctx, couponUri+tool.CurrentMethodName())
}

func (c CouponController) Weblist(ctx *gin.Context) {
	c.ServiceRewrite(ctx, couponUri+tool.CurrentMethodName())
}

func (c CouponController) GetCouponGoods(ctx *gin.Context) {
	c.ServiceRewrite(ctx, couponUri+tool.CurrentMethodName())
}
