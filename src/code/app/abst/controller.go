package abst

import (
	"bytes"
	"code/app/logger"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// bytes Buffer池
var bufferPool sync.Pool

type MethodConfig struct {
	HttpMethods              []string
	ControllerNameFirstUpper bool
	MethodNameFirstUpper     bool
}

type Controller struct {
	Config map[string]MethodConfig
}

func init() {
	bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 4096))
		},
	}
}

func (c Controller) HttpPost(ctx *gin.Context, uri string, timeout time.Duration) (string, error) {
	if uri == "" {
		return "", errors.New("request uri cannot be an empty string")
	}
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	cli := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, uri, ctx.Request.Body); err != nil {
		return "", err
	} else {
		if ctx.Request.URL.RawQuery != "" {
			req.URL.RawQuery = ctx.Request.URL.RawQuery
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if resp, err := cli.Do(req.WithContext(ctx)); err != nil {
			return "", err
		} else {
			defer resp.Body.Close()
			buffer := bufferPool.Get().(*bytes.Buffer)
			buffer.Reset()
			defer func() {
				if buffer != nil {
					bufferPool.Put(buffer)
					buffer = nil
				}
			}()
			if _, err = io.Copy(buffer, resp.Body); err != nil {
				return "", fmt.Errorf("io.copy error: %v", err)
			}
			bodyStr := buffer.String()
			bufferPool.Put(buffer)
			buffer = nil
			return bodyStr, nil
		}
	}
}

func (c Controller) ServiceRewrite(ctx *gin.Context, uri string, options ...interface{}) {
	timeout := 10 * time.Second
	if len(options) > 0 {
		timeout = options[0].(time.Duration)
	}
	if respBody, err := c.HttpPost(ctx, uri, timeout); err == nil {
		ctx.String(http.StatusOK, respBody)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": err.Error(), "data": nil})
	}
}

// ServiceMultiRewrite 把 multiKey 分 corNum 个协程并行去发起请求
func (c Controller) ServiceMultiRewrite(ctx *gin.Context, uri string, multiKey string, corNum int, respBodyHandler func([]string), options ...interface{}) {
	if multiKey == "" {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "multiKey is empty", "data": nil})
		return
	}
	if corNum == 0 {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "corNum is empty", "data": nil})
		return
	}
	reqMethod := http.MethodPost
	multiValue := ctx.DefaultPostForm(multiKey, "")
	if multiValue == "" {
		multiValue = ctx.DefaultQuery(multiKey, "")
		reqMethod = http.MethodGet
	}
	if multiValue == "" {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "multiValue is empty", "data": nil})
		return
	}

	multiSlice := strings.Split(multiValue, ",")
	multiSliceLen := len(multiSlice)

	/*logger.Debug("", "ServiceMultiRewrite",
		zap.String("multiValue", multiValue),
		zap.Any("multiSlice", multiSlice),
		zap.String("reqMethod", reqMethod),
	)*/

	timeout := 10 * time.Second
	if len(options) > 0 {
		timeout = options[0].(time.Duration)
	}

	step := 1
	if multiSliceLen > corNum {
		step = multiSliceLen / corNum
		if multiSliceLen%corNum > 0 {
			corNum++
		}
	} else {
		step = multiSliceLen
		corNum = 1
	}

	var wg sync.WaitGroup
	wg.Add(corNum)
	respBodes := make([]string, corNum)
	for i := 0; i < corNum; i++ {
		start := i * step
		end := start + step
		if end > multiSliceLen {
			end = multiSliceLen
		}
		go func(ctxPy *gin.Context, start, end int) {
			sliceStr := strings.Join(multiSlice[start:end], ",")
			//logger.Debug("", "corNum", zap.Int("start", start), zap.Int("end", end), zap.String("sliceStr", sliceStr))
			switch reqMethod {
			case http.MethodPost:
				ctxPy.Request.PostForm.Set(multiKey, sliceStr)
			case http.MethodGet:
				query := ctxPy.Request.URL.Query()
				query.Set(multiKey, sliceStr)
				ctxPy.Request.URL.RawQuery = query.Encode()
			}
			if respBody, err := c.HttpPost(ctxPy, uri, timeout); err == nil && respBody != "" {
				//logger.Debug("", "respBody", zap.String("respBody", respBody))
				respBodes = append(respBodes, respBody)
			} else {
				logger.Error("", "ServiceMultiRewrite->httpPost Error", zap.Error(err), zap.String("body", respBody))
			}
			wg.Done()
		}(ctx.Copy(), start, end)
	}
	wg.Wait()

	if respBodyHandler == nil {
		type respBody struct {
			Data map[string]interface{}
		}
		data := make(map[string]interface{})
		for _, body := range respBodes {
			if body != "" {
				obj := respBody{}
				if err := json.Unmarshal([]byte(body), &obj); err == nil {
					if len(obj.Data) > 0 {
						for k, v := range obj.Data {
							data[k] = v
						}
					}
				} else {
					logger.Error("", "ServiceMultiRewrite->respBodyHandler Unmarshal Error", zap.Error(err))
				}
			}
		}
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "", "data": data})
	} else {
		respBodyHandler(respBodes)
	}
}
