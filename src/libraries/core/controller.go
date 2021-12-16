package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"runtime/debug"
	appEnv "src/config/env"
	"src/libraries/helper"
	"src/libraries/logger"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

//  MtdCfg ...
//  @author Cloud|2021-12-14 13:01:08
type MtdCfg map[string]MethodConfig

//  IController ...
//  @author Cloud|2021-12-14 13:01:08
type IController interface {
	Init() MtdCfg
}

//  Controller ...
//  @author Cloud|2021-12-14 13:01:04
type Controller struct {
	Config MtdCfg
}

//  MethodConfig ...
//  @author Cloud|2021-12-14 13:01:01
type MethodConfig struct {
	Http        []string        // 请求方法
	Func        gin.HandlerFunc // 调用方法
	CNCap       bool            // 控制器名首字母大写
	MNCap       bool            // 方法名首字母大写
	Rest        string          // 接口地址别名
	SignCheck   bool            // APP签名校验
	XySignCheck bool            // 小易签名检验
	LoginCheck  bool            // APP登录检验
}

//  RPCXService ...
//  @receiver c *Controller
//  @author Cloud|2021-12-14 13:00:50
//  @param ctx *gin.Context ...
//  @param service string ...
//  @param method string ...
func (c *Controller) RPCXService(ctx *gin.Context, service, method string) {
	args := make(map[string]interface{})
	for key := range ctx.Request.URL.Query() {
		args[key] = ctx.Request.URL.Query().Get(key)
	}
	// 如果为Post请求，PostForm参数优先级高于Query参数
	if ctx.Request.Method == http.MethodPost {
		_ = ctx.Request.ParseForm()
		for key := range ctx.Request.PostForm {
			args[key] = ctx.Request.PostForm.Get(key)
		}
	}
	logger.SetLoggerAttr(service, ctx)
	reply, err := ForwardRPCXService(service, method, args)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}

	ctx.Abort()
	ctx.Status(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
	switch v := reply.(type) {
	case []byte:
		_, _ = ctx.Writer.Write(v)
	case string:
		_, _ = ctx.Writer.WriteString(v)
	default:
		_, _ = ctx.Writer.WriteString(fmt.Sprintf("%v", v))
	}
}

//  GetQueryPostForm ...
//  @receiver c *Controller
//  @author Cloud|2021-12-14 13:01:45
//  @param ctx *gin.Context ...
//  @param keys ...string ...
//  @return string ...
func (c *Controller) GetQueryPostForm(ctx *gin.Context, keys ...string) string {
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

//  ServiceRewrite ...
//  @receiver c *Controller
//  @author Cloud|2021-12-14 13:01:49
//  @param ctx *gin.Context ...
//  @param url string ...
//  @param options ...interface{} ...
func (c *Controller) ServiceRewrite(ctx *gin.Context, url string, options ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			errS, reqS, stackS := fmt.Sprintf("service rewrite recover: %v", err),
				ctx.GetString(appEnv.ReqStrKey), string(debug.Stack())
			logger.Use("rewrite").Error(errS)
			helper.PushContextMessage(ctx, errS, reqS, stackS, true)
			obj := gin.H{"code": http.StatusInternalServerError, "msg": errS}
			if gin.IsDebugging() {
				obj["request"] = strings.Split(reqS, "\r\n")
				obj["stack"] = strings.Split(strings.ReplaceAll(stackS, "\t", ""), "\n")
			}
			ctx.AbortWithStatusJSON(http.StatusOK, obj)
		}
	}()

	if url == "" {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusNotImplemented, "msg": "illegal request"})
		return
	}

	timeout := helper.DefaultHttpClientTimeout
	if len(options) > 0 {
		if val, ok := options[0].(time.Duration); ok && val > 0 {
			timeout = val
		}
	}

	var (
		err  error
		req  *http.Request
		resp *http.Response
	)

	if req, err = http.NewRequest(ctx.Request.Method, url, ctx.Request.Body); err != nil {
		msg := "http.NewRequest error: " + err.Error()
		logger.Use("rewrite").Error(msg)
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": msg})
		return
	}

	req.URL.RawQuery = ctx.Request.URL.RawQuery
	req.Header = ctx.Request.Header

	// 转发内部微服务不需要再去压缩，要不然返回客户端时候Nginx会进行二次压缩，导致内容无法解析！
	req.Header.Del("Accept-Encoding")

	clientIP := ctx.ClientIP()
	req.Header.Set("X-Real-Ip", clientIP)
	req.Header.Add("X-Forwarded-For", clientIP)
	req.Header.Add("User-Agent", "Go-Http-Client/1.1")

	ctxt, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if resp, err = helper.HttpClient.Do(req.WithContext(ctxt)); err != nil {
		msg := "http.Client.Do error: " + err.Error()
		logger.Use("rewrite").Error(msg)
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": msg})
		return
	}

	defer resp.Body.Close()

	for k, v := range resp.Header {
		ctx.Writer.Header().Del(k)
		for _, v2 := range v {
			ctx.Writer.Header().Add(k, v2)
		}
	}

	var bodyByte []byte
	if bodyByte, err = helper.IoCopy(resp.Body); err != nil {
		msg := "read resp.Body error: " + err.Error()
		logger.Use("rewrite").Error(msg)
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": msg})
		return
	}

	if gin.IsDebugging() {
		go c.respBodyCheck(ctx.Copy(), resp.Header.Get("Content-Type"), string(bodyByte))
	}

	_, _ = ctx.Writer.Write(bodyByte)
	ctx.AbortWithStatus(http.StatusOK)
}

//  ServiceMultiRewrite ...
//  @receiver c *Controller
//  @author Cloud|2021-12-14 13:07:31
//  @param ctx *gin.Context ...
//  @param urls string ...
//  @param key string ...
//  @param corNum uint32 ...
//  @param options ...interface{} ...
func (c *Controller) ServiceMultiRewrite(ctx *gin.Context, urls string, key string, corNum uint32, options ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			errS, reqS, stackS := fmt.Sprintf("service multi rewrite recover: %v", err),
				ctx.GetString(appEnv.ReqStrKey), string(debug.Stack())
			logger.Use("multi-rewrite").Error(errS)
			helper.PushContextMessage(ctx, errS, reqS, stackS, true)
			obj := gin.H{"code": http.StatusInternalServerError, "msg": errS}
			if gin.IsDebugging() {
				obj["request"] = strings.Split(reqS, "\r\n")
				obj["stack"] = strings.Split(strings.ReplaceAll(stackS, "\t", ""), "\n")
			}
			ctx.AbortWithStatusJSON(http.StatusOK, obj)
		}
	}()

	keyValue := ctx.DefaultPostForm(key, "")
	method := http.MethodPost
	if keyValue == "" {
		keyValue = ctx.DefaultQuery(key, "")
		method = http.MethodGet
	}

	if urls == "" || key == "" || corNum <= 0 || keyValue == "" {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusNotImplemented, "msg": "illegal request"})
		return
	}

	timeout := helper.DefaultHttpClientTimeout
	if len(options) > 0 {
		if val, ok := options[0].(time.Duration); ok && val > 0 {
			timeout = val
		}
	}

	keyValue, _ = url.QueryUnescape(keyValue)
	keySlice := strings.Split(keyValue, ",")
	keySliceLen := uint32(len(keySlice))

	step := keySliceLen / corNum
	if step > 0 {
		if keySliceLen%corNum > 0 {
			corNum++
		}
	} else {
		step = 1
		corNum = keySliceLen
	}

	var (
		headerOnce sync.Once
		respHeader http.Header
		chanNum    uint32
	)

	resChan := make(chan interface{}, corNum)
	dataMap := make(map[string]interface{}, keySliceLen) // 接口返回的 json 结果是 map 结构的时候使用
	var dataSlice []interface{}                          // 接口返回的 json 结果是 slice 结构的时候使用

	ccp := ctx.Copy()
	_ = ccp.Request.ParseForm()

	for i := uint32(0); i < corNum; i++ {
		start := i * step
		end := start + step
		if end > keySliceLen {
			end = keySliceLen
		}

		go func(start, end uint32) {
			defer func() {
				if err := recover(); err != nil {
					txt := fmt.Sprintf("service multi rewrite recover: %v", err)
					logger.Use("multi-rewrite").Error(txt)
					helper.PushSimpleMessage(txt, true)
				}
			}()

			var succeed bool
			defer func() {
				if !succeed {
					// 程序遇到异常而中断, 被 return 后, 返回 "" 到 channel 中, 避免协程一直在等待
					helper.SafeSendChan(resChan, "")
				}
			}()

			defer atomic.AddUint32(&chanNum, 1)

			var (
				err  error
				req  *http.Request
				resp *http.Response
			)

			keyStr := strings.Join(keySlice[start:end], ",")
			if keyStr == "" {
				return
			}

			if req, err = http.NewRequest(http.MethodPost, urls, nil); err != nil {
				logger.Use("multi-rewrite").Error("http.NewRequest error: " + err.Error())
				return
			}

			switch method {
			case http.MethodPost:
				postForm := url.Values{}
				for k, v := range ccp.Request.PostForm {
					for _, v2 := range v {
						postForm.Add(k, v2)
					}
				}
				postForm.Set(key, keyStr)
				req.Body = ioutil.NopCloser(bytes.NewBufferString(postForm.Encode()))
				req.URL.RawQuery = ccp.Request.URL.RawQuery
			case http.MethodGet:
				rawQuery := url.Values{}
				for k, v := range ccp.Request.URL.Query() {
					for _, v2 := range v {
						rawQuery.Add(k, v2)
					}
				}
				rawQuery.Set(key, keyStr)
				req.URL.RawQuery = rawQuery.Encode()
			}

			for k, v := range ccp.Request.Header {
				req.Header.Del(k)
				for _, v2 := range v {
					req.Header.Add(k, v2)
				}
			}

			// 转发内部微服务不需要再去压缩，要不然返回客户端时候Nginx会进行二次压缩，导致内容无法解析！
			req.Header.Del("Accept-Encoding")

			clientIP := ccp.ClientIP()
			req.Header.Set("X-Real-Ip", clientIP)
			req.Header.Add("X-Forwarded-For", clientIP)
			req.Header.Add("User-Agent", "Go-Http-Client/1.1")

			ctxt, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			if resp, err = helper.HttpClient.Do(req.WithContext(ctxt)); err != nil {
				logger.Use("multi-rewrite").Error("http.Client.Do error: " + err.Error())
				return
			}

			defer resp.Body.Close()

			headerOnce.Do(func() {
				respHeader = resp.Header
			})

			var bodyB []byte
			if bodyB, err = helper.IoCopy(resp.Body); err != nil {
				logger.Use("multi-rewrite").Error("read resp.Body error: " + err.Error())
				return
			}

			var res map[string]interface{}
			if err := json.Unmarshal(bodyB, &res); err != nil {
				logger.Use("multi-rewrite").Error("json.Unmarshal error: "+err.Error(),
					zap.String("body", string(bodyB)))
				return
			}

			if data, ok := res["data"]; ok {
				helper.SafeSendChan(resChan, data)
				succeed = true
			}
		}(start, end)
	}

LOOP:
	for {
		select {
		case <-time.After(timeout + time.Second):
			helper.SafeCloseChan(resChan)
			break LOOP
		case res := <-resChan:
			switch res.(type) {
			case map[string]interface{}:
				for k, v := range res.(map[string]interface{}) {
					dataMap[k] = v
				}
			case []interface{}:
				for _, v := range res.([]interface{}) {
					dataSlice = append(dataSlice, v)
				}
			}

			if atomic.LoadUint32(&chanNum) == corNum {
				helper.SafeCloseChan(resChan)
				break LOOP
			}
		}
	}

	for k, v := range respHeader {
		ctx.Writer.Header().Del(k)
		for _, v2 := range v {
			ctx.Writer.Header().Add(k, v2)
		}
	}

	if len(dataSlice) > 0 {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "", "data": dataSlice})
	} else {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "", "data": dataMap})
	}
}

//  respBodyCheck ...
//  @receiver c *Controller
//  @author Cloud|2021-12-14 13:07:18
//  @param ctx *gin.Context ...
//  @param respContentType string ...
//  @param respBody string ...
func (c *Controller) respBodyCheck(ctx *gin.Context, respContentType, respBody string) {
	defer helper.SafeDefer()
	if strings.Index(respContentType, "application/json") == 0 {
		var data interface{}
		if err := json.Unmarshal([]byte(respBody), &data); err != nil {
			html := respBody
			// 去除所有尖括号内的HTML代码
			re, _ := regexp.Compile(`<[\S\s]+?\>`)
			html = re.ReplaceAllString(html, "")
			msg := `resp.Header.Content-Type is application/json, but json.Unmarshal failed`
			helper.PushContextMessage(ctx, msg, ctx.GetString(appEnv.ReqStrKey), html, true)
		}
	}
}
