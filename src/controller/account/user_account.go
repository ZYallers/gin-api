package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"src/abs"
	"src/library/tool"
)

type UserAccountController struct {
	abs.Controller
}

const userAccountUri = "http://account.hxsapp.com/user/UserAccount/"
const imUri = "http://im.hxsapp.com/api/Tim/"

func UserAccount() UserAccountController {
	c := UserAccountController{}
	c.Config = map[string]abs.MethodConfig{
		"GetTimSig":                  {ControllerNameFirstUpper: true},
		"ErrorHandlerTIm":            {ControllerNameFirstUpper: true},
		"GetUserInfoByWeChatUnionId": {ControllerNameFirstUpper: true},
		"CallbackTIm":                {ControllerNameFirstUpper: true, HttpMethods: []string{http.MethodPut, http.MethodPost, http.MethodGet}},
		"RegisterServerTim":          {ControllerNameFirstUpper: true},
		"ModifyMobile":               {ControllerNameFirstUpper: true},
	}
	return c
}

func (c UserAccountController) RegisterServerTim(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c UserAccountController) GetTimSig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c UserAccountController) RegisterTIm(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c UserAccountController) UserTimAccount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c UserAccountController) CheckUserTim(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c UserAccountController) CallbackTIm(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c UserAccountController) ErrorHandlerTIm(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c UserAccountController) MakeCallByPhone(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

/**
 * 登录
 * @Author   lifeng
 * @DateTime 2017-09-15T17:58:22+0800
 * @return   [type]                   [description]
 */
func (c UserAccountController) Login(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

/**
 * 第三方登录
 * @Author   lifeng
 * @DateTime 2017-09-15T18:03:12+0800
 * @return   [type]                   [description]
 */
func (c UserAccountController) ThirdPartyLogin(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) GetUserInfoByWeChatUnionId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) OutsideLogin(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

/**
 * 注销登录
 * @Author   lifeng
 * @DateTime 2017-09-15T18:02:46+0800
 * @return   [type]                   [description]
 */
func (c UserAccountController) Logout(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

/**
 * 登录
 * @Author   lifeng
 * @DateTime 2017-09-15T17:58:35+0800
 * @return   [type]                   [description]
 */
func (c UserAccountController) Register(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

/**
 * 发短信
 * login - 登录; register - 注册; bindMobile - 绑定手机; modifyMobile - 修改手机
 * @Author   lifeng
 * @DateTime 2017-09-15T17:59:21+0800
 * @return   [type]                   [description]
 */
func (c UserAccountController) SendSMS(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

/**
 * 获取im账号
 * @Author   lifeng
 * @DateTime 2017-09-16T10:50:20+0800
 * @return   [type]                   [description]
 */
func (c UserAccountController) GetOpenIMAccount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

/**
 * 修改手机号
 * @Author   lifeng
 * @DateTime 2017-09-16T10:50:52+0800
 * @return   [type]                   [description]
 */
func (c UserAccountController) ModifyMobile(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) ResetPwd(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) HasPwd(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) ThirdPartyBindInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) ThirdPartyBindOrUnbind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) VerifySmsCode(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) RelateMobile(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}

func (c UserAccountController) BrmUserIdChange(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userAccountUri+tool.CurrentMethodName())
}
