package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"src/abs"
	"src/library/tool"
	"strings"
)

type UserInfoController struct {
	abs.Controller
}

const userInfoUri = "http://account.hxsapp.com/user/UserInfo/"
const imfoUri = "http://im.hxsapp.com/api/Brm/"

func UserInfo() UserInfoController {
	c := UserInfoController{}
	c.Config = map[string]abs.MethodConfig{
		"SaveInfo":                     {HttpMethods: []string{http.MethodPost}},
		"GetTalentFixed":               {ControllerNameFirstUpper: true},
		"GetTalentRandomTop":           {ControllerNameFirstUpper: true},
		"GetTalentFansRankList":        {ControllerNameFirstUpper: true},
		"GetChatInfoByChatAccount":     {ControllerNameFirstUpper: true},
		"CanImWithChatAccount":         {ControllerNameFirstUpper: true},
		"GetUserStatusByOpenImAccount": {ControllerNameFirstUpper: true},
		"GetSelfUserInfo2":             {HttpMethods: []string{http.MethodGet}},
	}
	return c
}

/**
 * 获取当前登录用户信息
 * @Author   lifeng
 * @DateTime 2017-09-16T10:53:29+0800
 * @return   [type]                   [description]
 */
func (c UserInfoController) GetSelfUserInfo(ctx *gin.Context) {
	//c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
	c.ServiceRewrite(ctx, userInfoUri+"getSelfUserInfo2")
}

func (c UserInfoController) GetSelfUserInfo2(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 通过用户id获取用户信息，多个用逗号分隔  如  60250,60251
 * 传 from_type = admin 可拿到手机号
 * @Author   lifeng
 * @DateTime 2017-09-16T10:54:20+0800
 * @return   [type]                   [description]
 */
func (c UserInfoController) GetUserInfo(ctx *gin.Context) {
	//c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
	c.ServiceMultiRewrite(ctx, userInfoUri+tool.CurrentMethodName(), "user_id", 10)
}

/**
 * 保存、修改个人资料
 * @Author   lifeng
 * @DateTime 2017-09-16T10:58:45+0800
 * @return   [type]                   [description]
 */
func (c UserInfoController) SaveInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 新增家庭用户
 * @Author   lifeng
 * @DateTime 2017-09-16T10:59:25+0800
 */
func (c UserInfoController) AddFamily(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 删除家庭用户
 * @Author   lifeng
 * @DateTime 2017-09-16T10:59:25+0800
 */
func (c UserInfoController) DeleteFamily(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 获取自己邀请码，包括邀请总数和邀请列表
 * @Author   lifeng
 * @DateTime 2017-09-16T11:01:00+0800
 * @return   [type]                   [description]
 */
func (c UserInfoController) GetInviteCode(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 分页获取邀请列表
 * @Author   lifeng
 * @DateTime 2017-09-16T11:02:20+0800
 * @return   [type]                   [description]
 */
func (c UserInfoController) GetInviteUserList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 获取地址列表
 * @Author   lifeng
 * @DateTime 2017-09-16T11:16:14+0800
 * @return   [type]                   [description]
 */
func (c UserInfoController) GetAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 保存地址
 * @Author   lifeng
 * @DateTime 2017-09-16T11:16:44+0800
 * @return   [type]                   [description]
 */
func (c UserInfoController) SaveAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 删除地址
 * @Author   lifeng
 * @DateTime 2017-09-16T11:16:44+0800
 * @return   [type]                   [description]
 */
func (c UserInfoController) DeleteAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 设置默认地址
 * @Author   lifeng
 * @DateTime 2017-09-16T11:16:44+0800
 * @return   [type]                   [description]
 */
func (c UserInfoController) SetAddressDefault(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 实名认证
 */
func (c UserInfoController) Realname(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

/**
 * 获取当前实名信息
 */
func (c UserInfoController) GetRealname(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) MutiGetUserInfoByMobile(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) GetTalentFixed(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) GetTalentRandomTop(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) GetTalentRandom(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) GetTalentFansRankList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) GetUserIdByOpenImAccount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) GetUserInfoByOpenImAccount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imfoUri+tool.CurrentMethodName())
	//c.ServiceMultiRewrite(ctx, imfoUri+tool.CurrentMethodName(), "openim_account", 20)
}

func (c UserInfoController) CanImWithChatAccount(ctx *gin.Context) {
	save := ctx.Request.Body
	save, ctx.Request.Body, _ = tool.DrainBody(ctx.Request.Body)
	key := "chat_account"
	keyValue := tool.GetQueryPostForm(ctx, key)
	ctx.Request.Body = save
	if keyValue, _ = url.QueryUnescape(keyValue); len(strings.Split(keyValue, ",")) > 50 {
		c.ServiceMultiRewrite(ctx, userInfoUri+tool.CurrentMethodName(), key, 10)
	} else {
		c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
	}
}

func (c UserInfoController) GetUserInfoByChatAccount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) GetUserStatusByOpenImAccount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) GetChatInfoByChatAccount(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) ChooseMobileToSummary(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}

func (c UserInfoController) GetGroupChatAuth(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userInfoUri+tool.CurrentMethodName())
}
