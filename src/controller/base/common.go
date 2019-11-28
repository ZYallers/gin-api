package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CommonController struct {
	abs.Controller
}

const commonUri = "http://base.hxsapp.com/base/common/"

func Common() CommonController {
	c := CommonController{}
	c.Config = map[string]abs.MethodConfig{
		"ConsultConfig":   {ControllerNameFirstUpper: true},
		"ConsultClose":    {ControllerNameFirstUpper: true},
		"IsConsultClosed": {ControllerNameFirstUpper: true},
		"GetPopUpState":   {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * [getAdListByType 取指定类型的广告列表]
 * @return [type] [description]
 */
func (c CommonController) GetAdListByType(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * [getSystemWords 获取系统文字说明配置]
 * @return [type] [description]
 */
func (c CommonController) GetSystemWords(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * [getClientVersion 获取客户端最新版本号信息]
 * @return [type] [description]
 */
func (c CommonController) GetClientVersion(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * [getCommonQuestions 获取系统常见问题列表]
 * @return [type] [description]
 */
func (c CommonController) GetCommonQuestions(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * [getWxShareJsApiSignature 获取微信分享样式定制jssdk需要的校验参数]
 * @return [type] [description]
 */
func (c CommonController) GetWxShareJsApiSignature(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * 获取微信公众号通用接口调用需要的access_token
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=12133816
 * @author ZYaller
 */
func (c CommonController) GetWxAccessToken(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * [shareCounts 分享计数]
 * @return [type] [description]
 */
func (c CommonController) ShareCounts(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * [saveFeedback 保存用户意见反馈]
 * @return [type] [description]
 */
func (c CommonController) SaveFeedback(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * [getUserFeedback 获取用户意见反馈列表]
 * @return [type] [description]
 */
func (c CommonController) GetUserFeedback(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * [isUserHaveUnreads 用户是否有各种未读的消息]
 * @return boolean [description]
 */
func (c CommonController) IsUserHaveUnreads(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * [doReadFeedback 把用户未读意见反馈置为已读]
 * @return [type] [description]
 */
func (c CommonController) DoReadFeedback(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * [getChannelConfig 获取渠道配置信息]
 * @return [type] [description]
 */
func (c CommonController) GetChannelConfig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetAllRecommend(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetUserShareLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * 配合 活动
 * @Author   lifeng
 * @DateTime 2017-07-25T11:09:54+0800
 * @return   [type]                   [description]
 */
func (c CommonController) GetUserShareNumByDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetRecords(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * 2.9 新增
 * 获取是否弹出提醒推送开关的状态
 * @Author   lifeng
 * @DateTime 2017-07-24T15:21:06+0800
 * @return   [type]                   [description]
 */
func (c CommonController) GetPopUpState(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) SavePopUpState(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * 获取分享窗口内容
 * @Author   lifeng
 * @DateTime 2017-08-04T19:58:53+0800
 * @return   [type]                   [description]
 */
func (c CommonController) GetSharebar(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * 获取弹窗内容
 */
func (c CommonController) GetPopup(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetHomePage(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetHomepageUserRecommend(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetMoreUserRecommendList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetHomepageIcon(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetHomepageLabel(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetCopywritingByKeyword(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetHomepageList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetGuidepageRecommend(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * 获取服务器时间
 */
func (c CommonController) GetTime(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) GetImgCode(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) CheckSecurityCode(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) ConsultConfig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) ConsultClose(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

func (c CommonController) IsConsultClosed(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}
