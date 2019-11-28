package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/search"
	"src/library/tool"
)

type SearchModule struct {
	abs.Module
}

func Search() SearchModule {
	return SearchModule{}
}

func (a SearchModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		a.BindMethodOfController(gp,
			search.ChatSearch(),
			search.ContactSearch(),
			search.FixSearch(),
			search.MallSearch(),
		)
	}
}
