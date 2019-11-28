package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserClockHealthController struct {
	abs.Controller
}

const userClockHealthUri = "http://bonus.hxsapp.com/bean/UserClockHealth/"

func UserClockHealth() UserClockHealthController {
	c := UserClockHealthController{}
	c.Config = map[string]abs.MethodConfig{
		"GetUserClockIndex":           {ControllerNameFirstUpper: true},
		"GetUserHealthState":          {ControllerNameFirstUpper: true},
		"SetUserHealthState":          {ControllerNameFirstUpper: true},
		"SetUserSport":                {ControllerNameFirstUpper: true},
		"GetUserSport":                {ControllerNameFirstUpper: true},
		"SetUserWeigh":                {ControllerNameFirstUpper: true},
		"GetUserWeighList":            {ControllerNameFirstUpper: true},
		"GetUserMeasurList":           {ControllerNameFirstUpper: true},
		"SetUserMeasur":               {ControllerNameFirstUpper: true},
		"SetUserMeasurWithSync":       {ControllerNameFirstUpper: true},
		"SetUserWeighToClock":         {ControllerNameFirstUpper: true},
		"SetUserDiet":                 {ControllerNameFirstUpper: true},
		"GetUserDiet":                 {ControllerNameFirstUpper: true},
		"GetClockStatus":              {ControllerNameFirstUpper: true},
		"DeleteDietItem":              {ControllerNameFirstUpper: true},
		"GetRewardByType":             {ControllerNameFirstUpper: true},
		"GetHadClockCountDay":         {ControllerNameFirstUpper: true},
		"GetUserDietHistory":          {ControllerNameFirstUpper: true},
		"SetUserCustomIntake":         {ControllerNameFirstUpper: true},
		"GetUserDietByUserId":         {ControllerNameFirstUpper: true},
		"SetUserDietByUserId":         {ControllerNameFirstUpper: true},
		"DeleteDietItemByUserId":      {ControllerNameFirstUpper: true},
		"SetUserCustomIntakeByUserId": {ControllerNameFirstUpper: true},
		"GetUserDietPhotoByUserId":    {ControllerNameFirstUpper: true},
		"GetUserConfigIntakeByUserId": {ControllerNameFirstUpper: true},
		"DeleteDietPhotoByUserId":     {ControllerNameFirstUpper: true},
		"UpdateUserDietItem":          {ControllerNameFirstUpper: true},
		"ResetUserDietByUserId":       {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 今日打卡主页
 * @Author   lifeng
 * @DateTime 2017-09-16T11:35:54+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) GetUserClockIndex(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 获取身体状况
 * @Author   lifeng
 * @DateTime 2017-09-16T11:42:53+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) GetUserHealthState(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 设置身体状况
 * @Author   lifeng
 * @DateTime 2017-09-16T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) SetUserHealthState(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 设置 用户今日运动
 * @Author   lifeng
 * @DateTime 2017-09-16T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) SetUserSport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 获取 用户今日运动
 * @Author   lifeng
 * @DateTime 2017-09-16T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) GetUserSport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 设置 用户今日称重
 * @Author   lifeng
 * @DateTime 2017-09-16T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) SetUserWeigh(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 获取 用户称重列表
 * @Author   lifeng
 * @DateTime 2017-09-16T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) GetUserWeighList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 获取 用户三围数据
 * @Author   lingjijian
 * @DateTime 2018-12-21T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) GetUserMeasurList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 设置 用户三围数据
 * @Author   lingjijian
 * @DateTime 2018-12-21T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) SetUserMeasur(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) SetUserMeasurWithSync(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) SetUserWeighToClock(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 设置 用户饮食
 * @Author   lingjijian
 * @DateTime 2018-12-21T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) SetUserDiet(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 获取 用户饮食
 * @Author   lingjijian
 * @DateTime 2018-12-21T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) GetUserDiet(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

/**
 * 获取 今日打卡状况
 * @Author   lingjijian
 * @DateTime 2018-12-21T11:51:51+0800
 * @return   [type]                   [description]
 */
func (c UserClockHealthController) GetClockStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) DeleteDietItem(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) GetRewardByType(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) GetHadClockCountDay(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) GetUserDietHistory(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) SetUserCustomIntake(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) GetUserDietByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) SetUserDietByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) DeleteDietItemByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) SetUserCustomIntakeByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) GetUserDietPhotoByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) GetUserConfigIntakeByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) DeleteDietPhotoByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) UpdateUserDietItem(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}

func (c UserClockHealthController) ResetUserDietByUserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userClockHealthUri+tool.CurrentMethodName())
}
