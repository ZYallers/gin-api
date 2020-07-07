package abs

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	app "src/config"
	"src/library/logger"
	"src/library/tool"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Controller struct {
	Config map[string]MethodConfig
}

type MethodConfig struct {
	HttpMethods              []string
	ControllerNameFirstUpper bool
	MethodNameFirstUpper     bool
	Rest                     string
}

func (c Controller) ServiceRewrite(ctx *gin.Context, url string, options ...interface{}) {
	if url == "" {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "url cannot be an empty string"})
		return
	}

	timeout := time.Second * 10
	if len(options) > 0 {
		timeout = options[0].(time.Duration)
	}

	var (
		err  error
		req  *http.Request
		resp *http.Response
	)

	if req, err = http.NewRequest(ctx.Request.Method, url, ctx.Request.Body); err != nil {
		go logger.Use("rewrite").Error("http.NewRequest error: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "http.NewRequest error: " + err.Error()})
		return
	}

	req.URL.RawQuery = ctx.Request.URL.RawQuery

	req.Header = ctx.Request.Header
	// 主动断开，响应更快，占用资源减少！
	req.Close = true
	req.Header.Add("Connection", "close")
	// 不要用编码类型去压缩，要不然body无法解析！
	req.Header.Del("Accept-Encoding")
	// 排查测试环境问题，临时使用！
	if app.ResetXForwardFor {
		req.Header.Set("X-Forwarded-For", ctx.ClientIP())
	}
	req.Header.Set("X-Real-Ip", ctx.ClientIP())
	req.Header.Add("User-Agent", "Go-http-client/1.1")

	ctxt, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client := http.Client{Transport: &http.Transport{DisableKeepAlives: true}}
	if resp, err = client.Do(req.WithContext(ctxt)); err != nil {
		go logger.Use("rewrite").Error("http.Client.Do error: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "http.DefaultClient.Do error: " + err.Error()})
		return
	}

	defer resp.Body.Close()

	for k, v := range resp.Header {
		ctx.Writer.Header().Del(k)
		for _, v2 := range v {
			ctx.Writer.Header().Add(k, v2)
		}
	}

	// 如果在本地IP+Port访问情况下，没有压缩程序会无法返回内容给前端，为方便本地测试加的！
	if strings.Contains(ctx.Request.Host, "127.0.0.1") {
		ctx.Writer.Header().Del("Content-Encoding")
	}

	if respBodyBytes, err := ioutil.ReadAll(resp.Body); err != nil {
		go logger.Use("rewrite").Error("ioutil.ReadAll error: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "ioutil.ReadAll error: " + err.Error()})
	} else {
		go func(ctx *gin.Context, bts []byte) {
			if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
				var data interface{}
				if err := app.Json.Unmarshal(bts, &data); err != nil {
					tool.PushContextMessage(ctx, "response content type is json, but unmarshal fails",
						ctx.GetString(app.ReqStrKey), string(bts), true)
				}
			}
		}(ctx.Copy(), respBodyBytes)
		ctx.String(http.StatusOK, string(respBodyBytes))
		ctx.AbortWithStatus(http.StatusOK)
	}
}

func (c Controller) ServiceMultiRewrite(ctx *gin.Context, urls string, key string, corNum uint32, options ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			go logger.Use("multi-rewrite").Error("recover error", zap.Any("error", err))
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "recover error"})
		}
	}()

	if urls == "" {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "url cannot be an empty string"})
		return
	}

	if key == "" {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "key cannot be an empty string"})
		return
	}

	if corNum <= 0 {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "corNum cannot be less than zero"})
		return
	}

	keyValue := ctx.DefaultPostForm(key, "")
	method := http.MethodPost
	if keyValue == "" {
		keyValue = ctx.DefaultQuery(key, "")
		method = http.MethodGet
	}
	if keyValue == "" {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "keyValue cannot be an empty string"})
		return
	}

	timeout := 10 * time.Second
	if len(options) > 0 {
		timeout = options[0].(time.Duration)
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

	//logger.Use("multi-info").Info("", zap.String("key", key), zap.String("value", keyValue), zap.Any("corNum", corNum))

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
					go logger.Use("multi-rewrite").Error("go recover error", zap.Any("error", err))
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
				go logger.Use("multi-rewrite").Error("http.NewRequest error: " + err.Error())
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
			// 主动断开，响应更快，占用资源减少！
			req.Close = true
			req.Header.Add("Connection", "close")
			// 不要用编码类型去压缩，要不然body无法解析！
			req.Header.Del("Accept-Encoding")
			// 排查测试环境问题，临时使用！
			if app.ResetXForwardFor {
				req.Header.Set("X-Forwarded-For", ccp.ClientIP())
			}
			req.Header.Add("User-Agent", "Go-http-client/1.1")

			ctxt, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			client := http.Client{Transport: &http.Transport{DisableKeepAlives: true}}
			if resp, err = client.Do(req.WithContext(ctxt)); err != nil {
				go logger.Use("multi-rewrite").Error("http.DefaultClient.Do error: " + err.Error())
				return
			}

			defer resp.Body.Close()

			headerOnce.Do(func() {
				respHeader = resp.Header
			})

			var body []byte
			if body, err = ioutil.ReadAll(resp.Body); err != nil {
				go logger.Use("multi-rewrite").Error("ioutil.ReadAll error: " + err.Error())
				return
			}
			if body == nil {
				return
			}
			var res map[string]interface{}
			//logger.Use("multi-body").Info("body string", zap.String("body", string(body)))
			if err := app.Json.Unmarshal(body, &res); err != nil {
				go logger.Use("multi-rewrite").Error("json.Unmarshal", zap.Error(err), zap.String("body", string(body)))
				return
			}
			if data, ok := res[`data`]; ok {
				tool.SafeSendChan(resChan, data)
			}
		}(start, end)
	}

GOTO:
	for {
		select {
		case <-time.After(timeout + time.Millisecond*100):
			tool.SafeCloseChan(resChan)
			break GOTO
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
				tool.SafeCloseChan(resChan)
				break GOTO
			}
		}
	}

	for k, v := range respHeader {
		ctx.Writer.Header().Del(k)
		for _, v2 := range v {
			ctx.Writer.Header().Add(k, v2)
		}
	}

	// 如果在本地IP+Port访问情况下，没有压缩程序会无法返回内容给前端，为方便本地测试加的！
	if strings.Contains(ctx.Request.Host, "127.0.0.1") {
		ctx.Writer.Header().Del("Content-Encoding")
	}

	if len(dataSlice) > 0 {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "", "data": dataSlice})
	} else {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "", "data": dataMap})
	}
}
