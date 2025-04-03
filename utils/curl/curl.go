package curl

import (
	"crypto/tls"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

func Post(url string, param string, headers map[string]string, options map[string]string) (resBody string, respHeader fasthttp.ResponseHeader, err error) {
	return exec("POST", url, param, headers, 30*time.Second, options)
}

func exec(method string, url string, param string, headers map[string]string, duration time.Duration, options map[string]string) (resBody string, respHeader fasthttp.ResponseHeader, err error) {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURI(url)
	// 设置参数
	req.Header.SetMethod(method)
	req.SetBody([]byte(param))

	// 设置头信息
	if headers != nil && len(headers) != 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	client := &fasthttp.Client{
		TLSConfig:          &tls.Config{InsecureSkipVerify: true},
		MaxConnWaitTimeout: duration,
	}
	// 设置代理
	if proxy, ok := options["proxy"]; ok {
		client.Dial = fasthttpproxy.FasthttpSocksDialer(proxy)
	}
	err = client.DoRedirects(req, resp, 10)

	respHeader = resp.Header
	if err != nil {
		return
	}

	resBody = string(resp.Body())

	return
}
