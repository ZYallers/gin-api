package v100

import (
	"code/app/abst"
	"code/app/logger"
	"code/model/user"
	"code/service/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type userInfo struct {
	abst.Controller
}

func NewUserInfo(e *gin.Engine) *userInfo {
	ui := &userInfo{}
	ui.SetEngine(e)
	return ui
}

func (ui *userInfo) BaseInfo(c *gin.Context) {
	userIdStr, _ := c.GetQuery("uid")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(600, gin.H{"code": 600, "msg": "非法传参！"})
	} else {
		userInfo := service.NewUserInfo().GetUserInfoByUserId(userId)
		logger.Info("userinfo", "mobile", zap.String("mobile", userInfo.Mobile))
		c.JSON(200, gin.H{"code": 200, "msg": "ok", "data": gin.H{"mobile": userInfo.Mobile}})
	}
}

func (ui *userInfo) LoginLog(c *gin.Context) {
	userIdStr, _ := c.GetQuery("uid")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(600, gin.H{"code": 600, "msg": "非法传参！"})
	} else {
		data := model.NewUserLoginLog().FetchByUserIdAndLimit(userId, "id,user_id,last_log_ip", 10, "id desc")
		c.JSON(200, gin.H{"code": 200, "msg": "ok", "data": gin.H{"last_login_data": data}})
	}
}
