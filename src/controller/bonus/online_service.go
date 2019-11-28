package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type OnlineServiceController struct {
	abs.Controller
}

const onlineServiceUri = "http://bonus.hxsapp.com/base/OnlineService/"

func OnlineService() OnlineServiceController {
	c := OnlineServiceController{}
	c.Config = map[string]abs.MethodConfig{
		"IsWhiteListUser":     {ControllerNameFirstUpper: true},
		"SetQuestionByUserId": {ControllerNameFirstUpper: true},
		"GetQuestionByUserId": {ControllerNameFirstUpper: true},
		"HasQuestionReport":   {ControllerNameFirstUpper: true},
		"GetQuestionReport":   {ControllerNameFirstUpper: true},
	}
	return c
}

func (c OnlineServiceController) IsWhiteListUser(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) SetQuestionByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetQuestionByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) HasQuestionReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetQuestionReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

