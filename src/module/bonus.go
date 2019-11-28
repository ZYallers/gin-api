package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/bonus"
	"src/library/tool"
)

type BonusModule struct {
	abs.Module
}

func Bonus() BonusModule {
	return BonusModule{}
}

func (a BonusModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		a.BindMethodOfController(gp,
			bonus.Bonus(),
			bonus.CopyWord(),
			bonus.OnlineService(),
			bonus.DailyTask(),
			bonus.Medal(),
			bonus.UserBean(),
			bonus.UserClockHealth(),
			bonus.UserScore(),
			bonus.UserSign(),
		)
	}
}
