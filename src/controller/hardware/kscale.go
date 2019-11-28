package hardware

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type KscaleController struct {
	abs.Controller
}

const kscaleUri  = "http://hardware.hxsapp.com/device/Kscale/"

func Kscale()  KscaleController{
	c := KscaleController{}
	return c
}

func(c KscaleController) ProduceInit(ctx *gin.Context)  {
	c.ServiceRewrite(ctx, kscaleUri+tool.CurrentMethodName())
}

func(c KscaleController) Register(ctx *gin.Context)  {
	c.ServiceRewrite(ctx, kscaleUri+tool.CurrentMethodName())
}

func(c KscaleController) Configure(ctx *gin.Context)  {
	c.ServiceRewrite(ctx, kscaleUri+tool.CurrentMethodName())
}

func(c KscaleController) Upload(ctx *gin.Context)  {
	c.ServiceRewrite(ctx, kscaleUri+tool.CurrentMethodName())
}

func(c KscaleController) CheckUpdate(ctx *gin.Context)  {
	c.ServiceRewrite(ctx, kscaleUri+tool.CurrentMethodName())
}

func(c KscaleController) ErrorLog(ctx *gin.Context)  {
	c.ServiceRewrite(ctx, kscaleUri+tool.CurrentMethodName())
}