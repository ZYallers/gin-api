package live

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type PopController struct {
	abs.Controller
}

const PopUri = "http://live.hxsapp.com/pop/"

func Pop() PopController {
	p := PopController{}
	return p
}

func (p PopController) GetRoomAwardList(ctx *gin.Context) {
	p.ServiceRewrite(ctx, PopUri+tool.CurrentMethodName())
}

func (p PopController) GetAwardById(ctx *gin.Context) {
	p.ServiceRewrite(ctx, PopUri+tool.CurrentMethodName())
}

func (p PopController) ReceiveBean(ctx *gin.Context) {
	p.ServiceRewrite(ctx, PopUri+tool.CurrentMethodName())
}
