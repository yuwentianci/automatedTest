package function

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"myapp/config"
	"net/http"
)

// PostByteDetails 发送Post请求
func PostByteDetails(url string, jsonData []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		//fmt.Println("创建请求时发生错误:", err)
		return nil, errors.New("创建请求时发生错误:" + err.Error())
	}

	// 设置请求头
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("Authorization", config.Token)
	//req.Header.Set("Token", config.FutureBackendToken)

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

// PostByteDetailsComplete 发送Post请求
func PostByteDetailsComplete(url string, rawData, target interface{}) error {
	// 将 JSON 数据序列化为字节数组
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		return errors.New("JSON 编码失败:" + err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		//fmt.Println("创建请求时发生错误:", err)
		return errors.New("创建请求时发生错误:" + err.Error())
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", config.Token)
	req.Header.Set("token", config.Token)
	// 发送POST请求
	client := http.Client{}     // 创建一个HTTP客户端
	resp, err := client.Do(req) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		//fmt.Println("发送请求时发生错误:", err.Error())
		return errors.New("发送请求时发生错误:" + err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(req.Body)

	// 检查响应的状态码
	if resp.StatusCode != http.StatusOK {
		//fmt.Printf("响应状态码非200 OK: %v\n", resp.Status)
		return errors.New("响应状态码非200 OK: " + resp.Status)
	}

	// 读取响应内容
	responseText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应时发生错误:", err)
	}

	if err := json.Unmarshal(responseText, target); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error())
	}

	return nil
}
