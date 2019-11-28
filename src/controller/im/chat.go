package im

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ChatController struct {
	abs.Controller
}

const chatUri = "http://im.hxsapp.com/im/Chat/"

func Chat() ChatController {
	c := ChatController{}
	c.Config = map[string]abs.MethodConfig{
		"GetCustomerGroups":                {ControllerNameFirstUpper: true},
		"ShowGroupCustomers":               {ControllerNameFirstUpper: true},
		"AddGroup":                         {ControllerNameFirstUpper: true},
		"EditGroupName":                    {ControllerNameFirstUpper: true},
		"DeleteCustomerGroup":              {ControllerNameFirstUpper: true},
		"MoveCustomerGroup":                {ControllerNameFirstUpper: true},
		"AssignCustomerToGroupByFirstChat": {ControllerNameFirstUpper: true},
		"GetCustomerNameAndGroup":          {ControllerNameFirstUpper: true},
		"GetCustomerInfo":                  {ControllerNameFirstUpper: true},
		"EditRemarkOrImpression":           {ControllerNameFirstUpper: true},
		"GetAdviserConfig":                 {ControllerNameFirstUpper: true},
		"SetSendShortcuts":                 {ControllerNameFirstUpper: true},
		"GetUserInfo":                      {ControllerNameFirstUpper: true},
		"GetCustomerTop":                   {ControllerNameFirstUpper: true},
		"OperateCustomerTop":               {ControllerNameFirstUpper: true},
		"GetCustomerStatusAndExclusive":    {ControllerNameFirstUpper: true},
		"OnlinePayData":                    {ControllerNameFirstUpper: true},
		"OrderListData":                    {ControllerNameFirstUpper: true},
		"CheckCanChat":                     {ControllerNameFirstUpper: true},
		"DiaryListData":                    {ControllerNameFirstUpper: true},
		"DiaryStateListData":               {ControllerNameFirstUpper: true},
		"GetUserInfoByEncodePhone":         {ControllerNameFirstUpper: true},
		"GetCustomerAllPhones":             {ControllerNameFirstUpper: true},
		"DataStatistics":                   {ControllerNameFirstUpper: true},
		"ChatTimeLimit":                    {ControllerNameFirstUpper: true},
		"AddGroupChat":                     {ControllerNameFirstUpper: true},
		"UpdateGroupChat":                  {ControllerNameFirstUpper: true},
		"DeleteGroupChat":                  {ControllerNameFirstUpper: true},
		"GetGroupChatInfo":                 {ControllerNameFirstUpper: true},
		"EnterGroupChat":                   {ControllerNameFirstUpper: true},
		"KickOffGroupChat":                 {ControllerNameFirstUpper: true},
		"ActiveAddToGroupChat":             {ControllerNameFirstUpper: true},
		"GetGroupMemberList":               {ControllerNameFirstUpper: true},
		"GetAdviserList":                   {ControllerNameFirstUpper: true},
		"IsExistGroup":                     {ControllerNameFirstUpper: true},
		"YunXinChat":                       {ControllerNameFirstUpper: true},
	}
	return c
}

func (c ChatController) GetCustomerGroups(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) ShowGroupCustomers(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) AddGroup(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) EditGroupName(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) DeleteCustomerGroup(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) MoveCustomerGroup(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) AssignCustomerToGroupByFirstChat(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetCustomerNameAndGroup(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetCustomerInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) EditRemarkOrImpression(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetAdviserConfig(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) SetSendShortcuts(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetUserInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetCustomerTop(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) OperateCustomerTop(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetCustomerStatusAndExclusive(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) OnlinePayData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) OrderListData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) CheckCanChat(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) DiaryListData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) DiaryStateListData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetUserInfoByEncodePhone(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetCustomerAllPhones(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) DataStatistics(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) ChatTimeLimit(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) AddGroupChat(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) UpdateGroupChat(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) DeleteGroupChat(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetGroupChatInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) EnterGroupChat(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) KickOffGroupChat(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) ActiveAddToGroupChat(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetGroupMemberList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

func (c ChatController) GetAdviserList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @DateTime:    2019-06-25 10:19:21
 * @Description: 判断该用户是否在群中
 */
func (c ChatController) IsExistGroup(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}

/**
 * @Author:      caius
 * @role:        role1,role2
 * @DateTime:    2019-07-29 11:10:52
 * @Description: 网易 语音聊天 这逻辑应在 canImWithChatAccount 之后
 */
func (c ChatController) YunXinChat(ctx *gin.Context) {
	c.ServiceRewrite(ctx, chatUri+tool.CurrentMethodName())
}
