package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CategoryController struct {
	abs.Controller
}

const categoryUri = "http://mall.hxsapp.com/base/Category/"

func Category() CategoryController {
	c := CategoryController{}
	c.Config = map[string]abs.MethodConfig{
		"GetCategory":  {ControllerNameFirstUpper: true},
		"GetStarGoods": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c CategoryController) GetCategory(ctx *gin.Context) {
	c.ServiceRewrite(ctx, categoryUri+tool.CurrentMethodName())
}

/**
 * @Author   pengxun
 * @DateTime 2018-12-18T11:17:21+0800
 * @desc     明星产品
 * @license  [license]
 * @version  [version]
 * @return   [type]                   [description]
 */
func (c CategoryController) GetStarGoods(ctx *gin.Context) {
	c.ServiceRewrite(ctx, categoryUri+tool.CurrentMethodName())
}
