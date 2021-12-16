package helper

import (
	"os"
	"runtime"
	"strings"
)

func CurrentMethodName() string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	return StrFirstToLower(name[strings.LastIndex(name, `.`)+1:])
}

func CurrentFileName() string {
	_, path, _, _ := runtime.Caller(1)
	file := path[strings.LastIndex(path, `/`)+1:]
	return file[0:strings.Index(file, `.`)]
}

// 检查文件是否存在
func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
