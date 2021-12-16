package logger

import (
	"fmt"
	"github.com/ZYallers/zgin/libraries/tool"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"regexp"
	"src/config/app"
	"src/config/env"
	"strings"
	"time"
)

type logger struct {
	Service string
	Ctx     *gin.Context
	handler func() *zap.Logger
}

var log *logger

func NewRPCXLogger() *logger {
	if log == nil {
		log = &logger{handler: func() *zap.Logger {
			return Use("rpcx-service")
		}}
	}
	return log
}

func SetLoggerAttr(service string, ctx *gin.Context) {
	if log == nil {
		return
	}
	log.Service = service
	log.Ctx = ctx
}

func (l *logger) Debug(v ...interface{}) {
	l.handler().Debug(fmt.Sprint(v...))
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.handler().Debug(fmt.Sprintf(format, v...))
}

func (l *logger) Info(v ...interface{}) {
	l.handler().Info(fmt.Sprint(v...))
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.handler().Info(fmt.Sprintf(format, v...))
}

func (l *logger) Warn(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.handler().Warn(s)
	l.sendMessage("Warn: "+s, false)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.handler().Warn(s)
	l.sendMessage("Warn: "+s, false)
}

func (l *logger) Error(v ...interface{}) {
	s := fmt.Sprint(v...)
	if strings.Contains(s, "i/o timeout") {
		return
	}
	l.handler().Error(s)
	l.sendMessage("Error: "+s, true)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.handler().Error(s)
	l.sendMessage("Error: "+s, true)
}

func (l *logger) Fatal(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.handler().Fatal(s)
	l.sendMessage("Fatal: "+s, true)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.handler().Fatal(s)
	l.sendMessage("Fatal: "+s, true)
}

func (l *logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.handler().Panic(s)
	l.sendMessage("Panic: "+s, true)
}

func (l *logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.handler().Panic(s)
	l.sendMessage("Panic: "+s, true)
}

func (l *logger) Handle(v ...interface{}) {
	l.Error(v...)
}

func (l *logger) sendMessage(msg string, isAtAll bool) {
	go func(msg string, isAtAll bool) {
		defer func() { recover() }()
		hostname, _ := os.Hostname()
		text := []string{
			msg + "\n---------------------------",
			"App: " + app.Name,
			"Env: " + env.App.Env,
			"Port: " + *app.HttpServerAddr,
			"Service: " + l.Service,
			"HostName: " + hostname,
			"Time: " + time.Now().Format("2006/01/02 15:04:05.000"),
			"SystemIP: " + tool.SystemIP(),
			"PublicIP: " + tool.PublicIP(),
		}
		if l.Ctx != nil {
			text = append(text, "ClientIP: "+tool.ClientIP(l.Ctx.ClientIP()))
			text = append(text, "URL: "+"https://"+l.Ctx.Request.Host+l.Ctx.Request.URL.String())
			// 去除连续的换行符
			re, _ := regexp.Compile(`\s{2,}`)
			if reqStr := l.Ctx.GetString(env.ReqStrKey); reqStr != "" {
				text = append(text, "\nRequest:\n"+strings.TrimSpace(re.ReplaceAllString(reqStr, "\n")))
			}
		}
		if gin.IsDebugging() {
			isAtAll = false
		}
		postData := map[string]interface{}{
			"msgtype": "text",
			"text":    map[string]string{"content": strings.Join(text, "\n") + "\n"},
			"at":      map[string]interface{}{"isAtAll": isAtAll},
		}
		jsonHeaders := map[string]string{"Content-Type": "application/json"}
		_, _ = tool.NewRequest(env.SimpleRobot).SetHeaders(jsonHeaders).SetPostData(postData).Post()
	}(msg, isAtAll)
}
