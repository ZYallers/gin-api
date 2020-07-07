package content

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type NcovController struct {
	abs.Controller
}

const ncovUri = "http://content.hxsapp.com/article/Ncov/"
//
func Ncov() NcovController {
	c := NcovController{}
	return c
}

func (c NcovController) GetNews(ctx *gin.Context) {
	c.ServiceRewrite(ctx, ncovUri+tool.CurrentMethodName())
}
