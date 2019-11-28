package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type Market400Controller struct {
	abs.Controller
}

const market400Uri = "http://mall.hxsapp.com/base/MarketingArea400/"

func Market400() Market400Controller {
	c := Market400Controller{}
	c.Config = map[string]abs.MethodConfig{
		"Marketing": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c Market400Controller) Marketing(ctx *gin.Context) {
	c.ServiceRewrite(ctx, market400Uri+tool.CurrentMethodName())
}
