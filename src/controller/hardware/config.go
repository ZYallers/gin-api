package hardware

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ConfigController struct {
	abs.Controller
}

const configUri = "http://hardware.hxsapp.com/device/config/"

func Config() ConfigController {
	c := ConfigController{}
	c.Config = map[string]abs.MethodConfig{
		"SaveUserConfig":   {ControllerNameFirstUpper: true},
		"GetUserConfig":    {ControllerNameFirstUpper: true},
		"CheckShoesIsBind": {ControllerNameFirstUpper: true},
		"GetVideo":         {ControllerNameFirstUpper: true},
		"GetHardwareToast": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 智能硬件用户配置
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112543
 * @author chs
 */
func (c ConfigController) SaveUserConfig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, configUri+tool.CurrentMethodName())
}

/**
 * 智能硬件配置读取
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112538
 * @author chs
 */
func (c ConfigController) GetUserConfig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, configUri+tool.CurrentMethodName())
}

/**
 * 检查MAC是否被绑定
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112554
 * @author chs
 */
func (c ConfigController) CheckShoesIsBind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, configUri+tool.CurrentMethodName())
}

func (c ConfigController) GetVideo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, configUri+tool.CurrentMethodName())
}

func (c ConfigController) GetHardwareToast(ctx *gin.Context) {
	c.ServiceRewrite(ctx, configUri+tool.CurrentMethodName())
}
