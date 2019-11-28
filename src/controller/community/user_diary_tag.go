package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserDiaryTagController struct {
	abs.Controller
}

const userDiaryTagUri = "http://community.hxsapp.com/user/userDiaryTag/"

func UserDiaryTag() UserDiaryTagController {
	c := UserDiaryTagController{}
	return c
}

/**
 * 获取用户话题列表
 */
func (c UserDiaryTagController) GetUserDiaryTagList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryTagUri+tool.CurrentMethodName())
}

/**
 * 获取用户动态话题详情页
 */
func (c UserDiaryTagController) GetUserDiaryTagById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryTagUri+tool.CurrentMethodName())
}

func (c UserDiaryTagController) GetTagRankWithWeek(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryTagUri+tool.CurrentMethodName())
}

func (c UserDiaryTagController) GetTagRankTagId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userDiaryTagUri+tool.CurrentMethodName())
}
