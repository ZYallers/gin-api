package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type PlatformWhiteController struct {
	abs.Controller
}

const platformWhiteUri = "http://community.hxsapp.com/common/platformWhite/"

func PlatformWhite() PlatformWhiteController {
	c := PlatformWhiteController{}
	c.Config = map[string]abs.MethodConfig{
		"IsUserWhite": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 是否白名单用户
 */
func (c PlatformWhiteController) IsUserWhite(ctx *gin.Context) {
	c.ServiceRewrite(ctx, platformWhiteUri+tool.CurrentMethodName())
}
