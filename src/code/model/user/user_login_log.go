package model

import (
	"code/app/abst"
	"time"
)

type UserLoginLogTable struct {
	Id            int       `json:"id"`
	UserId        int       `json:"user_id"`
	ModelVersion  string    `json:"model_version"`
	SystemVersion string    `json:"system_version"`
	AppVersion    string    `json:"app_version"`
	UpdateTime    time.Time `json:"update_time"`
	LastLogIp     string    `json:"last_log_ip"`
	ModelIdfa     string    `json:"model_idfa"`
	Platform      int       `json:"platform"`
	SessToken     string    `json:"sess_token"`
}

type userLoginLog struct {
	abst.Model
}

func NewUserLoginLog() *userLoginLog {
	ull := new(userLoginLog)
	ull.TableName = "et_user_login_log"
	return ull
}

func (ull *userLoginLog) FetchByUserIdAndLimit(userId int, fields string, limit uint8, orderBy string) []UserLoginLogTable {
	if userId == 0 {
		return nil
	}
	var rows []UserLoginLogTable
	pdo := ull.GetEnjoythin().Table(ull.TableName)
	if fields != "" {
		pdo = pdo.Select(fields)
	}
	pdo = pdo.Where("user_id = ?", userId)
	if orderBy != "" {
		pdo = pdo.Order(orderBy)
	}
	if limit == 0 {
		limit = 10
	}
	pdo.Limit(limit).Find(&rows)
	return rows
}
