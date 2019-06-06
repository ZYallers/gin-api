package router

import (
	v100User "application/controller/http/v100/user"
	v110Test "application/controller/http/v110/test"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Api = map[string]map[string]interface{}{
	"/test/index/allconfig": {"version": "1.0.0|1.1.0+", "method": map[string]byte{http.MethodGet: 0,},},
	"/user/info/base":       {"version": "1.0.0+", "method": map[string]byte{http.MethodGet: 0,},},
	"/user/info/loginlog":   {"version": "1.0.0+", "method": map[string]byte{http.MethodGet: 0,},},
}

func VersionGroup(engine *gin.Engine, handlerFunc ...gin.HandlerFunc) {
	v100 := engine.Group("/v100", handlerFunc...)
	{
		v100.GET("/user/info/base", v100User.NewUserInfo(engine).BaseInfo)
		v100.GET("/user/info/loginlog", v100User.NewUserInfo(engine).LoginLog)
	}
	v110 := engine.Group("/v110", handlerFunc...)
	{
		v110.GET("/test/index/allconfig", v110Test.NewTestIndex(engine).AllConfig)
	}
}
