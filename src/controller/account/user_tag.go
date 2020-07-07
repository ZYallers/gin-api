package account

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserTagController struct {
	abs.Controller
}

const userTagUri = "http://account.hxsapp.com/user/userTag/"

func Usertag() UserTagController {
	c := UserTagController{}
	return c
}

func (c UserTagController) GetList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userTagUri+tool.CurrentMethodName())
}

func (c UserTagController) Save(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userTagUri+tool.CurrentMethodName())
}

func (c UserTagController) GetCity(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userTagUri+tool.CurrentMethodName())
}