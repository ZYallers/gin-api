package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/base"
	"src/library/tool"
)

type BaseModule struct {
	abs.Module
}

func Base() BaseModule {
	return BaseModule{}
}

func (a BaseModule) freeBackUpper(rg *gin.RouterGroup, c *base.FeedbackController) {
	gp := rg.Group("/feedback")
	{
		gp.POST("/saveFeedback", c.SaveFeedback)
		gp.GET("/saveFeedback", c.SaveFeedback)
	}
}

func (a BaseModule) commonUpper(rg *gin.RouterGroup, c *base.CommonController) {
	gp := rg.Group("/Common")
	{
		gp.POST("/getWxShareJsApiSignature", c.GetWxShareJsApiSignature)
		gp.POST("/getImgCode", c.GetImgCode)
		gp.GET("/getImgCode", c.GetImgCode)
		gp.GET("/checkSecurityCode", c.CheckSecurityCode)
		gp.POST("/checkSecurityCode", c.CheckSecurityCode)
		gp.GET("/getHomePage", c.GetHomePage)
		gp.POST("/getHomePage", c.GetHomePage)
	}
}

func (a BaseModule) tagUpper(rg *gin.RouterGroup, c *base.TagController) {
	gp := rg.Group("/Tag")
	{
		gp.POST("/getTagById", c.GetTagById)
		gp.GET("/getTagById", c.GetTagById)
		gp.POST("getTagPopularityList", c.GetTagPopularityList)
		gp.GET("/getTagPopularityList", c.GetTagPopularityList)
	}
}

func (a BaseModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		feedback := base.Feedback()
		a.freeBackUpper(gp, &feedback)

		common := base.Common()
		a.commonUpper(gp, &common)

		tag := base.Tag()
		a.tagUpper(gp, &tag)

		a.BindMethodOfController(gp,
			base.Ad(),
			base.Android(),
			base.Assessment(),
			base.HomePage(),
			base.Notice(),
			base.Push(),
			base.UserMessage(),
			base.VisibleUser(),
			common,
			feedback,
			tag,
		)
	}
}
