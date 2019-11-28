package hardware

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type MassageController struct {
	abs.Controller
}

const massageUri = "http://hardware.hxsapp.com/device/massage/"

func Massage() MassageController {
	c := MassageController{}
	return c
}

func (c MassageController) CheckIsBind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, massageUri+tool.CurrentMethodName())
}

func (c MassageController) Bind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, massageUri+tool.CurrentMethodName())
}

func (c MassageController) Unbind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, massageUri+tool.CurrentMethodName())
}

func (c MassageController) SaveData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, massageUri+tool.CurrentMethodName())
}

func (c MassageController) Index(ctx *gin.Context) {
	c.ServiceRewrite(ctx, massageUri+tool.CurrentMethodName())
}

func (c MassageController) GetDailyListByXDays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, massageUri+tool.CurrentMethodName())
}

func (c MassageController) GetDetailByDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, massageUri+tool.CurrentMethodName())
}

