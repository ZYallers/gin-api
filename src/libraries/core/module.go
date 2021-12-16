package core

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"src/libraries/helper"
	"src/libraries/middleware"
	"strings"
)

type IModule interface {
	Group(eg *gin.Engine)
}

type Module struct{}

func (m *Module) GetControllerName(i IController) string {
	str := fmt.Sprintf("%T", i)
	return str[strings.Index(str, ".")+1:]
}

func (m *Module) BindMethodOfController(eg *gin.Engine, moduleName string, controllers ...IController) {
	for _, controller := range controllers {
		var rp []string

		for mtd, mtdCfg := range controller.Init() {
			if mtdCfg.Func == nil {
				continue
			}

			httpMtd := mtdCfg.Http
			if len(httpMtd) == 0 {
				httpMtd = []string{http.MethodGet, http.MethodPost}
			}

			var relativePath bytes.Buffer
			relativePath.WriteString("/" + moduleName)
			if mtdCfg.Rest == "" {
				controllerName := m.GetControllerName(controller)
				if mtdCfg.CNCap {
					relativePath.WriteString("/" + helper.StrFirstToUpper(controllerName))
				} else {
					relativePath.WriteString("/" + helper.StrFirstToLower(controllerName))
				}
				if mtdCfg.MNCap {
					relativePath.WriteString("/" + helper.StrFirstToUpper(mtd))
				} else {
					relativePath.WriteString("/" + helper.StrFirstToLower(mtd))
				}
			} else {
				relativePath.WriteString("/" + strings.TrimLeft(mtdCfg.Rest, "/"))
			}

			apiPath := relativePath.String()
			rp = append(rp, apiPath)

			var funcs []gin.HandlerFunc

			if mtdCfg.SignCheck {
				funcs = append(funcs, middleware.SignCheckHandler)
			}
			if mtdCfg.LoginCheck {
				funcs = append(funcs, middleware.LoginCheckHandler)
			}
			/* TODO::等updateSessionExpireApiPath接口迁移到Go之后再放开
			if apiPath == updateSessionExpireApiPath {
				funcs = append(funcs, middleware.UpdateSessionExpireHandler)
			}*/
			funcs = append(funcs, mtdCfg.Func)

			for _, method := range httpMtd {
				switch method {
				case http.MethodGet:
					eg.GET(apiPath, funcs...)
				case http.MethodPost:
					eg.POST(apiPath, funcs...)
				case http.MethodPut:
					eg.PUT(apiPath, funcs...)
				case http.MethodPatch:
					eg.PATCH(apiPath, funcs...)
				case http.MethodHead:
					eg.HEAD(apiPath, funcs...)
				case http.MethodOptions:
					eg.OPTIONS(apiPath, funcs...)
				case http.MethodDelete:
					eg.DELETE(apiPath, funcs...)
				}
			}
		}
	}
}
