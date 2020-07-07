package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type NewsTipController struct {
	abs.Controller
}

const newsTipUri = "http://bonus.hxsapp.com/base/NewsTip/"

func NewsTip() NewsTipController {
	c := NewsTipController{}
	c.Config = map[string]abs.MethodConfig{
	}
	return c
}

func (c NewsTipController) GetTips(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newsTipUri+tool.CurrentMethodName())
}
