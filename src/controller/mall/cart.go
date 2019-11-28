package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CartController struct {
	abs.Controller
}

const cartUri = "http://mall.hxsapp.com/base/Cart/"

func Cart() CartController {
	c := CartController{}
	c.Config = map[string]abs.MethodConfig{
		"GetCartNum":     {ControllerNameFirstUpper: true},
		"CartList":       {ControllerNameFirstUpper: true},
		"AddCart":        {ControllerNameFirstUpper: true},
		"DelCart":        {ControllerNameFirstUpper: true},
		"GetNewCartList": {ControllerNameFirstUpper: true},
		"GetNewCartNum":  {ControllerNameFirstUpper: true},
		"DelNewCart":     {ControllerNameFirstUpper: true},
		"UpdateCart":     {ControllerNameFirstUpper: true},
		"AddNewCart":     {ControllerNameFirstUpper: true},
		"CheckStock":     {ControllerNameFirstUpper: true},
	}
	return c
}

func (c CartController) GetCartNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, cartUri+tool.CurrentMethodName())
}

func (c CartController) CartList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, cartUri+tool.CurrentMethodName())
}

func (c CartController) AddCart(ctx *gin.Context) {
	c.ServiceRewrite(ctx, cartUri+tool.CurrentMethodName())
}

func (c CartController) DelCart(ctx *gin.Context) {
	c.ServiceRewrite(ctx, cartUri+tool.CurrentMethodName())
}

func (c CartController) GetNewCartList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, cartUri+tool.CurrentMethodName())
}

func (c CartController) GetNewCartNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, cartUri+tool.CurrentMethodName())
}

func (c CartController) DelNewCart(ctx *gin.Context) {
	c.ServiceRewrite(ctx, cartUri+tool.CurrentMethodName())
}

func (c CartController) UpdateCart(ctx *gin.Context) {
	c.ServiceRewrite(ctx, cartUri+tool.CurrentMethodName())
}

func (c CartController) AddNewCart(ctx *gin.Context) {
	c.ServiceRewrite(ctx, cartUri+tool.CurrentMethodName())
}

func (c CartController) CheckStock(ctx *gin.Context) {
	c.ServiceRewrite(ctx, cartUri+tool.CurrentMethodName())
}
