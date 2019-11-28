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
		config := valueOf.FieldByName("Config").Interface().(map[string]MethodConfig)
		controllerName := strings.Replace(typeOf.Name(), "Controller", "", 1)
		for i := 0; i < typeOf.NumMethod(); i++ {
			valueOfMethod := valueOf.Method(i)
			if valueOfMethod.Type().String() != "func(*gin.Context)" {
				continue
			}
			typeOfMethod := typeOf.Method(i)
			cname := tool.StrFirstToLower(controllerName)
			mame := tool.StrFirstToLower(typeOfMethod.Name)
			hms := []string{http.MethodGet, http.MethodPost}
			if methodConfig, exist := config[typeOfMethod.Name]; exist {
				if methodConfig.ControllerNameFirstUpper {
					cname = tool.StrFirstToUpper(cname)
				}
				if methodConfig.MethodNameFirstUpper {
					mame = typeOfMethod.Name
				}
				if len(methodConfig.HttpMethods) > 0 {
					hms = methodConfig.HttpMethods
				}
			}
			for _, m := range hms {
				reflect.ValueOf(rg).MethodByName(m).Call([]reflect.Value{reflect.ValueOf("/" + cname + "/" + mame), valueOfMethod})
			}
		}
	}
}
