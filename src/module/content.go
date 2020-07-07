package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/content"
	"src/library/tool"
)

type ContentModule struct {
	abs.Module
}

func Content() ContentModule {
	return ContentModule{}
}

func (a ContentModule) mediaUpper(rg *gin.RouterGroup, c *content.MediaController) {
	gp := rg.Group("/Media")
	{
		gp.POST("/incrMediaPlay", c.IncrMediaPlay)
		gp.GET("/incrMediaPlay", c.IncrMediaPlay)
	}
}

func (a ContentModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		media := content.Media()
		a.mediaUpper(gp, &media)

		a.BindMethodOfController(gp,
			content.Article(),
			content.Cases(),
			content.Fitness(),
			content.Gymnastics(),
			content.Ncov(),
			media,
		)
	}
}
