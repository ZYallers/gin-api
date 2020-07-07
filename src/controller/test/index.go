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
		"Isok":   {HttpMethods: []string{http.MethodGet, http.MethodPost}},
		"MyRest": {HttpMethods: []string{http.MethodGet}, Rest: "my/rest"},
	}
	return i
}

func (i IndexController) Isok(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
		"data": gin.H{
			"mode":        gin.Mode(),
			"debug_stack": app.DebugStack,
			"system_ip":   tool.SystemIP(),
			"client_ip":   tool.ClientIP(ctx.ClientIP()),
			"public_ip":   tool.PublicIP(),
			"keys":        ctx.Keys,
		},
	})
}

func (i IndexController) MyRest(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"path": ctx.Request.URL})
}

/*func (i IndexController) Multi(ctx *gin.Context) {
	//i.ServiceMultiRewrite(ctx, "http://account.hxsapp.com/user/userInfo/getUserInfo", "user_id", 10)
	//i.ServiceMultiRewrite(ctx, "http://im.hxsapp.com/api/Brm/getUserInfoByOpenImAccount", "openim_account", 20)
}*/

func (i IndexController) Demo(ctx *gin.Context) {
	i.ServiceRewrite(ctx, "http://base.hxsapp.com/base/common/testError")
}
