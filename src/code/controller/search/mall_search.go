package search

import (
	"code/app/abst"
	"code/app/tool"
	"github.com/gin-gonic/gin"
)

type MallSearchController struct {
	abst.Controller
}

const mallSearchUri = "http://search.hxsapp.com/api/MallSearch/"

func MallSearch() MallSearchController {
	c := MallSearchController{}
	c.Config = map[string]abst.MethodConfig{
		"DoSearch": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * @Author:      lingjijian
 * @DateTime:    2019-03-26 15:10:32
 * @Description: 商品搜索
 */
func (c MallSearchController) DoSearch(ctx *gin.Context) {
	c.ServiceRewrite(ctx, mallSearchUri+tool.CurrentMethodName())
}
