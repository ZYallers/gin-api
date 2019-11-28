package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type MarketController struct {
	abs.Controller
}

const marketUri = "http://mall.hxsapp.com/base/MarketingArea/"

func Market() MarketController {
	c := MarketController{}
	c.Config = map[string]abs.MethodConfig{
		"Marketing": {ControllerNameFirstUpper: true, MethodNameFirstUpper: true},
	}
	return c
}

func (c MarketController) Marketing(ctx *gin.Context) {
	c.ServiceRewrite(ctx, marketUri+tool.CurrentMethodName())
}
