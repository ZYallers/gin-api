package tool

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// bytes Buffer池
var bufferPool sync.Pool

// Request构造类
type Request struct {
	cli      *http.Client
	req      *http.Request
	Method   string
	Url      string
	Timeout  time.Duration
	Headers  map[string]string
	Cookies  map[string]string
	Queries  map[string]string
	PostData map[string]interface{}
}

func init() {
	bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 4096))
		},
	}
}

// 创建一个Request实例
func NewRequest(url string) *Request {
	this := &Request{}
	if url != "" {
		this.Url = url
	}
	return this
}

// 设置请求方法
func (r *Request) SetMethod(method string) *Request {
	r.Method = method
	return r
}

// 设置请求地址
func (r *Request) SetUrl(url string) *Request {
	r.Url = url
	return r
}

// 设置请求头
func (r *Request) SetHeaders(headers map[string]string) *Request {
	r.Headers = headers
	return r
}

// 将用户自定义请求头添加到http.Request实例上
func (r *Request) setHeaders() *Request {
	for k, v := range r.Headers {
		r.req.Header.Set(k, v)
	}
	return r
}

// 设置请求cookies
func (r *Request) SetCookies(cookies map[string]string) *Request {
	r.Cookies = cookies
	return r
}

// 将用户自定义cookies添加到http.Request实例上
func (r *Request) setCookies() *Request {
	for k, v := range r.Cookies {
		r.req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
	return r
}

// 设置url查询参数
func (r *Request) SetQueries(queries map[string]string) *Request {
	r.Queries = queries
	return r
}

// 将用户自定义url查询参数添加到http.Request上
func (r *Request) setQueries() *Request {
	q := r.req.URL.Query()
	for k, v := range r.Queries {
		q.Add(k, v)
	}
	r.req.URL.RawQuery = q.Encode()
	return r
}

// 设置post请求的提交数据
func (r *Request) SetPostData(postData map[string]interface{}) *Request {
	r.PostData = postData
	return r
}

// 发起get请求
func (r *Request) Get() (*Response, error) {
	return r.Send(r.Url, http.MethodGet)
}

// 发起post请求
func (r *Request) Post() (*Response, error) {
	return r.Send(r.Url, http.MethodPost)
}

// SetDialTimeOut
func (r *Request) SetTimeOut(timeout time.Duration) *Request {
	r.Timeout = timeout
	return r
}

func (r *Request) elapsedTime(n int64, resp *Response) {
	end := time.Now().UnixNano() / 1e6
	resp.spendTime = end - n
}

// 发起请求
func (r *Request) Send(url string, method string) (*Response, error) {
	start := time.Now().UnixNano() / 1e6
	if url == "" {
		return nil, errors.New("request url")
	}
	if method == "" {
		return nil, errors.New("request method")
	}
	response := NewResponse()
	defer r.elapsedTime(start, response)

	r.cli = &http.Client{}

	var body io.Reader
	if postLen := len(r.PostData); postLen > 0 {
		contentType, _ := r.Headers["Content-Type"]
		switch strings.ToLower(contentType) {
		case "application/json", "application/json;charset=utf-8":
			if bys, err := json.Marshal(r.PostData); err != nil {
				return nil, err
			} else {
				body = bytes.NewReader(bys)
			}
		default:
		case "application/x-www-form-urlencoded":
			var (
				bf bytes.Buffer
				i  int
			)
			for k, v := range r.PostData {
				i++
				bf.WriteString(k)
				bf.WriteString("=")
				bf.WriteString(fmt.Sprintf("%v", v))
				if i < postLen {
					bf.WriteString("&")
				}
			}
			body = strings.NewReader(bf.String())
		}
	}

	if req, err := http.NewRequest(method, url, body); err != nil {
		return nil, err
	} else {
		if r.Timeout <= 0 || r.Timeout > 30 {
			r.Timeout = 30 * time.Second
		}
		ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
		defer cancel()
		r.req = req.WithContext(ctx)
	}

	r.setHeaders().setCookies().setQueries()

	if resp, err := r.cli.Do(r.req); err != nil {
		return nil, err
	} else {
		response.Raw = resp
		defer response.Raw.Body.Close()
		response.parseHeaders()
		if err := response.parseBody(); err != nil {
			return nil, err
		}
		return response, nil
	}
}

// Response 构造类
type Response struct {
	Raw       *http.Response
	Headers   map[string]string
	Body      string
	spendTime int64
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) StatusCode() int {
	if r.Raw == nil {
		return 0
	}
	return r.Raw.StatusCode
}

func (r *Response) IsOk() bool {
	return r.StatusCode() == http.StatusOK
}

func (r *Response) SpendTime() string {
	return strconv.Itoa(int(r.spendTime)) + "ms"
}

func (r *Response) parseHeaders() {
	headers := map[string]string{}
	for k, v := range r.Raw.Header {
		headers[k] = v[0]
	}
	r.Headers = headers
}

func (r *Response) parseBody() error {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer func() {
		if buffer != nil {
			bufferPool.Put(buffer)
			buffer = nil
		}
	}()
	if _, err := io.Copy(buffer, r.Raw.Body); err != nil {
		return fmt.Errorf("io.Copy error: %v", err)
	} else {
		r.Body = buffer.String()
		bufferPool.Put(buffer)
		buffer = nil
		return nil
	}
}
