package module

import (
	"code/app/abst"
	"code/app/tool"
	"code/controller/search"
	"github.com/gin-gonic/gin"
)

type SearchModule struct {
	abst.Module
}

func Search() SearchModule {
	return SearchModule{}
}

func (a SearchModule) Group(rg *gin.RouterGroup) {
	api := rg.Group("/" + tool.CurrentFileName())
	{
		a.BindMethodOfController(api,
			search.ChatSearch(),
			search.ContactSearch(),
			search.FixSearch(),
			search.MallSearch(),
		)
	}
}
