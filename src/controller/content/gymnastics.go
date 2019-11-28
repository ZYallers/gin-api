package content

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type GymnasticsController struct {
	abs.Controller
}

const gymnasticsUri = "http://content.hxsapp.com/course/Gymnastics/"

func Gymnastics() GymnasticsController {
	c := GymnasticsController{}
	c.Config = map[string]abs.MethodConfig{
		"GetCourseList":        {ControllerNameFirstUpper: true},
		"GetCourseDetail":      {ControllerNameFirstUpper: true},
		"GetTrain":             {ControllerNameFirstUpper: true},
		"GetFragmentTrain":     {ControllerNameFirstUpper: true},
		"GetUserStatus":        {ControllerNameFirstUpper: true},
		"UserViewTrainStatus":  {ControllerNameFirstUpper: true},
		"TrainViewUser":        {ControllerNameFirstUpper: true},
		"GetUserLastCourse":    {ControllerNameFirstUpper: true},
		"GetHomeCourseList":    {ControllerNameFirstUpper: true},
		"GetHomeFragmentTrain": {ControllerNameFirstUpper: true},
		"CourseViewUser":       {ControllerNameFirstUpper: true},
	}
	return c
}

func (c GymnasticsController) GetCourseList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}

func (c GymnasticsController) GetCourseDetail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}

func (c GymnasticsController) GetTrain(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}

func (c GymnasticsController) GetFragmentTrain(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}

func (c GymnasticsController) GetUserStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}

func (c GymnasticsController) UserViewTrainStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}

func (c GymnasticsController) TrainViewUser(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}

func (c GymnasticsController) GetUserLastCourse(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}

func (c GymnasticsController) GetHomeCourseList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}

func (c GymnasticsController) GetHomeFragmentTrain(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}

func (c GymnasticsController) CourseViewUser(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gymnasticsUri+tool.CurrentMethodName())
}
