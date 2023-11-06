package text

import (
	"encoding/json"
	"errors"
	"fmt"
	"myapp/function"
	"myapp/redata"
	"strconv"
)

type FuturesTrading struct {
}

// MarginAndLeverage 查看币对杠杆和保证金模式
func (f *FuturesTrading) MarginAndLeverage() ([]byte, error) {
	formData := map[string]string{}
	url := "http://8.131.92.184:6883/api/v7/user/swap/wallet/leverage-pattern"

	var MarginAndLeverage = redata.MarginAndLeverage{}
	if err := function.PostFormData(url, formData, &MarginAndLeverage); err != nil {
		return nil, err
	}

	if MarginAndLeverage.Code == 0 {
		//创建 leveragePattern 结构体并初始化 Result 切片
		var leveragePattern []redata.LeveragePattern

		for _, entry := range MarginAndLeverage.Result {
			// 赋值数据到 LeveragePatternResult 结构体
			leveragePattern = append(leveragePattern, redata.LeveragePattern{
				ID:       entry.FutureSymbolsID,
				Symbol:   entry.Symbol,
				Leverage: entry.USDTLongLeverage,
				Pattern:  entry.Pattern,
			})
			// 这里不再输出每次迭代的结果
		}

		//使用JSON编码功能将数据转换为JSON格式
		jsonData, err := json.Marshal(leveragePattern)
		if err != nil {
			return nil, err
		}
		return jsonData, nil
	}
	return nil, errors.New(MarginAndLeverage.Message)
}

// symbolThumb 币对信息
func symbolThumb() (redata.SymbolResult, error) {
	formData := map[string]string{}
	url := "http://8.131.92.184:6883/api/v7/swap/symbol-thumb"

	var SymbolThumb = redata.SymbolThumb{}
	if err := function.PostFormData(url, formData, &SymbolThumb); err != nil {
		return redata.SymbolResult{}, err
	}
	for _, entry := range SymbolThumb.Result {
		return entry, nil
	}
	return redata.SymbolResult{}, nil
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
		fmt.Println("限价开多挂单成功")
		return 0
	} else {
		fmt.Println("限价开多挂单失败:", LimitOpenLong.Message)
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
		fmt.Println("限价开空挂单成功")
		return 0
	} else {
		fmt.Println("限价开空挂单失败:", LimitOpenClose.Message)
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
		fmt.Println("限价平多挂单成功")
	} else {
		fmt.Println("限价平多挂单失败:", MarketCloseLong.Message)
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
		fmt.Println("限价平空挂单成功")
	} else {
		fmt.Println("限价平空挂单失败:", MarketCloseShort.Message)
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

// Positions 当前持仓
func (f *FuturesTrading) Positions() (*[]redata.PositionsDetailsInfo, error) {
	formData := map[string]string{
		"market": "BTCUSDT",
		"side":   "",
	}

	url := "http://8.131.92.184:6883/api/v7/user/swap/order/position"
	// 发送HTTP请求并将响应解析为 Positions 结构
	var Positions = redata.Positions{}
	if err := function.PostFormData(url, formData, &Positions); err != nil {
		return nil, err
	}

	// 检查是否存在仓位
	if len(Positions.Result) > 0 {
		jsonVal, err := json.Marshal(Positions.Result)
		if err != nil {
			return nil, err
		}

		data := make([]redata.PositionsDetailsInfo, 0)
		err = json.Unmarshal(jsonVal, &data)
		if err != nil {
			return nil, err
		}
		// 创建并返回 PositionsInfo 对象
		return &data, nil
	}
	// 当没有仓位时返回 nil
	return nil, nil
}

// FuturesBalance 资产
func (f *FuturesTrading) FuturesBalance() (float64, float64, error) {
	formData := map[string]string{}
	url := "http://8.131.92.184:6883/api/v7/user/swap/wallet/balance"
	var FuturesBalance = redata.FuturesBalance{}
	if err := function.PostFormData(url, formData, &FuturesBalance); err != nil {
		return -1, -1, err
	}
	if FuturesBalance.Code == 0 {
		entry := FuturesBalance.Result
		available := function.FormatFloat(entry.Available)
		freeze := function.FormatFloat(entry.Freeze)
		return available, freeze, nil
	} else {
		return -2, -2, nil
	}
}

// UpdateBalance 预计平仓后的资产数据
func UpdateBalance(closePrice, CloseAmount float64) {
	futureTrades := new(FuturesTrading)

	// 使用辅助函数处理错误，以避免代码重复。
	handleErr := func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	available, _, err := futureTrades.FuturesBalance()
	handleErr(err)

	symbolThumb, err := symbolThumb()
	handleErr(err)

	PositionsInfo, err := futureTrades.Positions()
	handleErr(err)

	factor := 0.0
	updateBalance := 0.0
	LiquidationPrice := 0.0

	maintenanceMarginRate := function.FormatFloat(symbolThumb.MaintenanceMarginRate)
	takerFee := function.FormatFloat(symbolThumb.TakerFee)
	//平仓手续费 (平仓数 * 平仓费率)
	CloseFee := CloseAmount * takerFee

	//检查 PositionsInfo 是否为nil
	if PositionsInfo != nil {
		for i, entry := range *PositionsInfo {
			// 这里解引用 PositionsInfo 指针以获取切片
			openPrice := function.FormatFloat(entry.Price)
			position := function.FormatFloat(entry.Position)
			frozen := function.FormatFloat(entry.Frozen)
			principal := function.FormatFloat(entry.Principal)
			pattern := function.FuturesPatternMap(entry.Pattern)
			side := function.FuturesSideMap(entry.Side)

			// 逐仓爆仓价 爆仓价 = 开仓均价 * [ 1 ± (维持保证金率 + 平仓手续费率 - 仓位保证金 / 持仓总量)]
			if entry.Side == 1 && entry.Pattern == 1 { // 逐仓空仓
				LiquidationPrice = openPrice * (1 - (maintenanceMarginRate + takerFee - principal/position))
				factor = 1
			} else if entry.Side == 2 && entry.Pattern == 1 { // 逐仓多仓
				LiquidationPrice = openPrice * (1 + (maintenanceMarginRate + takerFee - principal/position))
				factor = -1
			} else if entry.Side == 1 && entry.Pattern == 2 { // 全仓空仓
				//全仓 只存在单方向 做空 爆仓价 = 开仓均价 X [ 1 - (维持保证金率 + 平仓手续费率 - （仓位保证金(空)+总余额+其他仓PNL-其他仓手续费-其他仓维持保证金)/持仓总量)]
				LiquidationPrice = openPrice * (1 - (maintenanceMarginRate + takerFee - (principal+available)/position))
				factor = 1
			} else { // 全仓多仓
				//全仓 只存在单方向 做多 爆仓价 = 开仓均价 X [ 1 + (维持保证金率 + 平仓手续费率 - （仓位保证金(多)+总余额+其他仓PNL-其他仓手续费-其他仓维持保证金）/持仓总量)]
				LiquidationPrice = openPrice * (1 + (maintenanceMarginRate + takerFee - (principal+available)/position))
				factor = -1
			}

			//盈亏 (±1) * ((开仓价 - 平仓价) / 开仓价 * 平仓数)
			ClosePNL := factor * ((openPrice - closePrice) / openPrice * CloseAmount)

			// 平仓后返还的保证金 (保证金 * (平仓数 / (持仓数 + 冻结数)))
			CloseMargin := principal * (CloseAmount / (position + frozen))

			// 平仓后剩余的保证金
			margin := principal - CloseMargin

			//当前资产 + 平仓后返还的保证金 - 手续费 + () * ((开仓价 - 平仓价)/ 开仓价 * 平仓数)
			updateBalance = available + CloseMargin - CloseFee + ClosePNL

			fmt.Printf("仓位%d %s %s %s 开仓价:%.2f 持仓数:%g 保证金:%g 爆仓价:%.2f 平仓后预计: 仓位保证金:%g 盈亏:%f 总资产:%f\n", 1+i, entry.Market, side, pattern, openPrice, position, principal, LiquidationPrice, margin, ClosePNL, updateBalance)
		}
	}
}

// OrderBook 铺订单薄
func (f *FuturesTrading) OrderBook(futureTrades *FuturesTrading) {
	startLong := 40000           //开多盘口(高)
	endLong := 35000             //开多峰值(低)
	startClose := endLong + 6000 //开空盘口(低)
	endClose := startLong + 6000 //开空峰值(高)

	for longPrice := endLong; longPrice <= startLong; longPrice += 1000 {
		longPriceStr := strconv.Itoa(longPrice)
		futureTrades.LimitOpenLong("10", longPriceStr, "10000")
	}
	for closePrice := startClose; closePrice <= endClose; closePrice += 1000 {
		closePriceStr := strconv.Itoa(closePrice)
		futureTrades.LimitOpenClose("10", closePriceStr, "10000")
	}
}
