package module

import (
	"code/app/abst"
	"code/app/tool"
	"code/controller/test"
	"github.com/gin-gonic/gin"
)

type TestModule struct {
	abst.Module
}

func Test() TestModule {
	return TestModule{}
}

func (t TestModule) Group(rg *gin.RouterGroup) {
	api := rg.Group("/" + tool.CurrentFileName())
	{
		t.BindMethodOfController(api, test.Index())
	}
}
