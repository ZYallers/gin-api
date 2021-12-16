package test

import (
	"fmt"
	"net/http"
	"src/config/env"
	"src/libraries/core"
	"src/libraries/handler"
	"src/libraries/helper"
	"time"

	"github.com/gin-gonic/gin"
)

type Index struct {
	core.Controller
}

func (i *Index) Init() core.MtdCfg {
	i.Config = core.MtdCfg{
		"Isok":        {Http: []string{http.MethodGet, http.MethodPost}, Func: i.Isok, Rest: "isok"},
		"SwaggerDocs": {Http: []string{http.MethodGet}, Func: i.SwaggerDocs, Rest: "swag/docs"},
		"BaseCheck":   {Http: []string{"GET", "POST"}, Func: i.BaseCheck, Rest: "base/check"},
	}
	return i.Config
}

// @title /test/index/isok
// @Summary 检测服务是否正常
// @Tags test,cloud
// @Accept json
// @Produce json
// @Router /test/isok [get]
// @Author cloud 2020/9/7 下午4:25
// @Update cloud 2020/9/7 下午4:25
// @Success 200 {string} json "{"code":200,"msg":"ok", "data":{}}"
func (i *Index) Isok(ctx *gin.Context) {
	reqStr, _ := ctx.Get(env.ReqStrKey)
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "OK",
		"data": gin.H{
			"mode":      gin.Mode(),
			"system_ip": helper.SystemIP(),
			"request":   reqStr,
		},
	})
}

// @title /test/swag/docs
// @Summary 获取微服务swagger.json
// @Tags test,cloud
// @Accept json
// @Produce json
// @Router /test/swag/docs [get]
// @Author cloud 2020/9/7 下午4:28
// @Update cloud 2020/9/7 下午4:28
// @Param m query string true "微服务"
// @Success 200 {string} json "{"code":200,"msg":"ok", "data":{}}"
func (i *Index) SwaggerDocs(ctx *gin.Context) {
	module := ctx.DefaultQuery("m", "")
	if module == "" {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": `the required parameter "m" is missing`})
		return
	}
	if module == "api" {
		handler.SwagDocHandler(ctx)
		return
	}
	i.ServiceRewrite(ctx, fmt.Sprintf("http://%s.hxsapp.com/swag/json", module), 3*time.Second)
}

func (i *Index) BaseCheck(ctx *gin.Context) {
	i.RPCXService(ctx, "rpcx-base", "test/check")
}
