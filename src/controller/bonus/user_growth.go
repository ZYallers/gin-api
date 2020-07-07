package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type UserGrowthController struct {
	abs.Controller
}

const userGrowthUri = "http://bonus.hxsapp.com/bean/userGrowth/"

func UserGrowth() UserGrowthController {
	c := UserGrowthController{}
	c.Config = map[string]abs.MethodConfig{}
	return c
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2020-03-09 10:30:56
 * @Description: 成长值
 */
func (c UserGrowthController) TaskInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userGrowthUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2020-03-09 16:43:55
 * @Description: 领取成长值
 */
func (c UserGrowthController) DoGrowth(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userGrowthUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2020-03-09 10:30:25
 * @Description: 用户的成长值
 */
func (c UserGrowthController) Growth(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userGrowthUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2020-03-14 16:39:39
 * @Description: 成长值兑换 消费
 */
func (c UserGrowthController) Exchange(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userGrowthUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2020-03-16 11:26:54
 * @Description: 成长值明细
 */
func (c UserGrowthController) Detail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userGrowthUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2020-03-26 16:51:14
 * @Description: 成长值兑换规则说明
 */
func (c UserGrowthController) Notice(ctx *gin.Context) {
	c.ServiceRewrite(ctx, userGrowthUri+tool.CurrentMethodName())
}
