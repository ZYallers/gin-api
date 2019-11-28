package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type GoodsController struct {
	abs.Controller
}

const goodsUri = "http://mall.hxsapp.com/base/Goods/"

func Goods() GoodsController {
	c := GoodsController{}
	c.Config = map[string]abs.MethodConfig{
		"GetGoodsDetails":            {ControllerNameFirstUpper: true},
		"GetGoodsList":               {ControllerNameFirstUpper: true},
		"GetGoodsListWithProperties": {ControllerNameFirstUpper: true},
		"PopularRedemption":          {ControllerNameFirstUpper: true},
	}
	return c
}

func (c GoodsController) GetGoodsDetails(ctx *gin.Context) {
	c.ServiceRewrite(ctx, goodsUri+tool.CurrentMethodName())
}

func (c GoodsController) GetGoodsList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, goodsUri+tool.CurrentMethodName())
}

func (c GoodsController) GetGoodsListWithProperties(ctx *gin.Context) {
	c.ServiceRewrite(ctx, goodsUri+tool.CurrentMethodName())
}

func (c GoodsController) PopularRedemption(ctx *gin.Context) {
	c.ServiceRewrite(ctx, goodsUri+tool.CurrentMethodName())
}
