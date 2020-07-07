package account

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserPhotoWallController struct {
	abs.Controller
}

const userPhotoWallUri = "http://account.hxsapp.com/user/userPhotoWall/"

func UserPhotoWall() UserPhotoWallController {
	c := UserPhotoWallController{}
	c.Config = map[string]abs.MethodConfig{
		"PhotoWall":    {ControllerNameFirstUpper: true},
		"GetPhotoWall": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2018-05-29 14:41:39
 * @Description: 编辑照片墙
 */
func (c UserPhotoWallController) PhotoWall(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userPhotoWallUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2018-05-29 14:41:42
 * @Description: 获取照片墙
 */
func (c UserPhotoWallController) GetPhotoWall(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userPhotoWallUri+tool.CurrentMethodName())
}

func (c UserPhotoWallController) UserPhotograph(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userPhotoWallUri+tool.CurrentMethodName())
}
