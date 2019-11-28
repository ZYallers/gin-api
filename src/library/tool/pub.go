package tool

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"runtime"
	"src/config"
	"strings"
	"time"
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

func SendDingTalkGroupMessage(ctx *gin.Context, msg string, reqstr string, stack string, isAtAll bool) {
	if !app.DingTalkMsgEnable {
		return
	}
	host, _ := os.Hostname()
	text := []string{
		msg + "\n------------------",
		"AppName: " + app.Name,
		"RunMode: " + gin.Mode(),
		"ListenAddr: " + *app.HttpServerAddr,
		"HostName: " + host,
		"ClientIP: " + ctx.ClientIP(),
		"PublicIP: " + PublicIP(),
		"SystemIP: " + SystemIP(),
		"URL: " + "https://" + ctx.Request.Host + ctx.Request.URL.String(),
		"NowTime: " + time.Now().Format("2006.01.02 15:04:05.000"),
	}
	if reqstr != "" {
		text = append(text, "\nRequest: "+strings.ReplaceAll(reqstr, "\n", ""))
	}
	if stack != "" {
		text = append(text, "\nDebugStack: "+strings.ReplaceAll(stack, "\n", ""))
	}
	posts := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": strings.Join(text, "\n") + "\n",
		},
		"at": map[string]interface{}{
			"isAtAll": isAtAll,
		},
	}
	_, _ = NewRequest(app.DingTalkGroupRobot).SetHeaders(map[string]string{"Content-Type": "application/json;charset=utf-8"}).SetPostData(posts).Post()
}

func SystemIP() string {
	if netInterfaces, err := net.Interfaces(); err == nil {
		for i := 0; i < len(netInterfaces); i++ {
			if (netInterfaces[i].Flags & net.FlagUp) != 0 {
				addrs, _ := netInterfaces[i].Addrs()

				for _, address := range addrs {
					if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
						if ipnet.IP.To4() != nil {
							return ipnet.IP.String()
						}
					}
				}
			}
		}
	}
	return "unknown"
}

func PublicIP() string {
	if resp, err := http.Get("http://ifconfig.co/ip"); err == nil {
		defer resp.Body.Close()
		if bs, err := ioutil.ReadAll(resp.Body); err == nil {
			return strings.Trim(string(bs), "\n")
		}
	}
	return "unknown"
}

func NowMemStats() string {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	return fmt.Sprintf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes) NumGoroutine:%d", ms.Alloc, ms.HeapIdle, ms.HeapReleased, runtime.NumGoroutine())
}

func SafeSendChan(ch chan<- interface{}, value interface{}) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value
	return false
}

func SafeCloseChan(ch chan interface{}) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = false
		}
	}()
	close(ch)
	return true
}

func DrainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
