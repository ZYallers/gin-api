package helper

import (
	"github.com/gin-gonic/gin"
	"os"
	"regexp"
	"src/config/app"
	"src/config/env"
	"strings"
	"time"
)

func PushSimpleMessage(msg string, isAtAll bool) {
	hostname, _ := os.Hostname()
	text := []string{
		msg + "\n---------------------------",
		"App: " + app.Name,
		"Env: " + env.App.Env,
		"Port: " + *app.HttpServerAddr,
		"HostName: " + hostname,
		"SystemIP: " + SystemIP(),
		"PublicIP: " + PublicIP(),
		"Time: " + time.Now().Format("2006/01/02 15:04:05.000"),
	}
	if gin.IsDebugging() {
		isAtAll = false
	}
	postData := map[string]interface{}{
		"msgtype": "text",
		"text":    map[string]string{"content": strings.Join(text, "\n") + "\n"},
		"at":      map[string]interface{}{"isAtAll": isAtAll},
	}
	jsonHeaders := map[string]string{"Content-Type": "application/json;charset=utf-8"}
	_, _ = NewRequest(env.SimpleRobot).SetHeaders(jsonHeaders).SetPostData(postData).Post()
}

func PushContextMessage(ctx *gin.Context, msg string, reqStr string, stack string, isAtAll bool) {
	hostname, _ := os.Hostname()
	text := []string{
		msg + "\n---------------------------",
		"App: " + app.Name,
		"Env: " + env.App.Env,
		"Port: " + *app.HttpServerAddr,
		"HostName: " + hostname,
		"SystemIP: " + SystemIP(),
		"PublicIP: " + PublicIP(),
		"ClientIP: " + ClientIP(ctx.ClientIP()),
		"Time: " + time.Now().Format("2006/01/02 15:04:05.000"),
		"URL: " + "https://" + ctx.Request.Host + ctx.Request.URL.String(),
	}
	// 去除连续的换行符
	re, _ := regexp.Compile(`\s{2,}`)
	if reqStr != "" {
		text = append(text, "\nRequest:\n"+strings.TrimSpace(re.ReplaceAllString(reqStr, "\n")))
	}
	if stack != "" {
		text = append(text, "\nStack:\n"+strings.TrimSpace(re.ReplaceAllString(stack, "\n")))
	}
	if gin.IsDebugging() {
		isAtAll = false
	}
	postData := map[string]interface{}{
		"msgtype": "text",
		"text":    map[string]string{"content": strings.Join(text, "\n") + "\n"},
		"at":      map[string]interface{}{"isAtAll": isAtAll},
	}
	jsonHeaders := map[string]string{"Content-Type": "application/json;charset=utf-8"}
	_, _ = NewRequest(env.ContextRobot).SetHeaders(jsonHeaders).SetPostData(postData).Post()
}
