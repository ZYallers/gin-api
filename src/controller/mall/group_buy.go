package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type GroupBuyController struct {
	abs.Controller
}

const groupBuyUri = "http://mall.hxsapp.com/base/GroupBuy/"

func GroupBuy() GroupBuyController {
	c := GroupBuyController{}
	c.Config = map[string]abs.MethodConfig{
		"GetGroupBuyItemList":       {ControllerNameFirstUpper: true},
		"GetGroupBuyList":           {ControllerNameFirstUpper: true},
		"GetGroupBuyInfo":           {ControllerNameFirstUpper: true},
		"GetGuessFavoriteGoodsItem": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c GroupBuyController) GetGroupBuyItemList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBuyUri+tool.CurrentMethodName())
}

func (c GroupBuyController) GetGroupBuyList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBuyUri+tool.CurrentMethodName())
}

func (c GroupBuyController) GetGroupBuyInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBuyUri+tool.CurrentMethodName())
}

func (c GroupBuyController) GetGuessFavoriteGoodsItem(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBuyUri+tool.CurrentMethodName())
}
