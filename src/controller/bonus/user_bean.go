package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserBeanController struct {
	abs.Controller
}

const userBeanUri = "http://bonus.hxsapp.com/bean/UserBean/"

func UserBean() UserBeanController {
	c := UserBeanController{}
	c.Config = map[string]abs.MethodConfig{
		"Bean":          {ControllerNameFirstUpper: true},
		"BeanTask":      {ControllerNameFirstUpper: true},
		"TaskData":      {ControllerNameFirstUpper: true},
		"GetTaskData":   {ControllerNameFirstUpper: true},
		"GetUserBean":   {ControllerNameFirstUpper: true},
		"GetBeanStatus": {ControllerNameFirstUpper: true},
		"GetBeanNum":    {ControllerNameFirstUpper: true},
		"BeanRank":      {ControllerNameFirstUpper: true},
		"Consume":       {ControllerNameFirstUpper: true},
		"OnlineTime":    {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:32
 * @Description: + - 好享豆
 */
func (c UserBeanController) Bean(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:36
 * @Description: 好享豆任务
 */
func (c UserBeanController) BeanTask(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:36
 * @Description: 日常任务操作
 */
func (c UserBeanController) TaskData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:36
 * @Description: 日常任务记录
 */
func (c UserBeanController) GetTaskData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:36
 * @Description: 获取多个用户的好享豆数目
 */
func (c UserBeanController) GetUserBean(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:36
 * @Description: 是否有好享豆可领取
 */
func (c UserBeanController) GetBeanStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-07-23 10:56:28
 * @Description: 未领取绿豆任务数
 */
func (c UserBeanController) GetBeanNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:36
 * @Description: 好享豆排行榜
 */
func (c UserBeanController) BeanRank(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:36
 * @Description: 购物消费
 */
func (c UserBeanController) Consume(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:36
 * @Description: 在线时长
 */
func (c UserBeanController) OnlineTime(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:36
 * @Description: 咨询时长
 */
func (c UserBeanController) ConsultTime(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2019-10-24 15:31:54
 * @Description: 4.2.6 弹窗领取绿豆
 */
func (c UserBeanController) SetBeanWithPopup(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userBeanUri+tool.CurrentMethodName())
}
