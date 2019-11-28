package search

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type FixSearchController struct {
	abs.Controller
}

const fixSearchUri = "http://search.hxsapp.com/api/FixSearch/"

func FixSearch() FixSearchController {
	c := FixSearchController{}
	c.Config = map[string]abs.MethodConfig{
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
