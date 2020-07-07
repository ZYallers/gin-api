package live

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type YunxinController struct {
	abs.Controller
}

const liveUri = "http://live.hxsapp.com/live/"
//const liveUri = "http://localhost:9140/live/"

func Yunxin() YunxinController {
	c := YunxinController{}
	return c
}

func (c YunxinController) IsAnchor(ctx *gin.Context) {
	c.ServiceRewrite(ctx, liveUri+tool.CurrentMethodName())
}

func (c YunxinController) GetChanAddr(ctx *gin.Context) {
	c.ServiceRewrite(ctx, liveUri+tool.CurrentMethodName())
}

func (c YunxinController) GetRoomListByType(ctx *gin.Context) {
	c.ServiceRewrite(ctx, liveUri+tool.CurrentMethodName())
}

func (c YunxinController) SetChanStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, liveUri+tool.CurrentMethodName())
}

func (c YunxinController) CountAnchorLikeNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, liveUri+tool.CurrentMethodName())
}

func (c YunxinController) Chstatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, liveUri+tool.CurrentMethodName())
}

func (c YunxinController) Mute(ctx *gin.Context) {
	c.ServiceRewrite(ctx, liveUri+tool.CurrentMethodName())
}