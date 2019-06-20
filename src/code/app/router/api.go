package router

import (
	v100User "code/controller/v100/user"
	v110Test "code/controller/v110/test"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Api = map[string]map[string]interface{}{
	"/test/index/isok":      {"version": "1.1.0+", "method": map[string]byte{http.MethodGet: 0}},
	"/test/index/main":      {"version": "1.1.0+", "method": map[string]byte{http.MethodGet: 0}},
	"/test/index/allconfig": {"version": "1.1.0+", "method": map[string]byte{http.MethodGet: 0}},

	"/user/info/base":     {"version": "1.0.0+", "method": map[string]byte{http.MethodGet: 0}},
	"/user/info/loginlog": {"version": "1.0.0+", "method": map[string]byte{http.MethodGet: 0}},
}

func VersionGroup(e *gin.Engine, hf ...gin.HandlerFunc) {
	v100 := e.Group("/v100", hf...)
	{
		v100.GET("/user/info/base", v100User.Info(e).BaseInfo)
		v100.GET("/user/info/loginlog", v100User.Info(e).LoginLog)
	}
	v110 := e.Group("/v110", hf...)
	{
		v110.GET("/test/index/isok", v110Test.Index(e).CheckOk)
		v110.GET("/test/index/main", v110Test.Index(e).Main)
		v110.GET("/test/index/allconfig", v110Test.Index(e).AllConfig)
	}
}
