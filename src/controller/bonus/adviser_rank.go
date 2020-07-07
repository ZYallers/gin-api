package bonus

import (
	"src/abs"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

// AdviserRankController 顾问成长值排行榜(462)
type AdviserRankController struct {
	abs.Controller
}

const adviserRankUri = "http://bonus.hxsapp.com/bean/AdviserRank/"

// AdviserRank 顾问成长值排行榜(462)
func AdviserRank() AdviserRankController {
	c := AdviserRankController{}
	c.Config = map[string]abs.MethodConfig{}
	return c
}

// GetMyPosition 获取最近一次排行榜我的排名
func (c AdviserRankController) GetMyPosition(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adviserRankUri+tool.CurrentMethodName())
}

// GrowthLog 按 id 倒序显示动态消息
func (c AdviserRankController) GrowthLog(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adviserRankUri+tool.CurrentMethodName())
}

// DoCommend 用户点赞 / 取消点赞
func (c AdviserRankController) DoCommend(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adviserRankUri+tool.CurrentMethodName())
}

// GetRanking 获取组长/组员排行
func (c AdviserRankController) GetRanking(ctx *gin.Context) {
	c.ServiceRewrite(ctx, adviserRankUri+tool.CurrentMethodName())
}