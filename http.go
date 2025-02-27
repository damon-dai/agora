package agora

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

func HttpGet(baseUrl string, headers map[string]string, params map[string]interface{}, timeout time.Duration) (string, error) {
	// url校验
	_, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	if timeout == 0 {
		timeout = 5 * time.Second // 默认超时时间 5 秒
	}

	// 构建 URL 带参数
	query := url.Values{}
	for key, value := range params {
		query.Set(key, fmt.Sprintf("%v", value)) // 转换参数为字符串
	}
	fullUrl := baseUrl
	if len(query) > 0 {
		fullUrl = fmt.Sprintf("%s?%s", baseUrl, query.Encode())
	}

	// 创建 fasthttp 请求
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(fullUrl)
	req.Header.SetMethod(fasthttp.MethodGet)

	// 添加自定义头部
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 创建 fasthttp 响应
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 发送请求
	client := &fasthttp.Client{}
	if err = client.DoTimeout(req, resp, timeout); err != nil {
		return "", err
	}

	// 获取响应
	body := string(resp.Body())
	return body, nil

}

func HttpRequest(baseUrl, method string, headers map[string]string, params map[string]interface{}, timeout time.Duration) (string, error) {
	// url校验
	_, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	if timeout == 0 {
		timeout = 5 * time.Second // 默认超时时间 5 秒
	}

	// 构建 URL 带参数
	requestBody, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	// 创建 fasthttp 请求
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(baseUrl)
	req.SetBody(requestBody)
	req.Header.SetMethod(method)

	// 添加自定义头部
	req.Header.SetContentType("application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 创建 fasthttp 响应
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 发送请求
	client := &fasthttp.Client{}
	if err = client.DoTimeout(req, resp, timeout); err != nil {
		return "", err
	}

	// 获取响应
	body := string(resp.Body())
	return body, nil
}
