package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type HomePageController struct {
	abs.Controller
}

const homePageUri = "http://base.hxsapp.com/base/HomePage/"

func HomePage() HomePageController {
	c := HomePageController{}
	c.Config = map[string]abs.MethodConfig{
		"GetHomePage":     {ControllerNameFirstUpper: true},
		"NewHomePage":     {ControllerNameFirstUpper: true},
		"CommunityPage":   {ControllerNameFirstUpper: true},
		"GetHomeFlamingo": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c HomePageController) GetHomePage(ctx *gin.Context) {
	c.ServiceRewrite(ctx, homePageUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-12-19 14:43:38
 * @Description: 首页楼层
 */
func (c HomePageController) NewHomePage(ctx *gin.Context) {
	c.ServiceRewrite(ctx, homePageUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-12-20 20:34:46
 * @Description: 社区页
 */
func (c HomePageController) CommunityPage(ctx *gin.Context) {
	c.ServiceRewrite(ctx, homePageUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2018-12-21 16:21:44
 * @Description: 首页火烈鸟
 */
func (c HomePageController) GetHomeFlamingo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, homePageUri+tool.CurrentMethodName())
}
