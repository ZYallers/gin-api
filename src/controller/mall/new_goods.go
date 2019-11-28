package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type NewGoodsController struct {
	abs.Controller
}

const newGoodsUri = "http://mall.hxsapp.com/base/NewGoods/"

func NewGoods() NewGoodsController {
	c := NewGoodsController{}
	c.Config = map[string]abs.MethodConfig{
		"NhomeGoods":    {ControllerNameFirstUpper: true, MethodNameFirstUpper: true},
		"HomeFloor":     {ControllerNameFirstUpper: true},
		"NewGoodsDet":   {ControllerNameFirstUpper: true},
		"GroupGoodsDet": {ControllerNameFirstUpper: true},
		"DetectUserLevel": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c NewGoodsController) NhomeGoods(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}

func (c NewGoodsController) HomeFloor(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}

func (c NewGoodsController) NewGoodsDet(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}

func (c NewGoodsController) GroupGoodsDet(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}
func (c NewGoodsController) DetectUserLevel(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}