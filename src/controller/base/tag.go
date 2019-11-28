package base

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type TagController struct {
	abs.Controller
}

const tagUri = "http://community.hxsapp.com/common/tag/"

func Tag() TagController {
	c := TagController{}
	c.Config = map[string]abs.MethodConfig{
		"GetTagByName":         {ControllerNameFirstUpper: true},
		"SaveTag":              {ControllerNameFirstUpper: true},
		"GetTagList":           {ControllerNameFirstUpper: true},
		"CategoryTagList":      {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 通过ID获取对应标签
 */
func (c TagController) GetTagById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, tagUri+tool.CurrentMethodName())
}

/**
 * 根据名字查询标签
 */
func (c TagController) GetTagByName(ctx *gin.Context) {
	c.ServiceRewrite(ctx, tagUri+tool.CurrentMethodName())
}

/**
 * 用户保存自定义标签
 */
func (c TagController) SaveTag(ctx *gin.Context) {
	c.ServiceRewrite(ctx, tagUri+tool.CurrentMethodName())
}

/**
 * 获取标签列表
 */
func (c TagController) GetTagList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, tagUri+tool.CurrentMethodName())
}

/**
 * 获取话题全部人气列表
 */
func (c TagController) GetTagPopularityList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, tagUri+tool.CurrentMethodName())
}

/**
 * 首页-话题-热门话题(按照排序号)
 */
func (c TagController) GetTagSortList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, tagUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2019-03-30 16:03:50
 * @Description: 分类下的话题列表
 */
func (c TagController) CategoryTagList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, tagUri+tool.CurrentMethodName())
}
