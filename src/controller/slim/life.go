package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type LifeController struct {
	abs.Controller
}

const lifeUri = "http://slim.hxsapp.com/slim/Life/"

func Life() LifeController {
	c := LifeController{}
	return c
}

/**
 * 获取生活日期
 */
func (c LifeController) GetLifeRecord(ctx *gin.Context) {
	c.ServiceRewrite(ctx, lifeUri+tool.CurrentMethodName())
}

/**
 * 保存生活日期
 */
func (c LifeController) SaveLifeRecord(ctx *gin.Context) {
	c.ServiceRewrite(ctx, lifeUri+tool.CurrentMethodName())
}
