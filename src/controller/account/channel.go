package account

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ChannelController struct {
	abs.Controller
}

const channelUri = "http://account.hxsapp.com/api/channel/"

func Channel() ChannelController {
	c := ChannelController{}
	c.Config = map[string]abs.MethodConfig{
		"Adclick":     {ControllerNameFirstUpper: true},
		"UniqueCheck": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 当用户发生点击的时候，趣米将把用户的mac和idfa按照指定格式通知给广告主：
 */
func (c ChannelController) Adclick(ctx *gin.Context) {
	c.ServiceRewrite(ctx, channelUri+tool.CurrentMethodName())
}

/**
 * 排重接口
 * 当用户在渠道平台点击广告时，为提高转化率，渠道需先对idfa排重（该接口只用于查询设备是否已经下载过，并不做其他操作）。
 */
func (c ChannelController) UniqueCheck(ctx *gin.Context) {
	c.ServiceRewrite(ctx, channelUri+tool.CurrentMethodName())
}

/**
 * 广告查询返回
 */
func (c ChannelController) Adcallback(ctx *gin.Context) {
	c.ServiceRewrite(ctx, channelUri+tool.CurrentMethodName())
}
