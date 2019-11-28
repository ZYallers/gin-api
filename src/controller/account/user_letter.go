package account

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserLetterController struct {
	abs.Controller
}

const userLetterUri = "http://account.hxsapp.com/user/UserLetter/"

func UserLetter() UserLetterController {
	c := UserLetterController{}
	c.Config = map[string]abs.MethodConfig{
		"Letter": {ControllerNameFirstUpper: true},
		"Agree":  {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * @Author:      caius
 * @DateTime:    2019-02-22 15:14:44
 * @Description: 查看信封
 */
func (c UserLetterController) Letter(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userLetterUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2019-02-22 15:55:31
 * @Description: 信协议
 */
func (c UserLetterController) Agree(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userLetterUri+tool.CurrentMethodName())
}
