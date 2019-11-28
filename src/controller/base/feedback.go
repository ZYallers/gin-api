package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type FeedbackController struct {
	abs.Controller
}

const feedbackUri = "http://base.hxsapp.com/base/feedback/"

func Feedback() FeedbackController {
	c := FeedbackController{}
	c.Config = map[string]abs.MethodConfig{
		"SaveFeedback": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * [saveFeedback 保存用户意见反馈]
 * @return [type] [description]
 */
func (c FeedbackController) SaveFeedback(ctx *gin.Context) {
	c.ServiceRewrite(ctx, feedbackUri+tool.CurrentMethodName())
}

/**
 * [getUserFeedback 获取用户意见反馈列表]
 * @return [type] [description]
 */
func (c FeedbackController) GetUserFeedback(ctx *gin.Context) {
	c.ServiceRewrite(ctx, feedbackUri+tool.CurrentMethodName())
}

/**
 * [doReadFeedback 把用户未读意见反馈置为已读]
 * @return [type] [description]
 */
func (c FeedbackController) DoReadFeedback(ctx *gin.Context) {
	c.ServiceRewrite(ctx, feedbackUri+tool.CurrentMethodName())
}
