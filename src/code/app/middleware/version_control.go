package middleware

import (
	"code/app/cons"
	"code/app/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

func VersionControl(engine *gin.Engine, apis map[string]map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !regexp.MustCompile(`^\/v[0-9]{3,6}\/.*$`).MatchString(c.Request.URL.Path) {
			if api, ok := apis[c.Request.URL.Path]; !ok {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "Api not found"})
			} else {
				if _, ok := api["method"].(map[string]byte)[c.Request.Method]; !ok {
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden, "msg": "Api method not allowed"})
				} else {
					var version string
					if c.Request.Method == http.MethodPost {
						version = c.DefaultPostForm("version", cons.Version)
					} else {
						version = c.DefaultQuery("version", cons.Version)
					}
					var flag bool
					for _, item := range strings.Split(api["version"].(string), "|") {
						if ilene := len(item); item[ilene-1:] == "+" {
							vs := item[0 : ilene-1]
							if tool.VersionCompare(version, vs, ">=") {
								version = vs
								flag = true
								break
							}
						} else {
							if tool.VersionCompare(version, item, "=") {
								flag = true
								break
							}
						}
					}
					if !flag {
						c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "Api version not found"})
					} else {
						c.Request.URL.Path = "/v" + strings.Join(strings.Split(version, "."), "") + c.Request.URL.Path
						engine.HandleContext(c)
						c.Abort()
					}
				}
			}
		}
	}
}
