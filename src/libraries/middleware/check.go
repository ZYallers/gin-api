package middleware

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"src/config/env"
	"src/libraries/helper"
	"strconv"
	"time"
)

// app签名验证
func SignCheckHandler(ctx *gin.Context) {
	var backup io.ReadCloser
	backup, ctx.Request.Body, _ = helper.DrainBody(ctx.Request.Body)
	defer func() { ctx.Request.Body = backup }()
	if !signCheck(ctx) {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusForbidden, "msg": "signature error"})
		return
	}
}

// app登录验证
func LoginCheckHandler(ctx *gin.Context) {
	var backup io.ReadCloser
	backup, ctx.Request.Body, _ = helper.DrainBody(ctx.Request.Body)
	defer func() { ctx.Request.Body = backup }()
	if !loginCheck(ctx) {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "msg": "please login first"})
		return
	}
}

// SignCheck
func signCheck(ctx *gin.Context) (pass bool) {
	sign := helper.QueryPostForm(ctx, "sign")
	if sign == "" {
		return
	}

	// 开发测试模式下，用固定sign判断
	if env.App.IsDevMode && sign == env.App.DevSign {
		return true
	}

	timestampStr := helper.QueryPostForm(ctx, "utime")
	if timestampStr == "" {
		return
	}

	timestamp, err := strconv.ParseInt(timestampStr, 10, 0)
	if err != nil {
		return
	}

	if time.Now().Unix()-timestamp > env.App.SignTimeExpiration {
		return
	}

	hash := md5.New()
	hash.Write([]byte(timestampStr + env.App.TokenKey))
	md5str := hex.EncodeToString(hash.Sum(nil))
	if sign == base64.StdEncoding.EncodeToString([]byte(md5str)) {
		pass = true
	}

	return
}

// loginCheck
func loginCheck(ctx *gin.Context) bool {
	if vars := sessionData(helper.QueryPostForm(ctx, env.Session.TokenKey)); vars != nil {
		return true
	}
	return false
}
