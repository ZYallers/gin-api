package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/test"
	"src/library/tool"
)

type TestModule struct {
	abs.Module
}

func Test() TestModule {
	return TestModule{}
}

func (t TestModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		t.BindMethodOfController(gp, test.Index())
	}
}
