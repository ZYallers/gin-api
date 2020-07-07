package base

import (
	"src/abs"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

// BottomItemController App 底部按钮
type BottomItemController struct {
	abs.Controller
}

const bottomItemUri = "http://base.hxsapp.com/base/BottomItem/"

// BottomItem App底部按钮
func BottomItem() BottomItemController {
	c := BottomItemController{}
	c.Config = map[string]abs.MethodConfig{}
	return c
}

// GetConfig 获取后台按钮配置
func (c BottomItemController) GetConfig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, bottomItemUri+tool.CurrentMethodName())
}
