package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserAccusationController struct {
	abs.Controller
}

const userAccusationUri = "http://community.hxsapp.com/user/userAccusation/"

func UserAccusation() UserAccusationController {
	c := UserAccusationController{}
	return c
}

/**
 * 保存举报
 */
func (c UserAccusationController) SaveAccusation(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccusationUri+tool.CurrentMethodName())
}
