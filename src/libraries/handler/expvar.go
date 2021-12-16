package handler

import (
	"encoding/json"
	"expvar"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var (
	lastPause uint32
	start     = time.Now() // 开始时间
)

func init() {
	// 这些都是我自定义的变量，发布到 expvar 中，每次请求接口，expvar会自动去获取这些变量，并返回给我
	expvar.Publish("A:运行时间", expvar.Func(caleRuntime))
	expvar.Publish("B:Go语言版本", expvar.Func(GoVersion))
	expvar.Publish("D:系统CPU数量", expvar.Func(getNumCPUs))
	expvar.Publish("C:Go运行系统", expvar.Func(getGoOS))
	expvar.Publish("F:CGO调用次数", expvar.Func(getNumCgoCall))
	expvar.Publish("E:协程数量", expvar.Func(getNumGoroutins))
	expvar.Publish("G:上次GC的暂停时间", expvar.Func(getLastGCPauseTime))
}

// 计算运行时间
func caleRuntime() interface{} {
	return time.Since(start).String()
}

// 当前 Golang 版本
func GoVersion() interface{} {
	return runtime.Version()
}

// 获取 CPU 核心数量
func getNumCPUs() interface{} {
	return runtime.NumCPU()
}

// 当前系统类型
func getGoOS() interface{} {
	return runtime.GOOS
}

// 当前 goroutine 数量
func getNumGoroutins() interface{} {
	return runtime.NumGoroutine()
}

// CGo 调用次数
func getNumCgoCall() interface{} {
	return runtime.NumCgoCall()
}

// 获取上次 GC 的暂停时间
func getLastGCPauseTime() interface{} {
	var gcPause uint64
	ms := new(runtime.MemStats)
	statString := expvar.Get("memstats").String()
	if statString != "" {
		_ = json.Unmarshal([]byte(statString), ms)
		if lastPause == 0 || lastPause != ms.NumGC {
			gcPause = ms.PauseNs[(ms.NumGC+255)%256]
			lastPause = ms.NumGC
		}
	}
	return gcPause
}

func sizeFormat(b uint64) string {
	i := 0
	// b大于是1024字节时，开始循环，当循环到第4次时跳出
	flat, _ := strconv.ParseFloat(strconv.FormatInt(int64(b), 10), 64)
	for {
		if flat >= 1024 {
			flat, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", flat/1024), 64)
			i++
		}
		if flat < 1024 || i == 4 {
			break
		}
	}
	//将B,KB,MB,GB,TB定义成一维数组
	units := []string{"B", "KB", "MB", "GB", "TB"}
	return fmt.Sprintf("%.2f%s", flat, units[i])
}

// 返回当前运行信息
func ExpHandler(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
	first := true
	report := func(key string, value interface{}) {
		if !first {
			_, _ = fmt.Fprintf(ctx.Writer, ",\n")
		}
		first = false
		if str, ok := value.(string); ok {
			_, _ = fmt.Fprintf(ctx.Writer, "%q: %q", key, str)
		} else {
			switch key {
			case "cmdline":
				_, _ = fmt.Fprintf(ctx.Writer, "%q: %v", "H:程序路径", value)
			case "memstats":
				if str, ok := value.(expvar.Func); ok {
					ms := runtime.MemStats{}
					_ = json.Unmarshal([]byte(str.String()), &ms)
					obj := make(map[string]interface{}, 30)
					obj["Alloc:框架堆空间分配"] = sizeFormat(ms.Alloc)
					obj["TotalAlloc:从服务开始运行至今分配器为分配的堆空间总和(只有增加,释放的时候不减少)"] = sizeFormat(ms.TotalAlloc)
					obj["Sys:服务现在系统使用的内存"] = sizeFormat(ms.Sys)
					obj["Lookups:被runtime监视的指针数"] = ms.Lookups
					obj["Mallocs:服务malloc的次数"] = ms.Mallocs
					obj["Frees:服务回收的HeapObjects字节数"] = sizeFormat(ms.Frees)
					obj["HeapAlloc:服务分配的堆内存字节数"] = sizeFormat(ms.HeapAlloc)
					obj["HeapSys:系统分配的作为运行栈的内存"] = sizeFormat(ms.HeapSys)
					obj["HeapIdle:申请但是未分配的堆内存或者回收了的堆内存(空闲)字节数"] = sizeFormat(ms.HeapIdle)
					obj["HeapInuse:正在使用的堆内存字节数"] = sizeFormat(ms.HeapInuse)
					obj["HeapReleased:返回给OS的堆内存"] = sizeFormat(ms.HeapReleased)
					obj["HeapObjects:堆内存块申请的数量"] = ms.HeapObjects
					obj["StackInuse:正在使用的栈字节数"] = sizeFormat(ms.StackInuse)
					obj["StackSys:系统分配的作为运行栈的内存"] = sizeFormat(ms.StackSys)
					obj["MSpanInuse:用于测试用的结构体使用的字节数"] = sizeFormat(ms.MSpanInuse)
					obj["MSpanSys:系统为测试用的结构体分配的字节数"] = sizeFormat(ms.MSpanSys)
					obj["MCacheInuse:结构体MCache申请的字节数(不会被视为垃圾回收)"] = sizeFormat(ms.MCacheInuse)
					obj["MCacheSys:操作系统申请的堆空间用于mcache的字节数"] = sizeFormat(ms.MCacheSys)
					obj["BuckHashSys:用于剖析桶散列表的堆空间"] = sizeFormat(ms.BuckHashSys)
					obj["GCSys:垃圾回收标记元信息使用的内存"] = sizeFormat(ms.GCSys)
					obj["OtherSys:系统架构占用的额外空间"] = sizeFormat(ms.OtherSys)
					obj["NextGC:垃圾回收器检视的内存大小"] = sizeFormat(ms.NextGC)
					obj["PauseTotalNs:垃圾回收或者其他信息收集导致服务暂停的次数"] = ms.PauseTotalNs
					obj["NumForcedGC:服务调用runtime.GC()强制使用垃圾回收的次数"] = ms.NumForcedGC
					obj["GCCPUFraction:垃圾回收占用服务CPU工作的时间总和(如果有100个协程,垃圾回收时间为1s,那么就占用了100s)"] = ms.GCCPUFraction
					//obj["PauseNs:记录最近垃圾回收系统中断的时间"] = ms.PauseNs
					//obj["PauseEnd:记录最近垃圾回收系统中断的时间开始点"] = ms.PauseEnd
					//obj["BySize:内存分配器使用情况"] = ms.BySize
					bt, _ := json.Marshal(obj)
					_, _ = fmt.Fprintf(ctx.Writer, "%q: %s", "I:内存状态", string(bt))
				} else {
					_, _ = fmt.Fprintf(ctx.Writer, "%q: %v", key, value)
				}
			default:
				_, _ = fmt.Fprintf(ctx.Writer, "%q: %v", key, value)
			}
		}
	}
	_, _ = fmt.Fprintf(ctx.Writer, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	_, _ = fmt.Fprintf(ctx.Writer, "\n}\n")

	ctx.String(http.StatusOK, "")
}
