package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserMessageController struct {
	abs.Controller
}

const userMessageUri = "http://base.hxsapp.com/message/userMessage/"

func UserMessage() UserMessageController {
	c := UserMessageController{}
	return c
}

/**
 * [getMessageList 获取用户消息列表]
 * @return [type] [description]
 */
func (c UserMessageController) GetMessageList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMessageUri+tool.CurrentMethodName())
}

/**
 * [getUnreadMessageTotal 获取未读消息总数]
 * @return [type] [description]
 */
func (c UserMessageController) GetUnreadMessageTotal(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMessageUri+tool.CurrentMethodName())
}

/**
 * [getUserAllUnreadMessageInfo 取用户所有类型未读消息信息]
 * @return [type] [description]
 */
func (c UserMessageController) GetUserAllUnreadMessageInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMessageUri+tool.CurrentMethodName())
}

/**
 * [updateMessageReadStatus 将消息置为已读]
 * @return [type] [description]
 */
func (c UserMessageController) UpdateMessageReadStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMessageUri+tool.CurrentMethodName())
}

/**
 * [updateMessageReadStatusByAction 根据action将消息置为已读]
 * @return [type] [description]
 */
func (c UserMessageController) UpdateMessageReadStatusByAction(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMessageUri+tool.CurrentMethodName())
}

/**
 * [updateOneMessageReadStatus 将单条消息置为已读]
 * @return [type] [description]
 */
func (c UserMessageController) UpdateOneMessageReadStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userMessageUri+tool.CurrentMethodName())
}
