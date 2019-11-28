package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/home"
	"src/library/tool"
)

type HomeModule struct {
	abs.Module
}

func Home() HomeModule {
	return HomeModule{}
}

func (a HomeModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		a.BindMethodOfController(gp,
			home.Weight(),
		)
	}
}
