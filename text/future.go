package text

import (
	"encoding/json"
	"errors"
	"fmt"
	"myapp/function"
	"myapp/redata"
	"strconv"
)

const (
	contractCoinId = 1
	Limit          = 1
	Market         = 0
	ShortSide      = 1
	LongSide       = 2
	leverage       = 5.0
	pattern        = 1
)

type FuturesTrading struct {
}

// MarginAndLeverage 查看币对杠杆和保证金模式
func (f *FuturesTrading) MarginAndLeverage() ([]redata.LeveragePattern, error) {
	formData := map[string]interface{}{}
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
		return leveragePattern, nil
	}
	return nil, errors.New(MarginAndLeverage.Message)
}

// symbolThumb 币对信息
func symbolThumb() ([]redata.SymbolResult, error) {
	formData := map[string]interface{}{}
	url := "http://8.131.92.184:6883/api/v7/swap/symbol-thumb"

	var SymbolThumb = redata.SymbolThumb{}
	if err := function.PostFormData(url, formData, &SymbolThumb); err != nil {
		return []redata.SymbolResult{}, err
	}

	var symbolResults []redata.SymbolResult
	for _, entry := range SymbolThumb.Result {
		symbolResults = append(symbolResults, entry)
	}

	return symbolResults, nil
}

// LimitOpenLong 限价开多
func (f *FuturesTrading) LimitOpenLong(leverage string, entrustPrice string, volume string) (int, float64) {
	// 获取当前时间戳
	currentTime := function.NowToUnix()

	formData := map[string]interface{}{
		"contractCoinId": contractCoinId,
		"side":           LongSide,
		"type":           Limit,
		"entrustPrice":   entrustPrice,
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        pattern,
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/open"
	var LimitOpenLong = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &LimitOpenLong); err != nil {
		fmt.Println("限价开多:", err)
		return 1, currentTime
	}

	if LimitOpenLong.Code == 0 {
		fmt.Println("限价开多挂单成功")
		return 0, currentTime
	} else {
		fmt.Println("限价开多挂单失败:", LimitOpenLong.Message)
		return 1, currentTime
	}
}

// LimitOpenClose 限价开空
func (f *FuturesTrading) LimitOpenClose(leverage, entrustPrice, volume string) (int, float64) {
	// 获取当前时间戳
	currentTime := function.NowToUnix()

	formData := map[string]interface{}{
		"contractCoinId": "1",
		"side":           ShortSide,
		"type":           Limit,
		"entrustPrice":   entrustPrice,
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        pattern,
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/open"
	var LimitOpenClose = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &LimitOpenClose); err != nil {
		fmt.Println("限价开空:", err)
		return 1, currentTime
	}

	if LimitOpenClose.Code == 0 {
		fmt.Println("限价开空挂单成功")
		return 0, currentTime
	} else {
		fmt.Println("限价开空挂单失败:", LimitOpenClose.Message)
		return 1, currentTime
	}
}

// MarketOpenLong 市价开多
func (f *FuturesTrading) MarketOpenLong(leverage, volume string) (int, float64) {
	// 获取当前时间戳
	currentTime := function.NowToUnix()

	formData := map[string]interface{}{
		"contractCoinId": "1",
		"side":           LongSide,
		"type":           Market,
		"entrustPrice":   "0",
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        pattern,
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/open"
	var MarketOpenLong = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketOpenLong); err != nil {
		fmt.Println("市价开多:", err)
		return 1, currentTime
	}

	if MarketOpenLong.Code == 0 {
		fmt.Println("市价开多成功")
		return 0, currentTime
	} else {
		fmt.Println("市价开多失败:", MarketOpenLong.Message)
		return 1, currentTime
	}
}

// MarketOpenClose 市价开空
func (f *FuturesTrading) MarketOpenClose(leverage, volume string) (int, float64) {
	// 获取当前时间戳
	currentTime := function.NowToUnix()

	formData := map[string]interface{}{
		"contractCoinId": "1",
		"side":           ShortSide,
		"type":           Market,
		"entrustPrice":   "0",
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        pattern,
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/open"
	var MarketOpenClose = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketOpenClose); err != nil {
		fmt.Println("市价开空:", err)
		return 1, currentTime
	}

	if MarketOpenClose.Code == 0 {
		fmt.Println("市价开空成功")
		return 0, currentTime
	} else {
		fmt.Println("市价开空失败:", MarketOpenClose.Message)
		return 1, currentTime
	}
}

// LimitCloseLong 限价平多
func (f *FuturesTrading) LimitCloseLong(leverage, entrustPrice, volume string) (int, float64) {
	// 获取当前时间戳
	currentTime := function.NowToUnix()

	formData := map[string]interface{}{
		"contractCoinId": "1",
		"side":           LongSide,
		"type":           Limit,
		"entrustPrice":   entrustPrice,
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        pattern,
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/close"
	var MarketCloseLong = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketCloseLong); err != nil {
		fmt.Println("限价平多:", err)
		return 1, currentTime
	}

	if MarketCloseLong.Code == 0 {
		fmt.Println("限价平多挂单成功")
		return 0, currentTime
	} else {
		fmt.Println("限价平多挂单失败:", MarketCloseLong.Message)
		return 1, currentTime
	}
}

// LimitCloseShort 限价平空
func (f *FuturesTrading) LimitCloseShort(leverage, entrustPrice, volume string) (int, float64) {
	// 获取当前时间戳
	currentTime := function.NowToUnix()

	formData := map[string]interface{}{
		"contractCoinId": "1",
		"side":           ShortSide,
		"type":           Limit,
		"entrustPrice":   entrustPrice,
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        pattern,
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/close"
	var MarketCloseShort = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketCloseShort); err != nil {
		fmt.Println("限价平空:", err)
		return 1, currentTime
	}

	if MarketCloseShort.Code == 0 {
		fmt.Println("限价平空挂单成功")
		return 0, currentTime
	} else {
		fmt.Println("限价平空挂单失败:", MarketCloseShort.Message)
		return 1, currentTime
	}
}

// MarketCloseLong 市价平多
func (f *FuturesTrading) MarketCloseLong(leverage, volume string) (int, float64) {
	// 获取当前时间戳
	currentTime := function.NowToUnix()

	formData := map[string]interface{}{
		"contractCoinId": "1",
		"side":           LongSide,
		"type":           Market,
		"entrustPrice":   "0",
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        pattern,
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/close"
	var MarketCloseLong = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketCloseLong); err != nil {
		fmt.Println("市价平多:", err)
		return 1, currentTime
	}

	if MarketCloseLong.Code == 0 {
		fmt.Println("市价平多成功")
		return 0, currentTime
	} else {
		fmt.Println("市价平多失败:", MarketCloseLong.Message)
		return 1, currentTime
	}
}

// MarketCloseShort 市价平空
func (f *FuturesTrading) MarketCloseShort(leverage, volume string) (int, float64) {
	// 获取当前时间戳
	currentTime := function.NowToUnix()

	formData := map[string]interface{}{
		"contractCoinId": "1",
		"side":           ShortSide,
		"type":           Market,
		"entrustPrice":   "0",
		"triggerPrice":   "",
		"leverage":       leverage,
		"volume":         volume,
		"pattern":        pattern,
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/close"
	var MarketCloseShort = redata.OrdersPositions{}
	if err := function.PostFormData(url, formData, &MarketCloseShort); err != nil {
		fmt.Println("市价平空:", err)
		return 1, currentTime
	}

	if MarketCloseShort.Code == 0 {
		fmt.Println("市价平空成功")
		return 0, currentTime
	} else {
		fmt.Println("市价平空失败:", MarketCloseShort.Message)
		return 1, currentTime
	}
}

// OpenOrders 当前委托单
func (f *FuturesTrading) OpenOrders() {
	formData := map[string]interface{}{
		"limit":  "100",
		"market": "BTCUSDT",
		"offset": "1",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/order/pending"
	var OpenOrders = redata.OpenOrders{}
	if err := function.PostFormData(url, formData, &OpenOrders); err != nil {
		fmt.Println("当前委托:", err)
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
func (f *FuturesTrading) Positions() ([]*redata.PositionsDetailsInfo, error) {
	formData := map[string]interface{}{
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
			return nil, errors.New("解析JSON响应时发生错误:" + err.Error())
		}

		data := make([]*redata.PositionsDetailsInfo, 0)
		err = json.Unmarshal(jsonVal, &data)
		if err != nil {
			return nil, errors.New("解析JSON响应时发生错误:" + err.Error())
		}
		// 创建并返回 PositionsDetailsInfo 对象
		return data, nil
	}
	// 当没有仓位时返回 nil
	return nil, errors.New("当前" + "无持仓")
}

// TradeHistory 交易历史记录
func (f *FuturesTrading) TradeHistory() ([]redata.TradeHistoryNeed, error) {
	formData := map[string]interface{}{
		"market": "BTCUSDT",
		"offset": "1",
		"limit":  "100",
	}
	url := "http://8.131.92.184:6883/api/v7/user/swap/trade/history"
	// 发送HTTP请求并将响应解析为 TradeHistory 结构
	var TradeHistory = redata.TradeHistory{}
	if err := function.PostFormData(url, formData, &TradeHistory); err != nil {
		return nil, err
	}

	// 检查是否存在交易历史记录
	if len(TradeHistory.Result.Data) > 0 {
		jsonVal, err := json.Marshal(TradeHistory.Result.Data)
		if err != nil {
			return nil, errors.New("解析JSON响应时发生错误:" + err.Error())
		}

		data := make([]redata.TradeHistoryNeed, 0)
		err = json.Unmarshal(jsonVal, &data)
		if err != nil {
			return nil, errors.New("解析JSON响应时发生错误:" + err.Error())
		}
		return data, nil
	}
	// 当没有成交记录时返回 nil
	return nil, errors.New("此时没有订单成交")
}

// FuturesBalance 资产
func (f *FuturesTrading) FuturesBalance() (float64, float64, error) {
	formData := map[string]interface{}{}
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
		return 1, 1, errors.New(fmt.Sprintf("获取资产失败: %+v", FuturesBalance.Result))
	}
}

// GetBeforeOpenPosition 获取开仓前指定的仓位信息
func GetBeforeOpenPosition(sides int, symbol string) []redata.PositionsDetailsInfo {
	dd := FuturesTrading{}
	// 获取仓位信息
	PositionsInfo, err := dd.Positions()
	if err != nil {
		fmt.Println(err)
	}

	// 寻找匹配的仓位信息
	currentPosition := new(redata.PositionsDetailsInfo)
	for _, item := range PositionsInfo {
		if item.Side == sides && item.Market == symbol {
			currentPosition = item
			//continue
			break // 假设你想在找到匹配后停止搜索
		}
	}

	if currentPosition == nil {
		fmt.Println("未持仓")
		return nil
	} else {
		openPrice := function.FormatFloat(currentPosition.Price)
		position := function.FormatFloat(currentPosition.Position)
		principal := function.FormatFloat(currentPosition.Principal)
		side := function.FuturesSideMap(currentPosition.Side)
		pattern := function.FuturesPatternMap(currentPosition.Pattern)
		fmt.Printf("开仓前仓位信息: %s %s %s 持仓数:%g 仓位价:%.2f  保证金:%g\n", currentPosition.Market, side, pattern, position, openPrice, principal)
	}
	currentPositionSlice := []redata.PositionsDetailsInfo{*currentPosition}
	return currentPositionSlice
}

// GetAfterPosition 获取开仓后指定的仓位信息
func GetAfterPosition(sides int, symbol string) []redata.PositionsDetailsInfo {
	dd := FuturesTrading{}
	// 获取仓位信息
	PositionsInfo, err := dd.Positions()
	if err != nil {
		fmt.Println(err)
	}

	// 寻找匹配的仓位信息
	currentPosition := new(redata.PositionsDetailsInfo)
	for _, item := range PositionsInfo {
		if item.Side == sides && item.Market == symbol {
			currentPosition = item
			//continue
			break // 假设你想在找到匹配后停止搜索
		}
	}

	if currentPosition == nil {
		fmt.Println("未持仓")
		return nil
	} else {
		openPrice := function.FormatFloat(currentPosition.Price)
		position := function.FormatFloat(currentPosition.Position)
		principal := function.FormatFloat(currentPosition.Principal)
		side := function.FuturesSideMap(currentPosition.Side)
		pattern := function.FuturesPatternMap(currentPosition.Pattern)
		fmt.Printf("开仓后实际仓位信息: %s %s %s 持仓数:%g 仓位价:%.2f  保证金:%g\n", currentPosition.Market, side, pattern, position, openPrice, principal)
	}
	currentPositionSlice := []redata.PositionsDetailsInfo{*currentPosition}
	return currentPositionSlice
}

// OpenEstimatedBalance 预计开仓后的仓位信息和资产数据
func OpenEstimatedBalance(PositionsInfo *[]redata.PositionsDetailsInfo, available, newOpenPrice, newOpenAmount float64) {
	//futureTrades := new(FuturesTrading)

	// 使用辅助函数处理错误，以避免代码重复。
	handleErr := func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	// 获取币对费率
	symbolThumb, err := symbolThumb()
	handleErr(err)

	LiquidationPrice := 0.0
	leverage := leverage

	//检查 PositionsInfo 是否为nil
	if PositionsInfo != nil {
		for _, positions := range *PositionsInfo {
			var fee redata.SymbolResult
			for _, fees := range symbolThumb {
				if fees.Symbol == positions.Market {
					fee = fees
				}
			}
			// 这里解引用 PositionsInfo 指针以获取切片
			openPrice := function.FormatFloat(positions.Price)
			position := function.FormatFloat(positions.Position)
			principal := function.FormatFloat(positions.Principal)
			side := function.FuturesSideMap(positions.Side)
			pattern := function.FuturesPatternMap(positions.Pattern)
			maintenanceMarginRate := function.FormatFloat(fee.MaintenanceMarginRate)
			takerFee := function.FormatFloat(fee.TakerFee)

			//新开仓手续费 (开仓数 * 开仓费率)
			newOpenFee := newOpenAmount * takerFee

			// 新开仓保证金 = 新开仓数量 / 杠杆
			newOpenMargin := newOpenAmount / leverage

			// 新仓位保证金 = 当前仓位保证金 + 新开仓保证金 - 新开仓手续费
			newPrincipal := principal + newOpenMargin - newOpenFee

			// 新资产余额 = 当前资产 - 新开仓保证金
			newBalance := available - newOpenMargin

			// 新开仓均价 = (开仓价 * 开仓数 + 新开仓均价 * 新开仓数) / (开仓数 + 新开仓数)
			openAveragePrice := (openPrice*position + newOpenPrice*newOpenAmount) / (position + newOpenAmount)

			// 新开仓总量 = 仓位数 + 新开仓数
			newPosition := position + newOpenAmount

			if positions.Side == 1 && positions.Pattern == 1 { // 逐仓空仓
				//  逐仓做空 爆仓价 = 新开仓均价 * [ 1 - (维持保证金率 + 平仓手续费率 - 仓位保证金 / 持仓总量)]
				LiquidationPrice = openAveragePrice * (1 - (maintenanceMarginRate + takerFee - newPrincipal/newPosition))
			} else if positions.Side == 2 && positions.Pattern == 1 { // 逐仓多仓
				// 逐仓做多 爆仓价 = 新开仓均价 * [ 1 + (维持保证金率 + 平仓手续费率 - 仓位保证金 / 持仓总量)]
				LiquidationPrice = openAveragePrice * (1 + (maintenanceMarginRate + takerFee - newPrincipal/newPosition))
			} else if positions.Side == 1 && positions.Pattern == 2 { // 全仓空仓
				//全仓 只存在单方向 做空 爆仓价 = 新开仓均价 X [ 1 - (维持保证金率 + 平仓手续费率 - （仓位保证金(空)+总余额+其他仓PNL-其他仓手续费-其他仓维持保证金)/持仓总量)]
				LiquidationPrice = openAveragePrice * (1 - (maintenanceMarginRate + takerFee - (newPrincipal+newBalance)/newPosition))
			} else if positions.Side == 2 && positions.Pattern == 1 { // 全仓多仓
				//全仓 只存在单方向 做多 爆仓价 = 新开仓均价 X [ 1 + (维持保证金率 + 平仓手续费率 - （仓位保证金(多)+总余额+其他仓PNL-其他仓手续费-其他仓维持保证金）/持仓总量)]
				LiquidationPrice = openAveragePrice * (1 + (maintenanceMarginRate + takerFee - (newPrincipal+newBalance)/newPosition))
			}
			fmt.Printf("开仓后预计仓位信息: %s %s %s 新持仓数:%.0f 新开仓价:%.2f 新爆仓价:%.2f 新保证金:%g 预计可用资产:%f\n", positions.Market, side, pattern, newPosition, openAveragePrice, LiquidationPrice, newPrincipal, newBalance)
		}
	} else {
		fmt.Println("当前无持仓")
	}
}

// CloseEstimatedBalance 预计平仓后的仓位信息和资产数据
func CloseEstimatedBalance(closePrice, CloseAmount float64) {
	futureTrades := new(FuturesTrading)

	// 使用辅助函数处理错误，以避免代码重复。
	handleErr := func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	// 获取资产
	available, _, err := futureTrades.FuturesBalance()
	handleErr(err)

	// 获取币对费率
	symbolThumb, err := symbolThumb()
	handleErr(err)

	// 获取仓位信息
	PositionsInfo, err := futureTrades.Positions()
	handleErr(err)

	factor := 0.0
	updateBalance := 0.0
	LiquidationPrice := 0.0

	//检查 PositionsInfo 是否为nil
	if PositionsInfo != nil {
		for i, positions := range PositionsInfo {
			var fee redata.SymbolResult
			for _, fees := range symbolThumb {
				if fees.Symbol == positions.Market {
					fee = fees
				}
			}
			// 这里解引用 PositionsInfo 指针以获取切片
			openPrice := function.FormatFloat(positions.Price)
			position := function.FormatFloat(positions.Position)
			frozen := function.FormatFloat(positions.Frozen)
			principal := function.FormatFloat(positions.Principal)
			pattern := function.FuturesPatternMap(positions.Pattern)
			side := function.FuturesSideMap(positions.Side)

			maintenanceMarginRate := function.FormatFloat(fee.MaintenanceMarginRate)
			takerFee := function.FormatFloat(fee.TakerFee)

			//平仓手续费 (平仓数 * 平仓费率)
			CloseFee := CloseAmount * takerFee

			// 逐仓爆仓价 爆仓价 = 开仓均价 * [ 1 ± (维持保证金率 + 平仓手续费率 - 仓位保证金 / 持仓总量)]
			if positions.Side == 1 && positions.Pattern == 1 { // 逐仓空仓
				LiquidationPrice = openPrice * (1 - (maintenanceMarginRate + takerFee - principal/position))
				factor = 1
			} else if positions.Side == 2 && positions.Pattern == 1 { // 逐仓多仓
				LiquidationPrice = openPrice * (1 + (maintenanceMarginRate + takerFee - principal/position))
				factor = -1
			} else if positions.Side == 1 && positions.Pattern == 2 { // 全仓空仓
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

			//当前资产 = 平仓后返还的保证金 - 手续费 + () * ((开仓价 - 平仓价)/ 开仓价 * 平仓数)
			updateBalance = available + CloseMargin - CloseFee + ClosePNL

			fmt.Printf("仓位%d %s %s %s 开仓价:%.2f 持仓数:%g 保证金:%g 爆仓价:%.2f 平仓后预计: 仓位保证金:%g 盈亏:%f 总资产:%f\n", 1+i, positions.Market, side, pattern, openPrice, position, principal, LiquidationPrice, margin, ClosePNL, updateBalance)
		}
	}
}

// OrderBook 铺订单薄
func OrderBook(futureTrades FuturesTrading) {
	startLong := 30000         //开多(高)
	endLong := 29988           //开多(低)
	startClose := endLong + 10 //开空(低)
	endClose := startLong + 10 //开空(高)

	for longPrice := startLong; longPrice >= endLong; longPrice -= 1 {
		longPriceStr := strconv.Itoa(longPrice)
		futureTrades.LimitOpenLong("10", longPriceStr, "10000")
	}
	for closePrice := startClose; closePrice <= endClose; closePrice += 1 {
		closePriceStr := strconv.Itoa(closePrice)
		futureTrades.LimitOpenClose("10", closePriceStr, "10000")
	}
}
