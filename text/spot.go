package text

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"myapp/function"
	"myapp/redata"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 定义entry.Type和entry.Side常量
const (
	LimitOrderType  = 1
	MarketOrderType = 2
	SellSide        = 1
	BuySide         = 2
)

type PostSpotCreateData struct {
}

// LimitBuy 现货限价购买
func (this *PostSpotCreateData) LimitBuy() {
	var wg sync.WaitGroup
	counter := 1

	// 启动5个goroutines并发发送限价购买请求
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 限价购买从20到23进行数量循环
			for amount := 20; amount <= 24; amount++ {
				// 准备表单数据
				formData := &bytes.Buffer{}
				writer := multipart.NewWriter(formData)

				// 添加表单字段
				formFields := map[string]string{
					"market": "BONE_USDT",
					"side":   "2",
					"amount": strconv.Itoa(amount),
					"price":  "0.6",
				}

				for key, value := range formFields {
					if err := writer.WriteField(key, value); err != nil {
						return
					}
				}

				// 关闭表单写入器
				if err := writer.Close(); err != nil {
					return
				}

				// 创建一个POST请求
				url := "https://www.biconomy.com/api/v1/user/trade/limit"
				responseText, err := function.PostBytesDetails(url, formData, writer.FormDataContentType())
				if err != nil {
					println(err.Error())
				}

				// 解析JSON响应
				var responseSpotOrderData = redata.ResponseSpotOrderData{}

				if err := json.Unmarshal(responseText, &responseSpotOrderData); err != nil {
					fmt.Println("解析JSON响应时发生错误:", err)
					return
				}

				if responseSpotOrderData.Result != nil {
					entry := responseSpotOrderData.Result
					amount, _ := strconv.ParseFloat(entry.Amount, 64)
					price, _ := strconv.ParseFloat(entry.Price, 64)
					total := amount * price
					totalStr := strconv.FormatFloat(total, 'f', 6, 64)
					ctime := time.Unix(int64(entry.Ctime), 0)
					fmt.Println("序号:", counter, ctime.Format("2006-01-02 15:04:05"), "限价委托购买", entry.Market, "数量:", entry.Amount, "委托价格:", entry.Price, "购买合计:", totalStr, "手续费:", entry.DealFee)
				} else {
					fmt.Printf("限价购买失败原因: %s\n", responseSpotOrderData.Message)
				}
				// 递增计数器
				counter++
			}
		}()
	}

	// 等待所有goroutines完成
	wg.Wait()
}

// MarketBuy 现货市价购买
func (this *PostSpotCreateData) MarketBuy() {
	var wg sync.WaitGroup

	// 启动5个goroutines并发发送市价购买请求
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 市价购买从12到14进行购买总额循环
			for amount := 12; amount <= 12; amount++ {
				// 准备表单数据
				formData := &bytes.Buffer{}
				writer := multipart.NewWriter(formData)

				// 添加表单字段
				formFields := map[string]string{
					"market": "BONE_USDT",
					"side":   "2",
					"amount": strconv.Itoa(amount),
				}

				for key, value := range formFields {
					if err := writer.WriteField(key, value); err != nil {
						return
					}
				}

				// 关闭表单写入器
				if err := writer.Close(); err != nil {
					return
				}

				// 创建一个POST请求
				url := "https://www.biconomy.com/api/v1/user/trade/market"
				responseText, err := function.PostBytesDetails(url, formData, writer.FormDataContentType())
				if err != nil {
					println(err.Error())
				}

				// 解析JSON响应
				var responseSpotOrderData = redata.ResponseSpotOrderData{}

				if err := json.Unmarshal(responseText, &responseSpotOrderData); err != nil {
					fmt.Println("解析JSON响应时发生错误:", err)
					return
				}

				if responseSpotOrderData.Result != nil {
					dealMoney, _ := strconv.ParseFloat(responseSpotOrderData.Result.DealMoney, 64)
					dealStock, _ := strconv.ParseFloat(responseSpotOrderData.Result.DealStock, 64)
					tradePrice := dealMoney / dealStock
					ctime := time.Unix(int64(responseSpotOrderData.Result.Ctime), 0)
					fmt.Println(ctime.Format("2006-01-02 15:04:05"), "市价委托购买", responseSpotOrderData.Result.Market, "数量:", responseSpotOrderData.Result.DealStock, "成交价格:", tradePrice, "购买合计:", responseSpotOrderData.Result.DealMoney, "手续费:", responseSpotOrderData.Result.DealFee)
				} else {
					fmt.Printf("amount值为：%d，响应内容codedata: %d，messagedata: %s\n", amount,
						responseSpotOrderData.Code, responseSpotOrderData.Message)
				}
			}
		}()
	}

	// 等待所有goroutines完成
	wg.Wait()
}

// LimitSell 现货限价出售
func (this *PostSpotCreateData) LimitSell() {
	var wg sync.WaitGroup

	// 启动5个goroutines并发发送限价出售请求
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 限价出售从10到13进行数量循环
			for amount := 10; amount <= 13; amount++ {
				// 准备表单数据
				formData := &bytes.Buffer{}
				writer := multipart.NewWriter(formData)

				// 添加表单字段
				formFields := map[string]string{
					"market": "BONE_USDT",
					"side":   "1",
					"amount": strconv.Itoa(amount),
					"price":  "1.05",
				}

				for key, value := range formFields {
					if err := writer.WriteField(key, value); err != nil {
						return
					}
				}

				// 关闭表单写入器
				if err := writer.Close(); err != nil {
					return
				}

				// 创建一个POST请求
				url := "https://www.biconomy.com/api/v1/user/trade/limit"
				responseText, err := function.PostBytesDetails(url, formData, writer.FormDataContentType())
				if err != nil {
					println(err.Error())
				}

				// 解析JSON响应
				var responseSpotOrderData = redata.ResponseSpotOrderData{}

				if err := json.Unmarshal(responseText, &responseSpotOrderData); err != nil {
					fmt.Println("解析JSON响应时发生错误:", err)
					return
				}

				if responseSpotOrderData.Result != nil {
					amount, _ := strconv.ParseFloat(responseSpotOrderData.Result.Amount, 64)
					price, _ := strconv.ParseFloat(responseSpotOrderData.Result.Price, 64)
					total := amount * price
					totalStr := strconv.FormatFloat(total, 'f', 4, 64)
					ctime := time.Unix(int64(responseSpotOrderData.Result.Ctime), 0)
					fmt.Println(ctime.Format("2006-01-02 15:04:05"), "限价委托出售", responseSpotOrderData.Result.Market, "数量:", responseSpotOrderData.Result.Amount, "委托价格:", responseSpotOrderData.Result.Price, "合计:", totalStr, "费用deal_fee:", responseSpotOrderData.Result.DealFee)
				} else {
					fmt.Printf("限价出售失败原因:%s\n", responseSpotOrderData.Message)
				}
			}
		}()
	}

	// 等待所有goroutines完成
	wg.Wait()
}

// MarketSell 现货市价出售
func (this *PostSpotCreateData) MarketSell() {
	var wg sync.WaitGroup

	// 启动5个goroutines并发发送市价出售请求
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 市价出售从13到15进行出售数量循环
			for amount := 13; amount <= 16; amount++ {
				// 准备表单数据
				formData := &bytes.Buffer{}
				writer := multipart.NewWriter(formData)

				// 添加表单字段
				formFields := map[string]string{
					"market": "BONE_USDT",
					"side":   "1",
					"amount": strconv.Itoa(amount),
				}

				for key, value := range formFields {
					if err := writer.WriteField(key, value); err != nil {
						return
					}
				}

				// 关闭表单写入器
				if err := writer.Close(); err != nil {
					return
				}

				// 创建一个POST请求
				req, err := http.NewRequest("POST", "https://www.biconomy.com/api/v1/user/trade/market", formData)
				if err != nil {
					fmt.Println("创建请求时发生错误:", err)
					return
				}

				// 设置请求头
				req.Header.Set("Content-Type", writer.FormDataContentType())
				req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEyODIwMzEsIkxvZ2luVmVyaWZ5IjoxLCJVbmlxdWVUb2tlbiI6IjNmNjAxNzhlLTRkYTEtNDdkZS1hZDY2LWFjN2ExMzRjNGZkNyIsIkFnZW50Ijoid2ViIiwiZXhwIjoxNjk2OTI2OTUwfQ.Q0sHcXCRL4QcQLPYy8jtX7mZS5fe3uzZ-xDeQXmEa9s")

				// 发送POST请求
				client := http.Client{}     // 创建一个HTTP客户端
				resp, err := client.Do(req) // Do 方法发送请求，返回 HTTP 回复
				if err != nil {
					fmt.Println("发送请求时发生错误:", err.Error())
					return
				}
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {

					}
				}(req.Body)

				// 检查响应的状态码
				if resp.StatusCode != http.StatusOK {
					fmt.Printf("响应状态码非200 OK: %v\n", resp.Status)
				}

				// 读取响应内容
				responseText, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("读取响应时发生错误:", err)
					return
				}

				//res := string(responseText)
				//fmt.Println(res)

				// 解析JSON响应
				var responseSpotOrderData = redata.ResponseSpotOrderData{}

				if err := json.Unmarshal(responseText, &responseSpotOrderData); err != nil {
					fmt.Println("解析JSON响应时发生错误:", err)
					return
				}

				if responseSpotOrderData.Result != nil {
					ctime := time.Unix(int64(responseSpotOrderData.Result.Ctime), 0)
					dealMoney, _ := strconv.ParseFloat(responseSpotOrderData.Result.DealMoney, 64)
					amount, _ := strconv.ParseFloat(responseSpotOrderData.Result.Amount, 64)
					tradePrice := dealMoney / amount
					tradePriceStr := strconv.FormatFloat(tradePrice, 'f', 4, 64)
					fmt.Println(ctime.Format("2006-01-02 15:04:05"), "市价委托出售", responseSpotOrderData.Result.Market, "数量:", responseSpotOrderData.Result.Amount, "交易价格:", tradePriceStr, "出售合计:", responseSpotOrderData.Result.DealMoney, "手续费:", responseSpotOrderData.Result.DealFee)
				} else {
					fmt.Printf("amount值为：%d，响应内容codedata: %d，messagedata: %s\n", amount,
						responseSpotOrderData.Code, responseSpotOrderData.Message)
				}
			}
		}()
	}

	// 等待所有goroutines完成
	wg.Wait()
}

// OpenOrder 当前委托隐藏其他交易对
func (this *PostSpotCreateData) OpenOrder() {

	// 准备表单数据
	formData := &bytes.Buffer{}
	writer := multipart.NewWriter(formData)

	// 添加表单字段
	formFields := map[string]string{
		"limit":  "101",
		"market": "BONE_USDT",
		"offset": "1",
		"side":   "0",
	}

	for key, value := range formFields {
		if err := writer.WriteField(key, value); err != nil {
			return
		}
	}

	// 关闭表单写入器
	if err := writer.Close(); err != nil {
		return
	}

	// 创建一个POST请求
	req, err := http.NewRequest("POST", "https://www.biconomy.com/api/v4/user/order/openOrders", formData)
	if err != nil {
		fmt.Println("创建请求时发生错误:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEyODIwMzEsIkxvZ2luVmVyaWZ5IjoxLCJVbmlxdWVUb2tlbiI6IjNmNjAxNzhlLTRkYTEtNDdkZS1hZDY2LWFjN2ExMzRjNGZkNyIsIkFnZW50Ijoid2ViIiwiZXhwIjoxNjk2OTI2OTUwfQ.Q0sHcXCRL4QcQLPYy8jtX7mZS5fe3uzZ-xDeQXmEa9s")

	// 发送POST请求
	client := http.Client{}     // 创建一个HTTP客户端
	resp, err := client.Do(req) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		fmt.Println("发送请求时发生错误:", err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)

	// 检查响应的状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("响应状态码非200 OK: %v\n", resp.Status)
	}

	// 读取响应内容
	responseText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应时发生错误:", err)
		return
	}

	//res := string(responseText)
	//fmt.Println(res)

	// 解析JSON响应
	var responseSpotTradeData = redata.ResponseSpotTradeData{}

	if err := json.Unmarshal(responseText, &responseSpotTradeData); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}

	if len(responseSpotTradeData.Result.Data) > 0 {

		fmt.Println(len(responseSpotTradeData.Result.Data))
		// 输出result中的data的个别数组的个别数据
		//for i := 0; i < 10 && i < len(responseSpotTradeData.Result.Data); i++ {
		//	firstEntry := responseSpotTradeData.Result.Data[i] // 输出result中指定data的数据
		//	fmt.Println("Market:", firstEntry.Market)
		//}

		// 计算data数组的amount总和
		totalAmount := 0.0
		// 输出result中的data的所有数组的个别数据
		for _, entry := range responseSpotTradeData.Result.Data {
			amount, _ := strconv.ParseFloat(entry.Amount, 64)
			price, _ := strconv.ParseFloat(entry.Price, 64)
			total := float64(amount * price)
			totalAmount += total
			side := "未知"
			if entry.Side == 1 {
				side = "卖出"
			} else if entry.Side == 2 {
				side = "买入"
			}
			// 将 total 和 totalAmount 转换为字符串
			totalStr := strconv.FormatFloat(total, 'f', 4, 64)
			totalAmountStr := strconv.FormatFloat(totalAmount, 'f', 4, 64)
			fmt.Printf("%s 委托%s %s 数量: %s 委托价格: %s 合计: %s 所有订单总计金额: %s\n",
				entry.Time, side, entry.Market, entry.Amount, entry.Price, totalStr, totalAmountStr)
		}
	} else {
		fmt.Println("没有数据")
	}
}

// OrderHistory 历史委托隐藏其他交易对
func (this *PostSpotCreateData) OrderHistory() {

	// 准备表单数据
	formData := &bytes.Buffer{}
	writer := multipart.NewWriter(formData)

	// 添加表单字段
	formFields := map[string]string{
		"limit":  "101",
		"market": "FIL_USDT",
		"offset": "1",
		"side":   "0",
		"status": "0",
	}

	for key, value := range formFields {
		if err := writer.WriteField(key, value); err != nil {
			return
		}
	}

	// 关闭表单写入器
	if err := writer.Close(); err != nil {
		return
	}

	// 创建一个POST请求
	req, err := http.NewRequest("POST", "https://www.biconomy.com/api/v4/user/order/orderHistory", formData)
	if err != nil {
		fmt.Println("创建请求时发生错误:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEyODIwMzEsIkxvZ2luVmVyaWZ5IjoxLCJVbmlxdWVUb2tlbiI6IjNmNjAxNzhlLTRkYTEtNDdkZS1hZDY2LWFjN2ExMzRjNGZkNyIsIkFnZW50Ijoid2ViIiwiZXhwIjoxNjk2OTI2OTUwfQ.Q0sHcXCRL4QcQLPYy8jtX7mZS5fe3uzZ-xDeQXmEa9s")

	// 发送POST请求
	client := http.Client{}     // 创建一个HTTP客户端
	resp, err := client.Do(req) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		fmt.Println("发送请求时发生错误:", err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)

	// 检查响应的状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("响应状态码非200 OK: %v\n", resp.Status)
	}

	// 读取响应内容
	responseText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应时发生错误:", err)
		return
	}

	// 解析JSON响应
	var responseSpotTradeData = redata.ResponseSpotTradeData{}

	if err := json.Unmarshal(responseText, &responseSpotTradeData); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}

	if len(responseSpotTradeData.Result.Data) > 0 {
		fmt.Println(len(responseSpotTradeData.Result.Data))

		for _, entry := range responseSpotTradeData.Result.Data {
			amount, _ := strconv.ParseFloat(entry.Filled, 64)
			price, _ := strconv.ParseFloat(entry.Price, 64)
			dealMoney, _ := strconv.ParseFloat(entry.DealMoney, 64)

			tradePrice := price
			typeStr := "未知"
			sideStr := "未知"
			FeeStr := "0.00"

			if entry.Status == 3 {
				fmt.Println("订单已取消")
				continue
			}

			//限价买入：委托价格：price 成交价格：deal_money /  filled  成交数量：filled  合计：deal_money
			//市价买入：委托价格：price 成交价格：price 成交数量：filled 合计：deal_money
			//限价出售：委托价格：price 成交价格：deal_money / filled  成交数量：filled  合计：deal_money
			//市价出售：委托价格：price 成交价格：price 成交数量：filled  合计：deal_money
			if entry.Type == LimitOrderType {
				typeStr = "限价"
				tradePrice = dealMoney / amount
			} else if entry.Type == MarketOrderType {
				typeStr = "市价"
			}

			if entry.Side == SellSide {
				sideStr = "卖出"
				FeeStr = entry.DealFee + "USDT"
			} else if entry.Side == BuySide {
				sideStr = "买入"
				parts := strings.Split(entry.Market, "_")
				if len(parts) > 0 {
					result := parts[0]
					FeeStr = entry.DealFee + result
				}
			}

			// 格式化 tradePrice 为字符串，保留四位小数，不进行四舍五入
			tradePriceStr := fmt.Sprintf("%.4f", tradePrice)

			fmt.Printf("%s %s委托%s %s 委托数量: %s 委托价格: %s 成交数量: %s 成交价格: %s 实际花费: %s 手续费: %s\n",
				entry.Time, typeStr, sideStr, entry.Market, entry.Amount, entry.Price, entry.Filled, tradePriceStr, entry.DealMoney, FeeStr)
		}
	} else {
		fmt.Println("没有数据")
	}
}

// TradeHistory 历史成交隐藏其他交易对
func (this *PostSpotCreateData) TradeHistory() {

	// 准备表单数据
	formData := &bytes.Buffer{}
	writer := multipart.NewWriter(formData)

	// 添加表单字段
	formFields := map[string]string{
		"limit":  "10",
		"market": "FIL_USDT",
		"offset": "1",
		"side":   "0",
	}

	for key, value := range formFields {
		if err := writer.WriteField(key, value); err != nil {
			return
		}
	}

	// 关闭表单写入器
	if err := writer.Close(); err != nil {
		return
	}

	// 创建一个POST请求
	url := "https://www.biconomy.com/api/v4/user/order/tradeHistory"
	req, err := http.NewRequest("POST", url, formData)
	if err != nil {
		fmt.Println("创建请求时发生错误:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEyODIwMzEsIkxvZ2luVmVyaWZ5IjoxLCJVbmlxdWVUb2tlbiI6IjNmNjAxNzhlLTRkYTEtNDdkZS1hZDY2LWFjN2ExMzRjNGZkNyIsIkFnZW50Ijoid2ViIiwiZXhwIjoxNjk2OTI2OTUwfQ.Q0sHcXCRL4QcQLPYy8jtX7mZS5fe3uzZ-xDeQXmEa9s")

	// 发送POST请求
	client := http.Client{}     // 创建一个HTTP客户端
	resp, err := client.Do(req) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		fmt.Println("发送请求时发生错误:", err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)

	// 检查响应的状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("响应状态码非200 OK: %v\n", resp.Status)
	}

	// 读取响应内容
	responseText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应时发生错误:", err)
		return
	}

	//res := string(responseText)
	//fmt.Println(res)

	// 解析JSON响应
	var responseSpotTradeData = redata.ResponseSpotTradeData{}

	if err := json.Unmarshal(responseText, &responseSpotTradeData); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}

	if len(responseSpotTradeData.Result.Data) > 0 {
		fmt.Println(len(responseSpotTradeData.Result.Data))

		for _, entry := range responseSpotTradeData.Result.Data {
			RoleStr := "未知"
			sideStr := "未知"
			FeeStr := "0.00"

			if entry.Role == "1" {
				RoleStr = "Taker"
			} else if entry.Role == "2" {
				RoleStr = "Maker"
			}

			switch entry.Side {
			case 1:
				sideStr = "卖出"
				FeeStr = entry.DealFee + "USDT"
			case 2:
				sideStr = "买入"
				parts := strings.Split(entry.Market, "_")
				if len(parts) > 0 {
					result := parts[0]
					FeeStr = entry.DealFee + result
				}
			}

			fmt.Printf("%s %s委托%s %s 数量: %s 成交价格: %s 合计: %s 手续费: %s\n",
				entry.Time, RoleStr, sideStr, entry.Market, entry.Filled, entry.Price, entry.Total, FeeStr)
		}
	} else {
		fmt.Println("没有数据")
	}
}
