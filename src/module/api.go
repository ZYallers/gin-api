package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/api"
	"src/library/tool"
)

type ApiModule struct {
	abs.Module
}

func Api() ApiModule {
	return ApiModule{}
}

/**
 * TODO::Controller Upper, 解决方案：待APP强制升级版本时，让客户端替换！
 */
func (a ApiModule) apiTimUpper(rg *gin.RouterGroup, c *api.TimController) {
	gp := rg.Group("/Tim")
	{
		gp.POST("/afterSaleSystemNotice", c.AfterSaleSystemNotice)
		gp.POST("/fixedUser", c.FixedUser)
	}
}

func (a ApiModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		apiTim := api.Tim()
		a.apiTimUpper(gp, &apiTim)

		a.BindMethodOfController(gp,
			api.Brm(),
			apiTim,
		)
	}
}
