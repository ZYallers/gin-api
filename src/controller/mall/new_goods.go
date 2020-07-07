package mall

import (
	"src/abs"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

type NewGoodsController struct {
	abs.Controller
}

const newGoodsUri = "http://mall.hxsapp.com/base/NewGoods/"

func NewGoods() NewGoodsController {
	c := NewGoodsController{}
	c.Config = map[string]abs.MethodConfig{
		"NhomeGoods":        {ControllerNameFirstUpper: true, MethodNameFirstUpper: true},
		"HomeFloor":         {ControllerNameFirstUpper: true},
		"NewGoodsDet":       {ControllerNameFirstUpper: true},
		"GroupGoodsDet":     {ControllerNameFirstUpper: true},
		"DetectUserLevel":   {ControllerNameFirstUpper: true},
		"GetGoodsCoupon":    {ControllerNameFirstUpper: true},
		"WebGetGoodsCoupon": {ControllerNameFirstUpper: true},
		"ShareGoodsDet":       {ControllerNameFirstUpper: true},
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

func (c NewGoodsController) GetGoodsCoupon(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}

func (c NewGoodsController) WebGetGoodsCoupon(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}

func (c NewGoodsController) GetWatchingGoods(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}

func (c NewGoodsController) GetRecommendGoods(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}

func (c NewGoodsController) IsShowRecommend(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}


func (c NewGoodsController) ShareGoodsDet(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newGoodsUri+tool.CurrentMethodName())
}
