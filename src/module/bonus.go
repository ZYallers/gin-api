package module

import (
	"src/abs"
	"src/controller/bonus"
	"src/library/tool"

	"github.com/gin-gonic/gin"
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
			bonus.AdviserRank(),
			bonus.Bonus(),
			bonus.CopyWord(),
			bonus.DailyTask(),
			bonus.GroupBusiness(),
			bonus.Medal(),
			bonus.NewOnlineService(),
			bonus.NewsTip(),
			bonus.OnlineService(),
			bonus.UserBean(),
			bonus.UserClockHealth(),
			bonus.UserGrowth(),
			bonus.UserScore(),
			bonus.UserSign(),
			bonus.WeightManagementUri(),
		)
	}
}
