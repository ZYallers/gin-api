package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/live"
	"src/library/tool"
)

type LiveModule struct {
	abs.Module
}

func Live() LiveModule {
	return LiveModule{}
}

func (a LiveModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		a.BindMethodOfController(gp,
			live.Yunxin(),
			live.Pop(),
			live.Room(),
			live.Present(),
		)
	}
}
