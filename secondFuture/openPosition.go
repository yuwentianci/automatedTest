package secondFuture

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"myapp/config/login"
	"myapp/function"
	"sync"
)

type LoginResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  LoginResult `json:"result"`
}

type LoginResult struct {
	Token string `json:"token"`
}

func getToken(email string) (error, string) {
	password := "RmVpZ2UxMjMh"

	// 添加表单字段
	formData := map[string]interface{}{
		"username": email,
		"password": password,
	}

	// 解析JSON响应
	var loginResponse LoginResponse
	if err := function.PostFormData(login.LoginUrl, formData, &loginResponse); err != nil {
		fmt.Println(err.Error())
		return err, ""
	}

	if loginResponse.Code == 0 && loginResponse.Message == "Success" {
		return nil, loginResponse.Result.Token
	}
	return fmt.Errorf("登录失败的邮箱是:%s,失败原因是:%s", email, loginResponse.Message), ""
}
func getTokens(emails []string) (error, []string) {
	var tokens []string
	for _, email := range emails {
		err, token := getToken(email)
		if err != nil {
			log.Println("Error:", err)
			continue // 忽略错误，继续处理下一个 email
		}
		tokens = append(tokens, token)
	}
	return nil, tokens
}

type SafeResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Result  SafeResult `json:"result"`
}

type SafeResult struct {
	EmailOrigin string `json:"email_origin"`
	Token       string `json:"token"`
}

func getSafeToken(emails []string) (error, []string) {
	// 获取 tokens
	err, tokens := getTokens(emails)
	if err != nil {
		return err, nil
	}

	var safeTokens []string // 保存成功处理的 tokens

	// 为每个 token 调用 PostFormDataManyToken
	for _, token := range tokens {
		// 添加表单字段
		formData := map[string]interface{}{
			"two_step_code": 123456,
			"email_code":    123456,
			"phone_code":    123456,
		}

		// 解析 JSON 响应
		var safeResponse SafeResponse
		if err := function.PostFormDataManyToken(login.SafeLoginUrl, formData, &safeResponse, token); err != nil {
			log.Println("Error processing token:", err)
			continue
		}

		// 检查响应
		if safeResponse.Code == 0 && safeResponse.Message == "Success" {
			safeTokens = append(safeTokens, safeResponse.Result.Token) // 保存成功处理的 token
		} else {
			fmt.Println("错误码是:", safeResponse.Code, "错误信息是:", safeResponse.Message)
		}
	}

	// 返回成功处理的 tokens
	return nil, safeTokens
}

type MakeResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"` // 使用 interface{} 以支持 data 字段为 null 或其他类型的情况
}

//func MakePosition(roundId, amount, side int, emails []string) (error, []int) {
//	var orderIds []int
//	err, safeTokens := getSafeToken(emails)
//	if err != nil {
//		return err, nil
//	}
//	URL := fmt.Sprintf("%sasset=USDT&round_id=%d&amount=%d&side=%d", login.MakePositionUrl, roundId, amount, side)
//	for _, token := range safeTokens {
//		responseTest, err := function.GetDetailsManyToken(URL, token)
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		var makeResponse MakeResponse
//		if err := json.Unmarshal(responseTest, &makeResponse); err != nil {
//			return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
//		}
//
//		if makeResponse.Code == 200 {
//			orderIds = append(orderIds, makeResponse.Code)
//		} else {
//			fmt.Println("错误码:", makeResponse.Code, "错误信息:", makeResponse.Msg)
//		}
//	}
//	return nil, orderIds
//}

func MakePosition(roundId, amount, side int, emails []string) (error, []int) {
	var (
		orderIds   []int
		mu         sync.Mutex         // 用于保护对 orderIds 的并发访问
		wg         sync.WaitGroup     // 用于等待所有 goroutine 完成
		resultChan = make(chan int)   // 用于收集成功的订单ID
		errorChan  = make(chan error) // 用于收集错误信息
	)

	// 获取安全的 tokens
	err, safeTokens := getSafeToken(emails)
	if err != nil {
		return err, nil
	}

	URL := fmt.Sprintf("%sasset=USDT&round_id=%d&amount=%d&side=%d", login.MakePositionUrl, roundId, amount, side)

	// 为每个 token 启动 goroutine
	for _, token := range safeTokens {
		wg.Add(1)
		go func(token string) {
			defer wg.Done()

			// 调用接口
			responseTest, err := function.GetDetailsManyToken(URL, token)
			if err != nil {
				errorChan <- err
				return
			}

			var makeResponse MakeResponse
			if err := json.Unmarshal(responseTest, &makeResponse); err != nil {
				errorChan <- errors.New("解析JSON响应时发生错误: " + err.Error())
				return
			}

			// 处理响应
			if makeResponse.Code == 200 {
				resultChan <- makeResponse.Code
			} else {
				fmt.Println("错误码:", makeResponse.Code, "错误信息:", makeResponse.Msg)
			}
		}(token)
	}

	// 处理结果
	go func() {
		wg.Wait()         // 等待所有 goroutine 完成
		close(resultChan) // 关闭 resultChan 表示不再有数据写入
		close(errorChan)  // 关闭 errorChan 表示不再有错误写入
	}()

	for result := range resultChan {
		mu.Lock()
		orderIds = append(orderIds, result)
		mu.Unlock()
	}

	// 检查是否有错误发生
	if len(errorChan) > 0 {
		return <-errorChan, nil
	}

	return nil, orderIds
}
