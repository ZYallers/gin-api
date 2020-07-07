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
		"AddReportFeedback":   {ControllerNameFirstUpper: true},
		"GetReportFeedback":   {ControllerNameFirstUpper: true},
		"SaveHxsCareService":   {ControllerNameFirstUpper: true},
		"SetReminderListNotice":   {ControllerNameFirstUpper: true},
		"SetReminderList":   {ControllerNameFirstUpper: true},
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


func (c OnlineServiceController) GetClockReportByPlanId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetAllClockPlanIds(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) AddReportFeedback(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetReportFeedback(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) SaveHxsCareService(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) SetReminderListNotice(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) SetReminderList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

// 获取该吃什么模板数据
func (c OnlineServiceController) GetFood(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

// 录入吃了什么
func (c OnlineServiceController) SetFood(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetXYServicePlanDates(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}


func (c OnlineServiceController) OnlineServiceState(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) OnFinishSchemeForXY(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetBodyInfoReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetBodyClockDataByDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetActProgress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetSportRecord(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetServiceConfirmStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) SetServiceConfirm(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}

func (c OnlineServiceController) GetServiceOrderList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, onlineServiceUri+tool.CurrentMethodName())
}