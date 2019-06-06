package service

import (
	"application/app/abst"
	"application/app/constant"
	"application/model/user"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type userInfo struct {
	abst.Service
}

func NewUserInfo(c *gin.Context) *userInfo {
	userInfo := new(userInfo)
	userInfo.Ctx = c
	return userInfo
}

func (this *userInfo) GetUserInfoByUserId(userId int) model.UserInfoTable {
	this.InitCache()
	var userInfo model.UserInfoTable
	key := constant.UserInfoStringRdsKey + ":" + strconv.Itoa(userId)

	str, _ := this.Cache.Get(key).Result()
	if str == "" {
		userInfo = model.NewUserInfo(this.Ctx).FetchOneByUserId(userId)
		if userInfo.Id > 0 {
			if bytes, err := json.Marshal(userInfo); err == nil {
				this.Cache.Set(key, string(bytes), 300*time.Second)
			}
		} else {
			// 设置个短暂的缓存
			this.Cache.Set(key, "", 10*time.Second)
		}
	} else {
		json.Unmarshal([]byte(str), &userInfo)
	}
	return userInfo
}
