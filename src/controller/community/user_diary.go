package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserDiaryController struct {
	abs.Controller
}

const userDiaryUri = "http://community.hxsapp.com/user/userDiary/"

func UserDiary() UserDiaryController {
	c := UserDiaryController{}
	c.Config = map[string]abs.MethodConfig{
		"GetDiaryList":       {ControllerNameFirstUpper: true},
		"DiaryNum":           {ControllerNameFirstUpper: true},
		"DiaryStateList":     {ControllerNameFirstUpper: true},
		"SaveDiaryCircleTop": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 热门 动态
 */
func (c UserDiaryController) Hot(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * 获取动态列表
 */
func (c UserDiaryController) GetDiaryList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * 根据id获取动态详情
 */
func (c UserDiaryController) GetDiaryById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * 发送动态
 */
func (c UserDiaryController) SaveDiary(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * 获取用户自己的动态
 */
func (c UserDiaryController) GetSelfList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2018-08-07 09:49:44
 * @Description: 在线动态数，屏蔽动态数
 */
func (c UserDiaryController) DiaryNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2018-08-07 09:49:44
 * @Description: 在线动态/评论列表，屏蔽动态/评论列表, 禁言屏蔽列表
 */
func (c UserDiaryController) DiaryStateList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * 获取最新动态审核时间
 */
func (c UserDiaryController) GetNewAuditTime(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * 用户删除指定日记
 */
func (c UserDiaryController) DeleteById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * 获取动态同省瘦友
 */
func (c UserDiaryController) GetDiaryProvince(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * 获取圈子动态列表
 */
func (c UserDiaryController) GetCircleDiaryList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * 获取所有圈子动态的列表
 */
func (c UserDiaryController) GetAllCircleDiaryList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}

/**
 * 保存(圈子-动态-置顶)顶部动态
 */
func (c UserDiaryController) SaveDiaryCircleTop(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryUri+tool.CurrentMethodName())
}
