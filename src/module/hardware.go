package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/hardware"
	"src/library/tool"
)

type HardwareModule struct {
	abs.Module
}

func Hardware() HardwareModule {
	return HardwareModule{}
}

func (a HardwareModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		a.BindMethodOfController(gp,
			hardware.Config(),
			hardware.Shoes(),
			hardware.Dumbbell(),
			hardware.Manage(),
			hardware.Kscale(),
			hardware.Massage(),
		)
	}
}
