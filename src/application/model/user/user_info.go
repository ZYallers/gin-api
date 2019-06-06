package model

import (
	"application/app/abst"
	"github.com/gin-gonic/gin"
	"time"
)

type UserInfoTable struct {
	Id                  int       `json:"id"`
	UserId              int       `json:"user_id"`
	HeadImg             string    `json:"head_img"`
	Nickname            string    `json:"nickname"`
	Descr               string    `json:"descr"`
	Mobile              string    `json:"mobile"`
	Sex                 int       `json:"sex"`
	Birthday            string    `json:"birthday"`
	BodyHigh            string    `json:"body_high"`
	Province            string    `json:"province"`
	City                string    `json:"city"`
	Vocation            string    `json:"vocation"`
	ContinuousFoodDays  int       `json:"continuous_food_days"`
	ContinuousSportDays int       `json:"continuous_sport_days"`
	ContinuousWeight    int       `json:"continuous_weight"`
	TotalOutCalory      int       `json:"total_out_calory"`
	FirstWeight         string    `json:"first_weight"`
	StartWeight         string    `json:"start_weight"`
	DescWeight          string    `json:"desc_weight"`
	InvitedCode         string    `json:"invited_code"`
	IsShield            string    `json:"is_shield"`
	ScaleName           string    `json:"scale_name"`
	UpdateTime          time.Time `json:"update_time"`
	CreateTime          time.Time `json:"create_time"`
	V                   int       `json:"v"`
	ShakeName           string    `json:"shake_name"`
	ShakeDescr          string    `json:"shake_descr"`
	Platform            int       `json:"platform"`
	AdviserMasterId     int       `json:"adviser_master_id"`
	TreadmillName       string    `json:"treadmill_name"`
	TreadmillMac        string    `json:"treadmill_mac"`
	RelateMobile        string    `json:"relate_mobile"`
}

type userInfo struct {
	abst.Model
}

func NewUserInfo(c *gin.Context) *userInfo {
	model := new(userInfo)
	model.TableName = "et_user_info"
	model.Ctx = c
	return model
}

func (this *userInfo) FetchOneByUserId(userId int) UserInfoTable {
	this.InitEnjoythin()
	var userInfo UserInfoTable
	this.Enjoythin.Table(this.TableName).Where("user_id = ?", userId).Limit(1).Find(&userInfo)
	return userInfo
}
