package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"src/libraries/helper"
)

func SwagDocHandler(ctx *gin.Context) {
	// 生产环境不注册
	if !gin.IsDebugging() {
		return
	}

	pwd, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/doc/swagger.json", pwd)
	if !helper.FileIsExist(filePath) {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusNotFound, "msg": "swagger.json file not exist"})
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}

	defer f.Close()

	fd, err := helper.IoCopy(f)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}

	ctx.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
	_, _ = ctx.Writer.Write(fd)
	ctx.AbortWithStatus(http.StatusOK)
}
