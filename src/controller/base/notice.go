package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type NoticeController struct {
	abs.Controller
}

const noticeUri = "http://base.hxsapp.com/message/notice/"

func Notice() NoticeController {
	c := NoticeController{}
	return c
}

/**
 * [getNoticeById 根据id获取单条通知]
 * @return [type] [description]
 */
func (c NoticeController) GetNoticeById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, noticeUri+tool.CurrentMethodName())
}

/**
 * [getNoticeList 取系统消息列表]
 * @return [type] [description]
 */
func (c NoticeController) GetNoticeList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, noticeUri+tool.CurrentMethodName())
}

/**
 * [getNoticeTotalNum 取系统消息总数]
 * @return [type] [description]
 */
func (c NoticeController) GetNoticeTotalNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, noticeUri+tool.CurrentMethodName())
}

/**
 * [isUserHaveUnread 判断用户是否有新未读通知]
 * @return boolean [description]
 */
func (c NoticeController) IsUserHaveUnread(ctx *gin.Context) {
	c.ServiceRewrite(ctx, noticeUri+tool.CurrentMethodName())
}

/**
 * [getUnreadNoticeTotalNum 取用户未读通知数]
 * @return [type] [description]
 */
func (c NoticeController) GetUnreadNoticeTotalNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, noticeUri+tool.CurrentMethodName())
}

/**
 * [updateUserReadNum 清空用户未读通知数]
 * @return [type] [description]
 */
func (c NoticeController) UpdateUserReadNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, noticeUri+tool.CurrentMethodName())
}

/**
 * [updateUserRead 更新用户指定通知的状态为已读]
 * @return [type] [description]
 */
func (c NoticeController) UpdateUserRead(ctx *gin.Context) {
	c.ServiceRewrite(ctx, noticeUri+tool.CurrentMethodName())
}
