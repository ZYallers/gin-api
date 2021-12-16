package helper

import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"net"
	"strings"
	"time"
)

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
	if json.Unmarshal([]byte(body), &info) != nil {
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
