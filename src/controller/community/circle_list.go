package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CircleListController struct {
	abs.Controller
}

const circleListUri = "http://community.hxsapp.com/circle/circleList/"

func CircleList() CircleListController {
	c := CircleListController{}
	return c
}

/**
 * 获取圈子分类列表
 */
func (c CircleListController) GetCircleTypeList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}

/**
 * 获取圈子分类的信息列表
 */
func (c CircleListController) GetCircleTypeInfoList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}

/**
 * 保存用户圈子信息(加入圈子)
 */
func (c CircleListController) SaveUserCircleInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}

/**
 * 圈子-内首页信息
 */
func (c CircleListController) GetCircleHomeInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}

/**
 * 获取圈子详情
 */
func (c CircleListController) GetCircleById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}

/**
 * 修改圈子内的公告
 */
func (c CircleListController) SaveCircleAffiche(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}

/**
 * 获取我的圈子列表
 */
func (c CircleListController) GetMyCircleList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}

/**
 * 获取首页,我的圈子
 */
func (c CircleListController) GetHomeMyCircle(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}

/**
 * 圈子首页-推荐圈子
 */
func (c CircleListController) GetCircleRecommendList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}

/**
 * 退出圈子
 */
func (c CircleListController) OutCircle(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}

/**
 * 获取默认圈子
 */
func (c CircleListController) GetDefaultCircle(ctx *gin.Context) {
	c.ServiceRewrite(ctx, circleListUri+tool.CurrentMethodName())
}
