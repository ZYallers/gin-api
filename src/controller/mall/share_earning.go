package mall

import (
	"src/abs"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

// ShareEarningController 分享赚
type ShareEarningController struct {
	abs.Controller
}

const shareEarningUri = "http://mall.hxsapp.com/base/ShareEarning/"

// ShareEarning 构造函数
func ShareEarning() ShareEarningController {
	c := ShareEarningController{}
	return c
}

// WebGetGoodsList H5 获取所有分享赚商品
func (s ShareEarningController) WebGetGoodsList(ctx *gin.Context) {
	s.ServiceRewrite(ctx, shareEarningUri+tool.CurrentMethodName())
}

// WebEarningBean H5 用户领取绿豆奖励
func (s ShareEarningController) WebEarningBean(ctx *gin.Context) {
	s.ServiceRewrite(ctx, shareEarningUri+tool.CurrentMethodName())
}

// Callback 用户分享到微信后回调这个接口，用于奖励用户绿豆
func (s ShareEarningController) Callback(ctx *gin.Context) {
	s.ServiceRewrite(ctx, shareEarningUri+tool.CurrentMethodName())
}

// WebCallback H5 用户分享到微信后回调这个接口，用于奖励用户绿豆
func (s ShareEarningController) WebCallback(ctx *gin.Context) {
	s.ServiceRewrite(ctx, shareEarningUri+tool.CurrentMethodName())
}
