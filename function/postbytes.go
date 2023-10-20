package function

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func PostBytesDetails(url string, formData *bytes.Buffer, writer string) ([]byte, error) {

	req, err := http.NewRequest("POST", url, formData)
	if err != nil {
		//fmt.Println("创建请求时发生错误:", err)
		return nil, errors.New("创建请求时发生错误:" + err.Error())
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjExNTA0MzQsIkxvZ2luVmVyaWZ5IjoxLCJVbmlxdWVUb2tlbiI6IjBjMThiYzQ4LTYyMTktNDQ5Mi1hOTMxLTg3NTE2YjgxMzA4ZSIsIkFnZW50IjoiYW5kcm9pZCIsImV4cCI6MTY5ODU4NzgwMn0.AjFM7GIFacgIs9VAFTdyXn4aCbU4RHPfLnPik2ZLwBw")

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
