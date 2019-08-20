package search

import (
	"code/app/abst"
	"code/app/tool"
	"github.com/gin-gonic/gin"
)

type FixSearchController struct {
	abst.Controller
}

const fixSearchUri = "http://search.hxsapp.com/api/FixSearch/"

func FixSearch() FixSearchController {
	c := FixSearchController{}
	c.Config = map[string]abst.MethodConfig{
		"Index": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * @Author:      caius
 * @DateTime:    2018-01-29 15:10:32
 * @Description: + - 好享豆
 */
func (c FixSearchController) Index(ctx *gin.Context) {
	c.ServiceRewrite(ctx, fixSearchUri+tool.CurrentMethodName())
}
