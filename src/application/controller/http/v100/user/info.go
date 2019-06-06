package v100

import (
	"application/app/abst"
	"application/app/logger"
	"application/model/user"
	"application/service/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type userInfo struct {
	abst.Controller
}

func NewUserInfo(engine *gin.Engine) *userInfo {
	controller := new(userInfo)
	controller.Egn = engine
	return controller
}

func (this *userInfo) BaseInfo(c *gin.Context) {
	this.Ctx = c
	userIdStr, _ := c.GetQuery("uid")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(600, gin.H{"code": 600, "msg": "非法传参！"})
	} else {
		userInfo := service.NewUserInfo(c).GetUserInfoByUserId(userId)
		logger.Info("userinfo", "mobile", zap.String("mobile", userInfo.Mobile))
		c.JSON(200, gin.H{"code": 200, "msg": "ok", "data": gin.H{"mobile": userInfo.Mobile}})
	}
}

func (this *userInfo) LoginLog(c *gin.Context) {
	this.Ctx = c
	userIdStr, _ := c.GetQuery("uid")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(600, gin.H{"code": 600, "msg": "非法传参！"})
	} else {
		data := model.NewUserLoginLog(c).FetchByUserIdAndLimit(userId, "id,user_id,last_log_ip", 10, "id desc")
		data2 := model.NewUserLoginLog(c).FetchByUserIdAndLimit(2, "id,user_id,last_log_ip", 10, "id desc")
		c.JSON(200, gin.H{"code": 200, "msg": "ok", "data": gin.H{"last_login_data": data, "last_login_data2": data2}})
	}
}
