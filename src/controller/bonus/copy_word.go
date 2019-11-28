package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CopyWordController struct {
	abs.Controller
}

const copyWordUri = "http://bonus.hxsapp.com/base/CopyWord/"

func CopyWord() CopyWordController {
	c := CopyWordController{}
	c.Config = map[string]abs.MethodConfig{
		"ParseWord": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c CopyWordController) ParseWord(ctx *gin.Context) {
	c.ServiceRewrite(ctx, copyWordUri+tool.CurrentMethodName())
}
