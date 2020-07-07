package base

import (
	"src/abs"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

// GridlyController v460 首页标签信息流
type GridlyController struct {
	abs.Controller
}

const gridlyUri = "http://base.hxsapp.com/bbs/Gridly/"

func Gridly() GridlyController {
	c := GridlyController{}
	return c
}

// Detail 内容详情
func (c GridlyController) Detail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gridlyUri+tool.CurrentMethodName())
}

// ExtendDetail 内容详情扩展字段, 判断用户是否收藏、点赞过指定的内容
func (c GridlyController) ExtendDetail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gridlyUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2020-04-18 09:58:35
 * @Description: 标签信息流
 */
func (c GridlyController) GridlyList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gridlyUri+tool.CurrentMethodName())
}

// GetSocialData 获取点赞和收藏数
func (c GridlyController) GetSocialData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gridlyUri+tool.CurrentMethodName())
}

// DoCommend 点赞
func (c GridlyController) DoCommend(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gridlyUri+tool.CurrentMethodName())
}

// DoCollect 收藏
func (c GridlyController) DoCollect(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gridlyUri+tool.CurrentMethodName())
}

// GetMyCollect 收藏
func (c GridlyController) GetMyCollect(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gridlyUri+tool.CurrentMethodName())
}

// DoHotPoint 热点
func (c GridlyController) DoHotPoint(ctx *gin.Context) {
	c.ServiceRewrite(ctx, gridlyUri+tool.CurrentMethodName())
}

