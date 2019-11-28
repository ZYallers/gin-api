package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type LogisticController struct {
	abs.Controller
}

const logisticUri = "http://mall.hxsapp.com/base/Logistic/"

func Logistic() LogisticController {
	c := LogisticController{}
	c.Config = map[string]abs.MethodConfig{
		"GetExpress": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c LogisticController) GetExpress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, logisticUri+tool.CurrentMethodName())
}
