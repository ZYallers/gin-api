package search

import (
	"code/app/abst"
	"code/app/tool"
	"github.com/gin-gonic/gin"
)

type ContactSearchController struct {
	abst.Controller
}

const contactSearchUri = "http://search.hxsapp.com/api/ContactSearch/"

func ContactSearch() ContactSearchController {
	c := ContactSearchController{}
	c.Config = map[string]abst.MethodConfig{
		"DoSearch": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * @Author:      lingjijian
 * @DateTime:    2018-08-1 15:10:32
 * @Description: 搜索通讯录
 */
func (c ContactSearchController) DoSearch(ctx *gin.Context) {
	c.ServiceRewrite(ctx, contactSearchUri+tool.CurrentMethodName())
}
