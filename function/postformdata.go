package function

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
)

// PostFormData Post参数类型为formData
// func PostFormData(url string, formData map[string]string) ([]byte, error) {
func PostFormData(url string, formData map[string]string, target interface{}) error {
	// 准备表单数据
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range formData {
		if err := writer.WriteField(key, value); err != nil {
			//return nil, err
			return err
		}
	}

	// 关闭表单写入器
	if err := writer.Close(); err != nil {
		//return nil, err
		return err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		//return nil, errors.New("创建请求时发生错误:" + err.Error())
		return errors.New("创建请求时发生错误:" + err.Error())
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjUxMzI0NTM0LCJMb2dpblZlcmlmeSI6MSwiVW5pcXVlVG9rZW4iOiIwMjRjODIwMi1jMGYzLTQyOWMtYjkyMC05OTNlNGZlODBlMzkiLCJBZ2VudCI6IndlYiIsImV4cCI6MTY5ODU4NDIyMX0.4MO3tx1_jh81X1iGR-fIEbRNLIkv-t3z5Uk6v6uEp6w")

	// 发送POST请求
	client := &http.Client{}    // 创建一个HTTP客户端
	resp, err := client.Do(req) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		//return nil, errors.New("发送请求时发生错误:" + err.Error())
		return errors.New("发送请求时发生错误:" + err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(req.Body)

	// 检查响应的状态码
	if resp.StatusCode != http.StatusOK {
		//return nil, errors.New("响应状态码非200 OK: " + resp.Status)
		return errors.New("响应状态码非200 OK: " + resp.Status)
	}

	// 读取响应内容
	responseText, err := io.ReadAll(resp.Body)
	if err != nil {
		//return nil, errors.New("读取响应时发生错误: " + resp.Status)
		return errors.New("读取响应时发生错误: " + resp.Status)
	}
	//return responseText, nil
	if err := json.Unmarshal(responseText, target); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error())
	}

	return nil
}
