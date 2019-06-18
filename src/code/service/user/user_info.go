package service

import (
	"code/app/abst"
	"code/app/cons"
	"code/model/user"
	"encoding/json"
	"strconv"
	"time"
)

type userInfo struct {
	abst.Service
}

func NewUserInfo() *userInfo {
	return &userInfo{}
}

func (ui *userInfo) GetUserInfoByUserId(userId int) model.UserInfoTable {
	cache := ui.GetCache()
	var uit model.UserInfoTable
	key := cons.UserInfoStringRdsKey + ":" + strconv.Itoa(userId)

	str, _ := cache.Get(key).Result()
	if str == "" {
		uit = model.NewUserInfo().FetchOneByUserId(userId)
		if uit.Id > 0 {
			if bytes, err := json.Marshal(uit); err == nil {
				cache.Set(key, string(bytes), 300*time.Second)
			}
		} else {
			// 设置个短暂的缓存
			cache.Set(key, "", 10*time.Second)
		}
	} else {
		_ = json.Unmarshal([]byte(str), &uit)
	}
	return uit
}
