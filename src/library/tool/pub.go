package tool

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
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
	strArr := []rune(str)
	if strArr[0] < 97 || strArr[0] > 122 {
		strArr[0] += 32
	}
	return string(strArr)
}

func PushSimpleMessage(msg string, isAtAll bool) {
	if !app.RobotEnable {
		return
	}
	host, _ := os.Hostname()
	text := []string{
		msg + "\n---------------------------",
		"App: " + app.Name,
		"Mode: " + gin.Mode(),
		"Listen: " + *app.HttpServerAddr,
		"HostName: " + host,
		"Time: " + time.Now().Format("2006/01/02 15:04:05.000"),
		"SystemIP: " + SystemIP(),
		"PublicIP: " + PublicIP(),
	}
	postData := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": strings.Join(text, "\n") + "\n",
		},
		"at": map[string]interface{}{
			"isAtAll": isAtAll,
		},
	}
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + app.GracefulRobotToken
	_, _ = NewRequest(url).SetHeaders(map[string]string{"Content-Type": "application/json;charset=utf-8"}).SetPostData(postData).Post()
}

func PushContextMessage(ctx *gin.Context, msg string, reqStr string, stack string, isAtAll bool) {
	if !app.RobotEnable {
		return
	}
	host, _ := os.Hostname()
	text := []string{
		msg + "\n---------------------------",
		"App: " + app.Name,
		"Mode: " + gin.Mode(),
		"Listen: " + *app.HttpServerAddr,
		"HostName: " + host,
		"Time: " + time.Now().Format("2006/01/02 15:04:05.000"),
		"Url: " + "https://" + ctx.Request.Host + ctx.Request.URL.String(),
		"SystemIP: " + SystemIP(),
		"PublicIP: " + PublicIP(),
		"ClientIP: " + ClientIP(ctx.ClientIP()),
	}
	if reqStr != "" {
		text = append(text, "\nRequest:\n"+strings.ReplaceAll(reqStr, "\n", ""))
	}
	if stack != "" {
		text = append(text, "\nStack:\n"+stack)
	}
	postData := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": strings.Join(text, "\n") + "\n",
		},
		"at": map[string]interface{}{
			"isAtAll": isAtAll,
		},
	}
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + app.ErrorRobotToken
	_, _ = NewRequest(url).SetHeaders(map[string]string{"Content-Type": "application/json;charset=utf-8"}).SetPostData(postData).Post()
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

// 阻塞式的执行外部shell命令的函数, 等待执行完毕并返回标准输出
func ExecShell(name string, arg ...string) ([]byte, error) {
	// 函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(name, arg...)

	// 读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为[]byte类型
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run执行c包含的命令，并阻塞直到完成。这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了。
	if err := cmd.Run(); err != nil {
		return nil, err
	} else {
		return out.Bytes(), nil
	}
}

func GetIPByIPCN(ip string) string {
	url := "https://www.ip.cn/"
	if ip != "" {
		url += "?ip=" + ip
	}
	if bts, err := ExecShell("curl", "--connect-timeout", "1", "-m", "1", url); err != nil || bts == nil {
		return "unknown"
	} else {
		out := struct{ Ip, Country, City string }{}
		if app.Json.Unmarshal(bts, &out) != nil {
			return "unknown"
		}
		return fmt.Sprintf("[%s]%s %s", out.Ip, out.Country, out.City)
	}
}

func GetIPByCIPCC(ip string) string {
	var buf bytes.Buffer
	buf.WriteString("http://cip.cc/")
	if ip != "" {
		buf.WriteString(ip)
	}
	headers := map[string]string{`User-Agent`: `curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.27.1 zlib/1.2.3 libidn/1.18 libssh2/1.4.2`}
	resp, err := NewRequest(buf.String()).SetHeaders(headers).SetTimeOut(1 * time.Second).Get()
	if err != nil || !strings.Contains(resp.Body, "IP") {
		return "unknown"
	}
	buf.Reset()
	for k, v := range strings.Split(resp.Body, "\n") {
		if k > 2 {
			break
		}
		if strings.Contains(v, ":") {
			slit := strings.Split(v, ":")
			if len(slit) != 2 {
				continue
			}
			switch strings.TrimSpace(slit[0]) {
			case "IP":
				buf.WriteString("[" + strings.TrimSpace(slit[1]) + "]")
			case "地址":
				for _, v := range strings.Split(strings.TrimSpace(slit[1]), "  ") {
					if !strings.Contains(buf.String(), strings.TrimSpace(v)) {
						buf.WriteString(v)
					}
				}
			case "运营商":
				operator := strings.TrimSpace(strings.Split(slit[1], "/")[0])
				if !strings.Contains(buf.String(), operator) {
					buf.WriteString(operator)
				}
			}
		}
	}
	return buf.String()
}

func GetIPByPconline(ip string) string {
	var result, url = ip, "http://whois.pconline.com.cn/ipJson.jsp?json=true"
	if ip != "" {
		url += "&ip=" + ip
	}
	resp, err := NewRequest(url).SetTimeOut(1 * time.Second).Get()
	if err != nil || resp.Body == "" {
		return result
	}
	body := mahonia.NewDecoder("GBK").ConvertString(resp.Body)
	if body == "" {
		return result
	}
	info := struct{ Ip, Addr string }{}
	if app.Json.Unmarshal([]byte(body), &info) != nil {
		return result
	}
	if info.Ip != "" && info.Addr != "" {
		result = fmt.Sprintf("%s %s", info.Ip, strings.ReplaceAll(info.Addr, " ", ""))
	}
	return result
}

func PublicIP() string {
	return GetIPByPconline("")
}

func ClientIP(ip string) string {
	return GetIPByPconline(ip)
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

func GetQueryPostForm(ctx *gin.Context, keys ...string) string {
	if len(keys) == 0 {
		return ""
	}
	if val, ok := ctx.GetQuery(keys[0]); ok {
		return val
	}
	if val, ok := ctx.GetPostForm(keys[0]); ok {
		return val
	}
	if len(keys) == 2 {
		return keys[1]
	}
	return ""
}
