package function

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

const (
	uid = 1293728 // 簿单用户: 1293728 持仓用户: 51324534 拉盘用户: 1281997
)

func userBearer(uid int) string {
	// 簿单用户
	placeOrdersUser := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjExNjQ4MjgsIkxvZ2luVmVyaWZ5IjoxLCJVbmlxdWVUb2tlbiI6ImQ4MmIzOWVlLTNmNWQtNDhlNS04NmQ3LWQzMzI1MWEzNjI0MyIsIkFnZW50Ijoid2ViIiwiZXhwIjoxNzA0NDI0NTE2fQ.yLmI7_xGuILTx61KQHxGqDGH71qYUDlX2_iiJfYHhbs"

	// 持仓用户
	PositionUser := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjUxMzI0NTM0LCJMb2dpblZlcmlmeSI6MSwiVW5pcXVlVG9rZW4iOiJlYmJlNGM2Zi1kNTQ1LTQ0MzAtOWJjOC03YjZkN2M4MjhjMzgiLCJBZ2VudCI6IndlYiIsImV4cCI6MTcwMDEyMjA1N30.7TyI-EC3l_tQ1l3wAxsMef67JlM6NtV47p26sEmk8nY"

	// 拉盘用户
	pullPlateUser := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEyODE5OTcsIkxvZ2luVmVyaWZ5IjoxLCJVbmlxdWVUb2tlbiI6ImM4MzhhZTg5LTY2MTEtNDFmMS1iYmYxLTZhNjBmMTQyMzc3NSIsIkFnZW50Ijoid2ViIiwiZXhwIjoxNzAwMTQzMDUyfQ.fAh8yCV1xYiyLBlwRSTaW7Lqcb9_CDAjyEt8Uhws51w"
	if uid == 1293728 {
		return placeOrdersUser
	} else if uid == 51324534 {
		return PositionUser
	} else if uid == 1281997 {
		return pullPlateUser
	}
	return pullPlateUser
}

// PostFormData Post参数类型为formData
func PostFormData(url string, formData map[string]interface{}, target interface{}) error {
	// 准备表单数据
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range formData {
		if err := writer.WriteField(key, fmt.Sprint(value)); err != nil {
			return errors.New("写入表单字段失败: " + err.Error())
		}
	}

	// 关闭表单写入器
	if err := writer.Close(); err != nil {
		return errors.New("关闭表单写入器失败: " + err.Error())
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return errors.New("创建请求时发生错误:" + err.Error())
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("cookie", "_ga=GA1.2.168874534.1690809633; _ga_1VP9V64RZ9=GS1.1.1690809632.1.1.1690809813.15.0.0; _ga_MKMBT9R5FW=GS1.1.1690809632.1.1.1690809813.15.0.0; PHPSESSID=89a3dk4qe0g7jlqhot103f37bj; think_var=zh-cn")
	//req.Header.Set("Authorization", userBearer(uid))
	req.Header.Set("Authorization", "111100")

	// 发送POST请求
	client := &http.Client{}    // 创建一个HTTP客户端
	resp, err := client.Do(req) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		return errors.New("发送请求时发生错误:" + err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(req.Body)

	// 检查响应的状态码
	if resp.StatusCode != http.StatusOK {
		return errors.New("响应状态码非200 OK: " + resp.Status)
	}

	// 读取响应内容
	responseText, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("读取响应时发生错误: " + resp.Status)
	}

	if err := json.Unmarshal(responseText, target); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error())
	}

	return nil
}
