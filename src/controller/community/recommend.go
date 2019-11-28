package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type RecommendController struct {
	abs.Controller
}

const recommendUri = "http://community.hxsapp.com/common/recommend/"

func Recommend() RecommendController {
	c := RecommendController{}
	return c
}

func (c RecommendController) GetRecommendList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, recommendUri+tool.CurrentMethodName())
}
