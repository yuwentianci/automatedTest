package function

import (
	"errors"
	"fmt"
	"io"
	"myapp/config"
	"net/http"
	"time"
)

// GetDetails 发送get请求
func GetDetails(url string) ([]byte, error) {
	startTimeMilliseconds := time.Now().UnixNano() / int64(time.Millisecond) // 记录开始时间（毫秒）
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//fmt.Println("创建请求时发生错误:", err)
		return nil, errors.New("创建请求时发生错误:" + err.Error())
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("cookie", "")
	req.Header.Set("Authorization", config.Token)

	// 发送GET请求并获取响应
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
	}(resp.Body)

	// 检查响应的状态码
	if resp.StatusCode != http.StatusOK {
		//fmt.Printf("响应状态码非200 OK: %v\n", resp.Status)
		return nil, errors.New("响应状态码非200 OK: " + resp.Status)
	}

	// 计算响应时间
	endTimeMilliseconds := time.Now().UnixNano() / int64(time.Millisecond)
	elapsedTimeMilliseconds := endTimeMilliseconds - startTimeMilliseconds
	fmt.Printf("GET请求响应时间: %dms\n", elapsedTimeMilliseconds)

	// 读取响应内容
	responseText, err := io.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("读取响应时发生错误:", err)
		return nil, errors.New("读取响应时发生错误: " + resp.Status)
	}

	return responseText, nil
}
