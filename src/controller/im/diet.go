package im

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type DietController struct {
	abs.Controller
}

const dietUri = "http://im.hxsapp.com/im/Diet/"

func Diet() DietController {
	c := DietController{}
	c.Config = map[string]abs.MethodConfig{
		"KeywordSearch":          {ControllerNameFirstUpper: true},
		"GetPunchConfigData":     {ControllerNameFirstUpper: true},
		"ApplyUserDietTemplate":  {ControllerNameFirstUpper: true},
		"DeleteUserDietTemplate": {ControllerNameFirstUpper: true},
		"SaveTemplateConfig":     {ControllerNameFirstUpper: true},
		"GetTemplateDetail":      {ControllerNameFirstUpper: true},
		"EditTemplateDetail":     {ControllerNameFirstUpper: true},
		"GetPunchSituation":      {ControllerNameFirstUpper: true},
		"GetUserDietTemplate":    {ControllerNameFirstUpper: true},
	}
	return c
}

func (c DietController) KeywordSearch(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dietUri+tool.CurrentMethodName())
}

func (c DietController) GetPunchConfigData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dietUri+tool.CurrentMethodName())
}

func (c DietController) ApplyUserDietTemplate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dietUri+tool.CurrentMethodName())
}

func (c DietController) DeleteUserDietTemplate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dietUri+tool.CurrentMethodName())
}

func (c DietController) SaveTemplateConfig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dietUri+tool.CurrentMethodName())
}

func (c DietController) GetTemplateDetail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dietUri+tool.CurrentMethodName())
}

func (c DietController) EditTemplateDetail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dietUri+tool.CurrentMethodName())
}

func (c DietController) GetPunchSituation(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dietUri+tool.CurrentMethodName())
}

func (c DietController) GetUserDietTemplate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dietUri+tool.CurrentMethodName())
}
