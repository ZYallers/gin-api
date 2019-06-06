package middleware

import (
	"application/app/constant"
	"application/app/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

func VersionControl(engine *gin.Engine, apis map[string]map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if regexp.MustCompile(`^\/v[0-9]{3,6}\/.*$`).MatchString(path) {
			return
		}

		api, ok := apis[path]
		if !ok {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "Api not found"})
			return
		}

		method := c.Request.Method
		if _, ok := api["method"].(map[string]byte)[method]; !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden, "msg": "Api method not allowed"})
			return
		}

		var version string
		if method == http.MethodPost {
			version = c.DefaultPostForm("version", constant.Version)
		} else {
			version = c.DefaultQuery("version", constant.Version)
		}

		flag := false
		// 先获取 Router 支持的版本，然后遍历
		for _, item := range strings.Split(api["version"].(string), "|") {
			// 判断是否包含'+'支持以上版本
			if ilen := len(item); item[ilen-1:] == "+" {
				// 判断 version 是否大于等于要求的版本
				vs := item[0 : ilen-1]
				if tool.VersionCompare(version, vs, ">=") {
					version = vs
					flag = true
					break
				}
			} else {
				// 判断 version 是否等于要求的版本
				if tool.VersionCompare(version, item, "=") {
					flag = true
					break
				}
			}
		}
		if !flag {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "Api version not found"})
			return
		}

		c.Request.URL.Path = "/v" + strings.Join(strings.Split(version, "."), "") + path
		engine.HandleContext(c)
		c.Abort()
		return
	}
}
