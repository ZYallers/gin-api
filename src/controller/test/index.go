package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"src/abs"
	app "src/config"
	"src/library/tool"
)

type IndexController struct {
	abs.Controller
}

func Index() IndexController {
	i := IndexController{}
	i.Config = map[string]abs.MethodConfig{
		"Isok": {HttpMethods: []string{http.MethodGet, http.MethodPost}},
	}
	return i
}

func (i IndexController) Isok(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": gin.H{
			"mode":        gin.Mode(),
			"debug_stack": app.DebugStack,
			"client_ip":   ctx.ClientIP(),
			"public_ip":   tool.PublicIP(),
			"system_ip":   tool.SystemIP(),
			"keys":        ctx.Keys,
		},
	})
}

/*func (i IndexController) Multi(ctx *gin.Context) {
	//i.ServiceMultiRewrite(ctx, "http://account.hxsapp.com/user/userInfo/getUserInfo", "user_id", 10)
	//i.ServiceMultiRewrite(ctx, "http://im.hxsapp.com/api/Brm/getUserInfoByOpenImAccount", "openim_account", 20)
}*/
