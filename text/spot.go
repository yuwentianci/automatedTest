package text

import (
	"fmt"
	"myapp/function"
	"myapp/redata"
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
func (t *PostSpotCreateData) LimitBuy(market, price string, amount1, amount2 float64) {
	var wg sync.WaitGroup
	counter := 1

	// 启动5个goroutines并发发送限价购买请求
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			step := function.StepValue(amount1)
			// 限价购买从20到23进行数量循环
			for amount := amount1; amount < amount2; amount += step {

				// 添加表单字段
				formData := map[string]string{
					"market": market,
					"side":   "2",
					"amount": strconv.FormatFloat(amount, 'f', 6, 64),
					"price":  price,
				}

				url := "https://www.biconomy.com/api/v1/user/trade/limit"

				// 解析JSON响应
				var responseSpotOrderData = redata.ResponseSpotOrderData{}
				if err := function.PostFormData(url, formData, &responseSpotOrderData); err != nil {
					fmt.Println(err.Error())
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
func (*PostSpotCreateData) MarketBuy() {
	var wg sync.WaitGroup

	// 启动5个goroutines并发发送市价购买请求
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 市价购买从12到14进行购买总额循环
			for amount := 12; amount <= 12; amount++ {

				// 添加表单字段
				formData := map[string]string{
					"market": "BONE_USDT",
					"side":   "2",
					"amount": strconv.Itoa(amount),
				}

				// 创建一个POST请求
				url := "https://www.biconomy.com/api/v1/user/trade/market"

				// 解析JSON响应
				var responseSpotOrderData = redata.ResponseSpotOrderData{}
				if err := function.PostFormData(url, formData, &responseSpotOrderData); err != nil {
					fmt.Println(err.Error())
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
func (t *PostSpotCreateData) LimitSell() {
	var wg sync.WaitGroup

	// 启动5个goroutines并发发送限价出售请求
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 限价出售从10到13进行数量循环
			for amount := 10; amount <= 10; amount++ {

				// 添加表单字段
				formData := map[string]string{
					"market": "BONE_USDT",
					"side":   "1",
					"amount": strconv.Itoa(amount),
					"price":  "1.05",
				}

				// 创建一个POST请求
				url := "https://www.biconomy.com/api/v1/user/trade/limit"

				// 解析JSON响应
				var responseSpotOrderData = redata.ResponseSpotOrderData{}
				if err := function.PostFormData(url, formData, &responseSpotOrderData); err != nil {
					fmt.Println(err.Error())
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
func (*PostSpotCreateData) MarketSell() {
	var wg sync.WaitGroup

	// 启动5个goroutines并发发送市价出售请求
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 市价出售从13到15进行出售数量循环
			for amount := 13; amount <= 16; amount++ {

				// 添加表单字段
				formData := map[string]string{
					"market": "BONE_USDT",
					"side":   "1",
					"amount": strconv.Itoa(amount),
				}

				url := "https://www.biconomy.com/api/v1/user/trade/market"

				// 解析JSON响应
				var responseSpotOrderData = redata.ResponseSpotOrderData{}
				if err := function.PostFormData(url, formData, &responseSpotOrderData); err != nil {
					fmt.Println(err.Error())
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
func (t *PostSpotCreateData) OpenOrder() {
	// 添加表单字段
	formData := map[string]string{
		"limit":  "101",
		"market": "BONE_USDT",
		"offset": "1",
		"side":   "0",
	}

	url := "https://www.biconomy.com/api/v4/user/order/openOrders"

	// 解析JSON响应
	var responseSpotTradeData = redata.ResponseSpotTradeData{}
	if err := function.PostFormData(url, formData, &responseSpotTradeData); err != nil {
		fmt.Println(err.Error())
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
			sideStr := function.SpotSideMap(entry.Side)
			// 将 total 和 totalAmount 转换为字符串
			totalStr := strconv.FormatFloat(total, 'f', 4, 64)
			totalAmountStr := strconv.FormatFloat(totalAmount, 'f', 4, 64)
			fmt.Printf("%s 委托%s %s 数量: %s 委托价格: %s 合计: %s 所有订单总计金额: %s\n",
				entry.Time, sideStr, entry.Market, entry.Amount, entry.Price, totalStr, totalAmountStr)
		}
	} else {
		fmt.Println("没有数据")
	}
}

// OrderHistory 历史委托隐藏其他交易对
func (t *PostSpotCreateData) OrderHistory() {
	// 添加表单字段
	formData := map[string]string{
		"limit":  "101",
		"market": "FIL_USDT",
		"offset": "1",
		"side":   "0",
		"status": "0",
	}

	// 创建一个POST请求
	url := "https://www.biconomy.com/api/v4/user/order/orderHistory"

	// 解析JSON响应
	var responseSpotTradeData = redata.ResponseSpotTradeData{}
	if err := function.PostFormData(url, formData, &responseSpotTradeData); err != nil {
		fmt.Println(err.Error())
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
func (t *PostSpotCreateData) TradeHistory() {
	// 添加表单字段
	formData := map[string]string{
		"limit":  "10",
		"market": "FIL_USDT",
		"offset": "1",
		"side":   "0",
	}

	// 创建一个POST请求
	url := "https://www.biconomy.com/api/v4/user/order/tradeHistory"

	// 解析JSON响应
	var responseSpotTradeData = redata.ResponseSpotTradeData{}
	if err := function.PostFormData(url, formData, &responseSpotTradeData); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}

	if len(responseSpotTradeData.Result.Data) > 0 {
		fmt.Println(len(responseSpotTradeData.Result.Data))

		for _, entry := range responseSpotTradeData.Result.Data {
			roleStr := function.SpotRoleMap(entry.Role)
			sideStr := "未知"
			FeeStr := "0.00"

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
				entry.Time, roleStr, sideStr, entry.Market, entry.Filled, entry.Price, entry.Total, FeeStr)
		}
	} else {
		fmt.Println("没有数据")
	}
}
