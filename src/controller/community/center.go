package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CenterController struct {
	abs.Controller
}

const centerUri = "http://community.hxsapp.com/common/Center/"

func Center() CenterController {
	c := CenterController{}
	c.Config = map[string]abs.MethodConfig{
		"Banner": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * @Author:      caius
 * @DateTime:    2019-03-21 15:31:08
 * @Description: banner
 */
func (c CenterController) Banner(ctx *gin.Context) {
	c.ServiceRewrite(ctx, centerUri+tool.CurrentMethodName())
}
