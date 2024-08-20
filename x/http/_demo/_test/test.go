package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func postForm(urlStr string, data url.Values) (string, error) {
	resp, err := http.PostForm(urlStr, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	// 使用 HTTPbin 的 POST 端点
	urlStr := "http://httpbin.org/post"

	// 创建表单数据
	formData := url.Values{
		"name":  {"John Doe"},
		"email": {"johndoe@example.com"},
	}

	// 发送请求
	response, err := postForm(urlStr, formData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印响应
	fmt.Println("Response:", response)
}