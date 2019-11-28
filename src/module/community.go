package module

import (
	"src/abs"
	"src/controller/community"
	"src/library/tool"

	"github.com/gin-gonic/gin"
)

type CommunityModule struct {
	abs.Module
}

func Community() CommunityModule {
	return CommunityModule{}
}

func (a CommunityModule) userDiaryUpper(rg *gin.RouterGroup, c *community.UserDiaryController) {
	gp := rg.Group("/UserDiary")
	{
		gp.GET("/deleteById", c.DeleteById)
		gp.POST("/deleteById", c.DeleteById)
		gp.GET("/hot", c.Hot)
		gp.POST("/hot", c.Hot)
		gp.GET("/getNewAuditTime",c.GetNewAuditTime)
		gp.POST("/getNewAuditTime",c.GetNewAuditTime)
		gp.GET("/getSelfList",c.GetSelfList)
		gp.POST("/getSelfList",c.GetSelfList)
	}
}

func (a CommunityModule) userCommentUpper(rg *gin.RouterGroup, c *community.UserCommentController) {
	gp := rg.Group("/userComment")
	{
		gp.GET("/getCommentChildrenList", c.GetCommentChildrenList)
		gp.POST("/getCommentChildrenList", c.GetCommentChildrenList)
	}
}

func (a CommunityModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{

		userDiary := community.UserDiary()
		a.userDiaryUpper(gp, &userDiary)

		userComment := community.UserComment()
		a.userCommentUpper(gp, &userComment)

		a.BindMethodOfController(gp,
			community.Center(),
			community.CircleList(),
			community.Common(),
			community.FriendsEvent(),
			community.PlatformWhite(),
			community.Power(),
			community.Recommend(),
			community.UserAccusation(),
			community.UserCollect(),
			community.UserCommend(),
			community.UserDiaryTag(),
			community.UserPullBlack(),
			community.UserRelate(),
			community.Vote(),
			userComment,
			userDiary,
		)
	}
}
