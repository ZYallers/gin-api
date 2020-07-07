package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/im"
	"src/library/tool"
)

type ImModule struct {
	abs.Module
}

func Im() ImModule {
	return ImModule{}
}

func (a ImModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		a.BindMethodOfController(gp,
			im.Chat(),
			im.Diet(),
			im.YunXinIM(),
			im.Error(),
		)
	}
}
