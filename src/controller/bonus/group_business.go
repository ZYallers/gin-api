package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type GroupBusinessController struct {
	abs.Controller
}

const groupBusinessUri = "http://bonus.hxsapp.com/base/GroupBusiness/"

func GroupBusiness() GroupBusinessController {
	c := GroupBusinessController{}
	c.Config = map[string]abs.MethodConfig{
		"GetPhysicalStoreList":         {ControllerNameFirstUpper: true},
		"GetSetting":                   {ControllerNameFirstUpper: true},
		"GetArticleList":               {ControllerNameFirstUpper: true},
		"GetAuthority":                 {ControllerNameFirstUpper: true},
		"SendMessage":                  {ControllerNameFirstUpper: true},
		"GetMessage":                   {ControllerNameFirstUpper: true},
		"SendGreetings":                {ControllerNameFirstUpper: true},
		"Store":                        {ControllerNameFirstUpper: true},
	}
	return c
}

func (c GroupBusinessController) GetPhysicalStoreList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBusinessUri+tool.CurrentMethodName())
}

func (c GroupBusinessController) GetSetting(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBusinessUri+tool.CurrentMethodName())
}

func (c GroupBusinessController) GetArticleList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBusinessUri+tool.CurrentMethodName())
}

func (c GroupBusinessController) GetAuthority(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBusinessUri+tool.CurrentMethodName())
}

func (c GroupBusinessController) SendMessage(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBusinessUri+tool.CurrentMethodName())
}

func (c GroupBusinessController) GetMessage(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBusinessUri+tool.CurrentMethodName())
}

func (c GroupBusinessController) SendGreetings(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBusinessUri+tool.CurrentMethodName())
}

func (c GroupBusinessController) Store(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupBusinessUri+tool.CurrentMethodName())
}
