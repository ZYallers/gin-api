package module

import (
	"src/abs"
	"src/controller/bbs"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

type BbsModule struct {
	abs.Module
}

func Bbs() BbsModule {
	return BbsModule{}
}

func (a BbsModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		a.BindMethodOfController(gp,
			bbs.Diary(),
			bbs.Circle(),
			bbs.DiaryComment(),
			bbs.DiaryLike(),
			bbs.SpecialPopulation(),
		)
	}
}
