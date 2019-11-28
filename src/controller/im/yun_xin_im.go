package im

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type YunXinIMController struct {
	abs.Controller
}

const yunXinIMUri = "http://im.hxsapp.com/im/YunXinIM/"

func YunXinIM() YunXinIMController {
	c := YunXinIMController{}
	c.Config = map[string]abs.MethodConfig{
		"Token":      {ControllerNameFirstUpper: true},
		"ReceiveMsg": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c YunXinIMController) Token(ctx *gin.Context) {
	c.ServiceRewrite(ctx, yunXinIMUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2019-01-24 14:12:50
 * @Description: 网易云抄送 3.8.5
 */
func (c YunXinIMController) ReceiveMsg(ctx *gin.Context) {
	c.ServiceRewrite(ctx, yunXinIMUri+tool.CurrentMethodName())
}
