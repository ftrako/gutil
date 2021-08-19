package goutils

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var DefaultClient *http.Client

func DoGet(url string) (string, error) {
	return DoGetWithHeader(url, nil)
}

func DoGetWithHeader(url string, header map[string]string) (string, error) {
	b, err := doWithHeader(url, http.MethodGet, "", header)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func DoPost(url, body string) (string, error) {
	return DoPostWithHeader(url, body, nil)
}

func DoPostWithHeader(url, body string, header map[string]string) (string, error) {
	b, err := doWithHeader(url, http.MethodPost, body, header)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func DoPut(url, body string) (string, error) {
	return DoPutWithHeader(url, body, nil)
}

func DoPutWithHeader(url, body string, header map[string]string) (string, error) {
	b, err := doWithHeader(url, http.MethodPut, body, header)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// 通过代理Get请求
func DoProxyGet(urls, proxyIP string) (string, error) {
	proxy, _ := url.Parse(proxyIP)
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
			Proxy:               http.ProxyURL(proxy),
		},
		Timeout: time.Second * 5}

	resp, err := client.Get(urls)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// method 支持GET， POST等方式，应该传常量，而不是直接传字符串，参考http.MethodGet
func doWithHeader(url, method, body string, header map[string]string) ([]byte, error) {
	var bodyReader io.Reader
	if len(body) > 0 {
		bodyReader = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	if header != nil && len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// defer DefaultClient.CloseIdleConnections()

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func init() {
	DefaultClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
		},
		Timeout: time.Second * 5}
}
