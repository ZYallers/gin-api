package api

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type BrmController struct {
	abs.Controller
}

func Brm() BrmController {
	c := BrmController{}
	c.Config = map[string]abs.MethodConfig{
		"GetUserSignStatus":         {ControllerNameFirstUpper: true},
		"GetAdviserInfoByAdviserId": {ControllerNameFirstUpper: true},
		"InQueueForAuthorize":       {ControllerNameFirstUpper: true},
		"InQueue":                   {ControllerNameFirstUpper: true},
	}
	return c
}

const (
	slimUri      = "http://slim.hxsapp.com/other/Brm/"
	communityUri = "http://community.hxsapp.com/api/brm/"
	imUri        = "http://im.hxsapp.com/api/Brm/"
)

func (c BrmController) GetGeneList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, slimUri+tool.CurrentMethodName())
}

func (c BrmController) GetUserDiaryList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, communityUri+tool.CurrentMethodName())
}

func (c BrmController) GetBodyReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, slimUri+tool.CurrentMethodName())
}

func (c BrmController) InQueue(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c BrmController) GetBrmUserInfoByEncodePhone(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c BrmController) InQueueForAuthorize(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c BrmController) GetAdviserInfoByAdviserId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c BrmController) GetUserSignStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}

func (c BrmController) GetLogByBrmId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, imUri+tool.CurrentMethodName())
}
