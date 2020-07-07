package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type NewOnlineServiceController struct {
	abs.Controller
}

const newonlineServiceUri = "http://bonus.hxsapp.com/base/NewOnlineService/"

func NewOnlineService() NewOnlineServiceController {
	c := NewOnlineServiceController{}
	c.Config = map[string]abs.MethodConfig{
	}
	return c
}

func (c NewOnlineServiceController) SetQuestionByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newonlineServiceUri+tool.CurrentMethodName())
}

func (c NewOnlineServiceController) GetQuestionByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newonlineServiceUri+tool.CurrentMethodName())
}

func (c NewOnlineServiceController) GetQuestionReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newonlineServiceUri+tool.CurrentMethodName())
}

func (c NewOnlineServiceController) OnlineServiceState(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newonlineServiceUri+tool.CurrentMethodName())
}

func (c NewOnlineServiceController) GetSchedule(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newonlineServiceUri+tool.CurrentMethodName())
}

func (c NewOnlineServiceController) SetDietPhoto(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newonlineServiceUri+tool.CurrentMethodName())
}

func (c NewOnlineServiceController) SetClockCare(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newonlineServiceUri+tool.CurrentMethodName())
}

func (c NewOnlineServiceController) GetSummaryGoods(ctx *gin.Context) {
	c.ServiceRewrite(ctx, newonlineServiceUri+tool.CurrentMethodName())
}

