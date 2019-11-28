package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserAccountController struct {
	abs.Controller
}

const userAccountUri = "http://mall.hxsapp.com/base/UserAccount/"

func UserAccount() UserAccountController {
	c := UserAccountController{}
	c.Config = map[string]abs.MethodConfig{
		"GetUserAccount": {ControllerNameFirstUpper: true},
		"GetAccountInfo": {ControllerNameFirstUpper: true},
		"GetAccountLog":  {ControllerNameFirstUpper: true},
		"GetBalanceLog":  {ControllerNameFirstUpper: true},
	}
	return c
}

func (c UserAccountController) GetUserAccount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) GetAccountInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) GetAccountLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) GetBalanceLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}
