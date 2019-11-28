package mall

import (
    "src/abs"
    "src/library/tool"
    "github.com/gin-gonic/gin"
)

type DailyRecController struct {
	abs.Controller
}

const dailyRecUri = "http://mall.hxsapp.com/base/DailyRec/"

func DailyRec() DailyRecController {
	c := DailyRecController{}
    c.Config = map[string]abs.MethodConfig{
        "Goods": {ControllerNameFirstUpper: true},
    }
	return c
}

func (c DailyRecController) Goods(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dailyRecUri+tool.CurrentMethodName())
}
