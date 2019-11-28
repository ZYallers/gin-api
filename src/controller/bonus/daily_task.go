package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type DailyTaskController struct {
	abs.Controller
}

const dailyTaskUri = "http://bonus.hxsapp.com/dailyTask/DailyTask/"

func DailyTask() DailyTaskController {
	c := DailyTaskController{}
	c.Config = map[string]abs.MethodConfig{
		"GetTodayTaskData": {ControllerNameFirstUpper: true},
		"Draw":             {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 获取今日任务
 * @Author   lifeng
 * @DateTime 2017-09-16T11:23:44+0800
 * @return   [type]                   [description]
 */
func (c DailyTaskController) GetTodayTaskData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dailyTaskUri+tool.CurrentMethodName())
}

/**
 * 抽奖
 * @Author   lifeng
 * @DateTime 2017-09-16T11:32:50+0800
 * @return   [type]                   [description]
 */
func (c DailyTaskController) Draw(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dailyTaskUri+tool.CurrentMethodName())
}
