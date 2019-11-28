package content

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ArticleController struct {
	abs.Controller
}

const articleUri = "http://content.hxsapp.com/article/Article/"

func Article() ArticleController {
	c := ArticleController{}
	c.Config = map[string]abs.MethodConfig{
		"Detail":                    {ControllerNameFirstUpper: true},
		"GetDetailByAid":            {ControllerNameFirstUpper: true},
		"GetListByTypeId":           {ControllerNameFirstUpper: true},
		"GetVideoRecommend":         {ControllerNameFirstUpper: true},
		"GetRandomArticleByTagName": {ControllerNameFirstUpper: true},
		"GetTypeList":               {ControllerNameFirstUpper: true},
	}
	return c
}

func (c ArticleController) Detail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, articleUri+tool.CurrentMethodName())
}

func (c ArticleController) GetAggregationList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, articleUri+tool.CurrentMethodName())
}

func (c ArticleController) GetDetailByAid(ctx *gin.Context) {
	c.ServiceRewrite(ctx, articleUri+tool.CurrentMethodName())
}

func (c ArticleController) GetListByTypeId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, articleUri+tool.CurrentMethodName())
}

func (c ArticleController) GetVideoRecommend(ctx *gin.Context) {
	c.ServiceRewrite(ctx, articleUri+tool.CurrentMethodName())
}

func (c ArticleController) GetRandomArticleByTagName(ctx *gin.Context) {
	c.ServiceRewrite(ctx, articleUri+tool.CurrentMethodName())
}

func (c ArticleController) GetTypeList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, articleUri+tool.CurrentMethodName())
}
