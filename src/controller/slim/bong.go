package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type BongController struct {
	abs.Controller
}

const bongUri = "http://hardware.hxsapp.com/device/Bong/"

func Bong() BongController {
	c := BongController{}
	c.Config = map[string]abs.MethodConfig{
		"GetUserStepByRecordDate":  {ControllerNameFirstUpper: true},
		"GetUserStepLastXDays":     {ControllerNameFirstUpper: true},
		"GetUserHeartLastXDays":    {ControllerNameFirstUpper: true},
		"GetUserSleepByRecordDate": {ControllerNameFirstUpper: true},
		"GetStepRankTen":           {ControllerNameFirstUpper: true},
		"SaveAppleData":            {ControllerNameFirstUpper: true},
	}
	return c
}

/*
 * 用户上次同步数据时间戳
 */
func (c BongController) GetUserLastSyncTime(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 检查手环是否被绑定
 */
func (c BongController) CheckBongIsBind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 手环解绑
 */
func (c BongController) UnbindUserBong(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 获取用户手环配置
 */
func (c BongController) GetUserBongConfig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 获取用户N天手环步数记录
 */
func (c BongController) GetUserStepLastXDays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 获取用户某个日期手环步数记录
 */
func (c BongController) GetUserStepByRecordDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 获取用户N天睡眠记录
 */
func (c BongController) GetUserSleepLastXDays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 获取用户某个日期的睡眠记录
 */
func (c BongController) GetUserSleepByRecordDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 获取用户N天的心率记录
 */
func (c BongController) GetUserHeartLastXDays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 同步手环记录（步数、心率、睡眠）
 */
func (c BongController) SaveBongData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 保存手环个性化配置
 */
func (c BongController) SaveBongConfig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 获取步数排行榜前10名
 */
func (c BongController) GetStepRankTen(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 保存苹果步数
 */
func (c BongController) SaveAppleData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 保存安卓步数
 */
func (c BongController) SaveAndroidData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 获取用户排行榜开关
 */
func (c BongController) GetUserStepRankSwitch(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 设置用户排行榜开关
 */
func (c BongController) SetUserStepRankSwitch(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}

/*
 * 步数排行榜分页列表
 */
func (c BongController) GetStepRankByPage(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bongUri+tool.CurrentMethodName())
}
