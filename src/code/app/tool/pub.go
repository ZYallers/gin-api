package tool

import (
	"code/app/cons"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CurrentMethodName() string {
	pc, _, _, _ := runtime.Caller(1)
	fnSlice := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	return StrFirstToLower(fnSlice[len(fnSlice)-1])
}

func CurrentFileName() string {
	_, file, _, _ := runtime.Caller(1)
	fnSlice := strings.Split(file, "/")
	return strings.Split(fnSlice[len(fnSlice)-1], ".")[0]
}

/**
 * 字符串首字母转成大写
 */
func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArray := []rune(str)
	if strArray[0] >= 97 && strArray[0] <= 122 {
		strArray[0] -= 32
	}
	return string(strArray)
}

/**
 * 字符串首字母转成小写
 */
func StrFirstToLower(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] < 97 || strArry[0] > 122 {
		strArry[0] += 32
	}
	return string(strArry)
}

func SendDingTalkGroupMessage(c *gin.Context, msg string) {
	if !cons.DingTalkMsgEnable {
		return
	}
	scheme := "http://"
	if c.Request.TLS != nil {
		scheme = "https://"
	}
	host, _ := os.Hostname()
	text := []string{
		msg + "\n---------------",
		"AppName: " + cons.Name,
		"HostName: " + host,
		"PublicIP: " + PublicIP(),
		"SystemIP: " + SystemIP(),
		"ClientIP: " + c.ClientIP(),
		"HttpAddr: " + cons.HttpServerAddr,
		"RunMode: " + gin.Mode(),
		"Method: " + c.Request.Method,
		"Status: " + strconv.Itoa(c.Writer.Status()),
		"NowTime: " + time.Now().Format("2006/01/02 15:04:05"),
		"ReqURI: " + strings.Join([]string{scheme, c.Request.Host, c.Request.RequestURI}, ""),
		"UserAgent: " + c.Request.UserAgent(),
	}
	posts := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": strings.Join(text, "\n") + "\n",
		},
		"at": map[string]interface{}{
			"isAtAll": true,
		},
	}
	_, _ = NewRequest(cons.DingTalkGroupRobot).SetHeaders(map[string]string{"Content-Type": "application/json;charset=utf-8"}).SetPostData(posts).Post()
}

func SystemIP() string {
	ip := "unknown"
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, addr := range addrs {
			// 检查ip地址判断是否回环地址
			if inet, ok := addr.(*net.IPNet); ok && !inet.IP.IsLoopback() {
				if inet.IP.To4() != nil {
					ip = inet.IP.String()
					break
				}

			}
		}
	}
	return ip
}

func PublicIP() string {
	ip := "Unknow"
	if resp, err := http.Get("http://ip.360.cn/IPShare/info"); err == nil {
		if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
			info := struct {
				Ip string
			}{}
			if err := json.Unmarshal(bytes, &info); err == nil {
				ip = info.Ip
			}
		}
	}
	return ip
}

/**
 * 服务健康检查
 * TODO::兼容旧的API网关
 */
func HealthCheckHttpHandler(rg *gin.RouterGroup) {
	rg.GET("/HealthCheck", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, `"ok"`)
	})
}
