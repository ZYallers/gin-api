package abs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"src/library/tool"
	"strings"
)

type Module struct {
}

// 通过反射自动为http请求方式注册Method
func (m Module) BindMethodOfController(rg *gin.RouterGroup, controllers ...interface{}) {
	if len(controllers) == 0 {
		return
	}
	for _, controller := range controllers {
		typeOf := reflect.TypeOf(controller)
		valueOf := reflect.ValueOf(controller)
		controllerConfig := valueOf.FieldByName("Config").Interface().(map[string]MethodConfig)
		controllerName := strings.Replace(typeOf.Name(), "Controller", "", 1)
		for i := 0; i < typeOf.NumMethod(); i++ {
			valueOfMethod := valueOf.Method(i)
			if valueOfMethod.Type().String() != "func(*gin.Context)" {
				continue
			}
			typeOfMethod := typeOf.Method(i)
			controllerName = tool.StrFirstToLower(controllerName)
			methodName := tool.StrFirstToLower(typeOfMethod.Name)
			httpMethods := []string{http.MethodGet, http.MethodPost}
			var relativePath string
			if methodConfig, exist := controllerConfig[typeOfMethod.Name]; exist {
				if methodConfig.Rest != "" {
					relativePath = methodConfig.Rest
				} else {
					if methodConfig.ControllerNameFirstUpper {
						controllerName = tool.StrFirstToUpper(controllerName)
					}
					if methodConfig.MethodNameFirstUpper {
						methodName = typeOfMethod.Name
					}
				}
				if len(methodConfig.HttpMethods) > 0 {
					httpMethods = methodConfig.HttpMethods
				}
			}
			if relativePath == "" {
				relativePath = controllerName + "/" + methodName
			}
			for _, method := range httpMethods {
				reflect.ValueOf(rg).MethodByName(method).Call([]reflect.Value{reflect.ValueOf("/" + relativePath), valueOfMethod})
			}
		}
	}
}
