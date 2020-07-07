package account

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type EquityController struct {
	abs.Controller
}

const groupStaffEquityUri = "http://account.hxsapp.com/user/Equity/"

func Equity() EquityController {
	c := EquityController{}
	c.Config = map[string]abs.MethodConfig{
		"StaffEquity":     {ControllerNameFirstUpper: true},
		"Advertise":       {ControllerNameFirstUpper: true},
		"Popup":           {ControllerNameFirstUpper: true},
		"UnUseEquity":     {ControllerNameFirstUpper: true},
		"DrawEquity":      {ControllerNameFirstUpper: true},
		"GradeEquity":     {ControllerNameFirstUpper: true},
		"DrawWelfare":     {ControllerNameFirstUpper: true},
		"StaffGoods":      {ControllerNameFirstUpper: true},
	}
	return c
}

// StaffEquity 获取员工福利列表
func (c EquityController) StaffEquity(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupStaffEquityUri+tool.CurrentMethodName())
}

// Advertise 获取员工福利入口图
func (c EquityController) Advertise(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupStaffEquityUri+tool.CurrentMethodName())
}

// Popup 获取权益弹窗
func (c EquityController) Popup(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupStaffEquityUri+tool.CurrentMethodName())
}

// UnUseEquity 获取已领取未使用的权益
func (c EquityController) UnUseEquity(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupStaffEquityUri+tool.CurrentMethodName())
}

// DrawEquity 领取权益
func (c EquityController) DrawEquity(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupStaffEquityUri+tool.CurrentMethodName())
}

// GradeEquity 等级权益
func (c EquityController) GradeEquity(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupStaffEquityUri+tool.CurrentMethodName())
}

// DrawWelfare 领取福利
func (c EquityController) DrawWelfare(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupStaffEquityUri+tool.CurrentMethodName())
}

// StaffGoods 领取福利
func (c EquityController) StaffGoods(ctx *gin.Context) {
	c.ServiceRewrite(ctx, groupStaffEquityUri+tool.CurrentMethodName())
}