package account

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserMemberController struct {
	abs.Controller
}

const userMemberUri = "http://account.hxsapp.com/user/UserMember/"

func UserMember() UserMemberController {
	c := UserMemberController{}
	c.Config = map[string]abs.MethodConfig{
		"GetConsumeFromBrm":   {ControllerNameFirstUpper: true},
		"GetMsgByUserId":      {ControllerNameFirstUpper: true},
		"AddMsgByUserId":      {ControllerNameFirstUpper: true},
		"ReadMsgById":         {ControllerNameFirstUpper: true},
		"GetUnreadNum":        {ControllerNameFirstUpper: true},
		"ReadAllMsg":          {ControllerNameFirstUpper: true},
		"GetVipCardInfo":      {ControllerNameFirstUpper: true},
		"GetEquityAllPreview": {ControllerNameFirstUpper: true},
		"getEquityByGradeId":  {ControllerNameFirstUpper: true},
		"CustomerService":     {ControllerNameFirstUpper: true},
		"GetEquityPopup":      {ControllerNameFirstUpper: true},
		"GetUnUseEquity":      {ControllerNameFirstUpper: true},
		"GetAllGradeEquity":   {ControllerNameFirstUpper: true},
		"DrawEquityByIds":     {ControllerNameFirstUpper: true},
	}
	return c
}

func (c UserMemberController) GetConsumeFromBrm(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

func (c UserMemberController) GetMsgByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

func (c UserMemberController) AddMsgByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

func (c UserMemberController) ReadMsgById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

func (c UserMemberController) GetUnreadNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

func (c UserMemberController) ReadAllMsg(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

func (c UserMemberController) GetVipCardInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

func (c UserMemberController) GetEquityAllPreview(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

// 获取指定会员等级的权益
func (c UserMemberController) getEquityByGradeId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

func (c UserMemberController) CustomerService(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

// 获取权益弹窗
func (c UserMemberController) GetEquityPopup(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

// 获取可报名的权益
func (c UserMemberController) GetUnUseEquity(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

// 获取所有会员等级权益及领取状况
func (c UserMemberController) GetAllGradeEquity(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

// 领取权益
func (c UserMemberController) DrawEquityByIds(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}

// 判断用户是不是 一二线 新客户 450
func (c UserMemberController) UserTag(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMemberUri+tool.CurrentMethodName())
}
