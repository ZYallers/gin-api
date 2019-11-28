package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type LsuserController struct {
	abs.Controller
}

const lsuserUri = "http://mall.hxsapp.com/lsmember/Lsuser/"

func Lsuser() LsuserController {
	c := LsuserController{}
	c.Config = map[string]abs.MethodConfig{
		"UserMember": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c LsuserController) UserMember(ctx *gin.Context) {
	c.ServiceRewrite(ctx, lsuserUri+tool.CurrentMethodName())
}
