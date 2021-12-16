package helper

import (
	"fmt"
	"go.uber.org/zap"
	"reflect"
	"runtime/debug"
	"src/libraries/logger"
)

func SafeDefer(params ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("%+v", r)
			stack := string(debug.Stack())
			logger.Use("panic").Error(msg, zap.String("debug_stack", stack))
			PushSimpleMessage(fmt.Sprintf("recovery from panic:\n%s\n%s", msg, stack), true)
			return
		}
	}()

	r := recover()
	if r == nil {
		return
	}

	err := fmt.Errorf("%+v", r)
	stack := string(debug.Stack())
	logger.Use("panic").Error(err.Error(), zap.String("debug_stack", stack))
	PushSimpleMessage(fmt.Sprintf("recovery from panic:\n%s\n%s", err.Error(), stack), true)

	if paramsLen := len(params); paramsLen > 0 {
		if reflect.TypeOf(params[0]).String()[0:4] != "func" {
			return
		}
		var args []reflect.Value
		if paramsLen > 1 {
			args = append(args, reflect.ValueOf(err))
			for _, v := range params[1:] {
				args = append(args, reflect.ValueOf(v))
			}
		}
		reflect.ValueOf(params[0]).Call(args)
	}
}
