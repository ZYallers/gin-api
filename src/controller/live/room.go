package live

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type RoomController struct {
	abs.Controller
}

const roomUri = "http://live.hxsapp.com/room/"
//const roomUri = "http://127.0.0.1:9140/room/"

func Room() RoomController {
	c := RoomController{}
	return c
}

func (c RoomController) WatchNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, roomUri+tool.CurrentMethodName())
}

func (c RoomController) GetRole(ctx *gin.Context) {
	c.ServiceRewrite(ctx, roomUri+tool.CurrentMethodName())
}

func (c RoomController) GetConf(ctx *gin.Context) {
	c.ServiceRewrite(ctx, roomUri+tool.CurrentMethodName())
}
