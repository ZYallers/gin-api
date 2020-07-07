package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/account"
	"src/library/tool"
)

type AccountModule struct {
	abs.Module
}

func Account() AccountModule {
	return AccountModule{}
}

/**
 * TODO::Controller Upper, 解决方案：待APP强制升级版本时，让客户端替换！
 */
func (a AccountModule) userAccountUpper(rg *gin.RouterGroup, c *account.UserAccountController) {
	api := rg.Group("/UserAccount")
	{
		api.POST("/login", c.Login)
		api.POST("/hasPwd", c.HasPwd)
		api.POST("/logout", c.Logout)
		api.POST("/sendSMS", c.SendSMS)
		api.POST("/userTimAccount", c.UserTimAccount)
		api.GET("/getOpenIMAccount", c.GetOpenIMAccount)
		api.POST("/getOpenIMAccount", c.GetOpenIMAccount)
		api.POST("/thirdPartyLogin", c.ThirdPartyLogin)
		api.POST("/register", c.Register)
		api.POST("/resetPwd", c.ResetPwd)
		api.POST("/thirdPartyBindInfo", c.ThirdPartyBindInfo)
		api.POST("/verifySmsCode", c.VerifySmsCode)
		api.POST("/thirdPartyBindOrUnbind", c.ThirdPartyBindOrUnbind)
		api.GET("/relateMobile", c.RelateMobile)
		api.POST("/relateMobile", c.RelateMobile)
		api.POST("/outsideLogin", c.OutsideLogin)
		api.POST("/registerTIm", c.RegisterTIm)
	}
}

/**
 * TODO::Controller Upper, 解决方案：待APP强制升级版本时，让客户端替换！
 */
func (a AccountModule) userInfoUpper(rg *gin.RouterGroup, c *account.UserInfoController) {
	gp := rg.Group("/UserInfo")
	{
		gp.POST("/saveInfo", c.SaveInfo)
		gp.GET("/getUserInfoByOpenImAccount", c.GetUserInfoByOpenImAccount)
		gp.POST("/getUserInfoByOpenImAccount", c.GetUserInfoByOpenImAccount)
		gp.POST("/getSelfUserInfo", c.GetSelfUserInfo)
		gp.GET("/getUserInfoByChatAccount", c.GetUserInfoByChatAccount)
		gp.POST("/getUserInfoByChatAccount", c.GetUserInfoByChatAccount)
		gp.POST("/getTalentRandom", c.GetTalentRandom)
		gp.POST("/getInviteCode", c.GetInviteCode)
		gp.POST("/chooseMobileToSummary", c.ChooseMobileToSummary)
		gp.POST("/getUserIdByOpenImAccount", c.GetUserIdByOpenImAccount)
		gp.POST("/getGroupChatAuth", c.GetGroupChatAuth)
		gp.POST("/mutiGetUserInfoByMobile", c.MutiGetUserInfoByMobile)
		gp.POST("/getInviteUserList", c.GetInviteUserList)
		gp.POST("/getUserInfo", c.GetUserInfo)
		gp.GET("/getUserInfo", c.GetUserInfo)
	}
}

/**
 * TODO::Controller Upper, 解决方案：待APP强制升级版本时，让客户端替换！
 */
func (a AccountModule) userPotoWall(rg *gin.RouterGroup, c *account.UserPhotoWallController) {
	gp := rg.Group("/userPhotoWall")
	{
		gp.POST("/getPhotoWall", c.GetPhotoWall)
		gp.POST("/photoWall", c.PhotoWall)
	}
}

func (a AccountModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		userAccount := account.UserAccount()
		a.userAccountUpper(gp, &userAccount)

		userInfo := account.UserInfo()
		a.userInfoUpper(gp, &userInfo)

		userPotoWall := account.UserPhotoWall()
		a.userPotoWall(gp, &userPotoWall)

		a.BindMethodOfController(gp,
			account.Adviser(),
			account.Channel(),
			account.Power(),
			account.UserLetter(),
			account.UserMember(),
			userPotoWall,
			userAccount,
			userInfo,
			account.Usertag(),
			account.Equity(),
		)
	}
}
