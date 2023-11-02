package text

import (
	"encoding/json"
	"fmt"
	"myapp/function"
	"myapp/redata"
	"strconv"
)

type FuturesTrading struct {
}

// MarginAndLeverage 查看币对杠杆和保证金模式
func (f *FuturesTrading) MarginAndLeverage() {
	formData := map[string]string{}
	url := "http://8.131.92.184:6883/api/v7/user/swap/wallet/leverage-pattern"

	var MarginAndLeverage = redata.MarginAndLeverage{}
	if err := function.PostFormData(url, formData, &MarginAndLeverage); err != nil {
		fmt.Println(err.Error())
	}

	if MarginAndLeverage.Code == 0 {
		//创建 LeveragePattern 结构体并初始化 Result 切片
		leveragePattern := redata.LeveragePattern{
			Result: []redata.LeveragePatternResult{},
		}

		for _, entry := range MarginAndLeverage.Result {
			// 赋值数据到 LeveragePatternResult 结构体
			leveragePattern.Result = append(leveragePattern.Result, redata.LeveragePatternResult{
				ID:       entry.FutureSymbolsID,
				Symbol:   entry.Symbol,
				Leverage: entry.USDTLongLeverage,
				Pattern:  entry.Pattern,
			})
			// 这里不再输出每次迭代的结果
		}
		// 使用JSON编码功能将数据转换为JSON格式
		jsonData, err := json.Marshal(leveragePattern)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 打印JSON格式的数据
		fmt.Println(string(jsonData))
		//fmt.Println(entry.Symbol, "多仓杠杆:", entry.USDTLongLeverage, "空仓杠杆:", entry.USDTShortLeverage)

	} else {
		fmt.Println("查看杠杆失败:", MarginAndLeverage.Message)
	}
}

// LimitOpenLong 限价开多
func (f *FuturesTrading) LimitOpenLong(leverage string, entrustPrice string, volume string) int {

	formData := map[string]string{
		"contractCoinId": "1",
		"side":           "2",
		"type":           "1",
		"entrustPrice":   entrustPrice,
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        "1",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/open"
	var LimitOpenLong = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &LimitOpenLong); err != nil {
		fmt.Println(err)
		return -1
	}

	if LimitOpenLong.Code == 0 {
		fmt.Println("限价开多成功")
		return 0
	} else {
		fmt.Println("限价开多失败:", LimitOpenLong.Message)
		return -1
	}
}

// LimitOpenClose 限价开空
func (f *FuturesTrading) LimitOpenClose(leverage, entrustPrice, volume string) int {

	formData := map[string]string{
		"contractCoinId": "1",
		"side":           "1",
		"type":           "1",
		"entrustPrice":   entrustPrice,
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        "1",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/open"
	var LimitOpenClose = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &LimitOpenClose); err != nil {
		fmt.Println(err)
		return -1
	}

	if LimitOpenClose.Code == 0 {
		fmt.Println("限价开空成功")
		return 0
	} else {
		fmt.Println("限价开空失败:", LimitOpenClose.Message)
		return -2
	}
}

// MarketOpenLong 市价开多
func (f *FuturesTrading) MarketOpenLong(leverage, volume string) int {

	formData := map[string]string{
		"contractCoinId": "1",
		"side":           "2",
		"type":           "0",
		"entrustPrice":   "0",
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        "1",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/open"
	var MarketOpenLong = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketOpenLong); err != nil {
		fmt.Println(err)
		return -1
	}

	if MarketOpenLong.Code == 0 {
		fmt.Println("市价开多成功")
		return 0
	} else {
		fmt.Println("市价开多失败:", MarketOpenLong.Message)
		return -1
	}
}

// MarketOpenClose 市价开空
func (f *FuturesTrading) MarketOpenClose(leverage, volume string) int {

	formData := map[string]string{
		"contractCoinId": "1",
		"side":           "1",
		"type":           "0",
		"entrustPrice":   "0",
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        "1",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/open"
	var MarketOpenClose = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketOpenClose); err != nil {
		fmt.Println(err)
		return -1
	}

	if MarketOpenClose.Code == 0 {
		fmt.Println("市价开空成功")
		return 0
	} else {
		fmt.Println("市价开空失败:", MarketOpenClose.Message)
		return -2
	}
}

// LimitCloseLong 限价平多
func (f *FuturesTrading) LimitCloseLong(leverage, entrustPrice, volume string) {

	formData := map[string]string{
		"contractCoinId": "1",
		"side":           "2",
		"type":           "1",
		"entrustPrice":   entrustPrice,
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        "1",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/close"
	var MarketCloseLong = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketCloseLong); err != nil {
		fmt.Println(err)
		return
	}

	if MarketCloseLong.Code == 0 {
		fmt.Println("限价平多成功")
	} else {
		fmt.Println("限价平多失败:", MarketCloseLong.Message)
	}
}

// LimitCloseShort 限价平空
func (f *FuturesTrading) LimitCloseShort(leverage, entrustPrice, volume string) {

	formData := map[string]string{
		"contractCoinId": "1",
		"side":           "1",
		"type":           "1",
		"entrustPrice":   entrustPrice,
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        "1",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/close"
	var MarketCloseShort = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketCloseShort); err != nil {
		fmt.Println(err)
		return
	}

	if MarketCloseShort.Code == 0 {
		fmt.Println("限价平空成功")
	} else {
		fmt.Println("限价平空失败:", MarketCloseShort.Message)
	}
}

// MarketCloseLong 市价平多
func (f *FuturesTrading) MarketCloseLong(leverage, volume string) {

	formData := map[string]string{
		"contractCoinId": "1",
		"side":           "2",
		"type":           "0",
		"entrustPrice":   "0",
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        "1",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/close"
	var MarketCloseLong = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketCloseLong); err != nil {
		fmt.Println(err)
		return
	}

	if MarketCloseLong.Code == 0 {
		fmt.Println("市价平多成功")
	} else {
		fmt.Println("市价平多失败:", MarketCloseLong.Message)
	}
}

// MarketCloseShort 市价平空
func (f *FuturesTrading) MarketCloseShort(leverage, volume string) {

	formData := map[string]string{
		"contractCoinId": "1",
		"side":           "1",
		"type":           "0",
		"entrustPrice":   "0",
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        "1",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/close"
	var MarketCloseShort = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketCloseShort); err != nil {
		fmt.Println(err)
		return
	}

	if MarketCloseShort.Code == 0 {
		fmt.Println("市价平空成功")
	} else {
		fmt.Println("市价平空失败:", MarketCloseShort.Message)
	}
}

// OpenOrders 当前委托单
func (f *FuturesTrading) OpenOrders() {
	formData := map[string]string{
		"limit":  "100",
		"market": "BTCUSDT",
		"offset": "1",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/pending"
	var OpenOrders = redata.OpenOrders{}
	if err := function.PostFormData(url, formData, &OpenOrders); err != nil {
		fmt.Println(err)
		return
	}

	// 计算amount的总和
	totalAmount := 0
	if OpenOrders.Code == 0 {
		for _, entry := range OpenOrders.Result.Data {
			amount, err := strconv.Atoi(entry.Amount)
			if err != nil {
				fmt.Println("string类型解析为int类型失败:", err)
			}
			totalAmount += amount
			ctime := function.To24h(entry.CTime)
			fmt.Println("创建时间:", ctime, "价格:", entry.Price, "数量:", entry.Amount)
		}
	} else {
		fmt.Println(OpenOrders.Message)
	}
	fmt.Println("订单合计冻结:", totalAmount)
}
