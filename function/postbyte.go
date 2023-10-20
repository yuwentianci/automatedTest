package function

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// PostByteDetails 发送get请求
func PostByteDetails(url string, jsonData []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		//fmt.Println("创建请求时发生错误:", err)
		return nil, errors.New("创建请求时发生错误:" + err.Error())
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjExNjQ4MjgsIkxvZ2luVmVyaWZ5IjowLCJVbmlxdWVUb2tlbiI6ImYxNjczMWQxLTZlMDktNGY4OS04MThiLWM3ZjIxMDc0M2I0NyIsIkFnZW50IjoiIiwiZXhwIjoxNjk3MDI5NjI1fQ.Xr_DKz4Y0HFkQsE7OZiXkt5KbbSGSSss9-hS_f2B2xY")

	// 发送POST请求
	client := http.Client{}     // 创建一个HTTP客户端
	resp, err := client.Do(req) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		//fmt.Println("发送请求时发生错误:", err.Error())
		return nil, errors.New("发送请求时发生错误:" + err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(req.Body)

	// 检查响应的状态码
	if resp.StatusCode != http.StatusOK {
		//fmt.Printf("响应状态码非200 OK: %v\n", resp.Status)
		return nil, errors.New("响应状态码非200 OK: " + resp.Status)
	}

	// 读取响应内容
	responseText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应时发生错误:", err)
	}
	return responseText, nil
}
