package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type SportController struct {
	abs.Controller
}

const sportUri = "http://slim.hxsapp.com/slim/sport/"

func Sport() SportController {
	c := SportController{}
	return c
}

/**
 * 获取分类
 */
func (c SportController) GetCategory(ctx *gin.Context) {
	c.ServiceRewrite(ctx, sportUri+tool.CurrentMethodName())
}

/**
 * 搜索名称
 */
func (c SportController) Search(ctx *gin.Context) {
	c.ServiceRewrite(ctx, sportUri+tool.CurrentMethodName())
}

/**
 * 保存记录
 */
func (c SportController) SaveRecord(ctx *gin.Context) {
	c.ServiceRewrite(ctx, sportUri+tool.CurrentMethodName())
}

/**
 * 删除一条记录
 */
func (c SportController) DeleteRecord(ctx *gin.Context) {
	c.ServiceRewrite(ctx, sportUri+tool.CurrentMethodName())
}

/**
 * 获取指定运动的规格
 */
func (c SportController) GetUnit(ctx *gin.Context) {
	c.ServiceRewrite(ctx, sportUri+tool.CurrentMethodName())
}

/**
 * 获取运动方案
 */
func (c SportController) GetSportScheme(ctx *gin.Context) {
	c.ServiceRewrite(ctx, sportUri+tool.CurrentMethodName())
}
