package module

import (
	"github.com/gin-gonic/gin"
	"src/controller/test"
	"src/libraries/core"
	"src/libraries/helper"
)

type Test struct {
	core.Module
}

func (t *Test) Group(eg *gin.Engine) {
	moduleName := helper.CurrentFileName()
	t.BindMethodOfController(eg, moduleName, &test.Index{})
}
