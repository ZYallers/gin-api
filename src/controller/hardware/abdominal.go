package hardware

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type AbdominalController struct {
	abs.Controller
}

const abdominalUri = "http://hardware.hxsapp.com/device/abdominal/"

func Abdominal() AbdominalController {
	c := AbdominalController{}
	return c
}

func (c AbdominalController) CheckIsBind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, abdominalUri+tool.CurrentMethodName())
}

func (c AbdominalController) Bind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, abdominalUri+tool.CurrentMethodName())
}

func (c AbdominalController) Unbind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, abdominalUri+tool.CurrentMethodName())
}

func (c AbdominalController) SaveData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, abdominalUri+tool.CurrentMethodName())
}

func (c AbdominalController) Index(ctx *gin.Context) {
	c.ServiceRewrite(ctx, abdominalUri+tool.CurrentMethodName())
}

func (c AbdominalController) GetDailyListByXDays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, abdominalUri+tool.CurrentMethodName())
}

func (c AbdominalController) GetDetailByDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, abdominalUri+tool.CurrentMethodName())
}

func (c AbdominalController) GetAnalyseListByXDays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, abdominalUri+tool.CurrentMethodName())
}

func (c AbdominalController) GetAnalyseByDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, abdominalUri+tool.CurrentMethodName())
}
