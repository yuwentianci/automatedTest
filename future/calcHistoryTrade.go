package future

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	config "myapp/config/future"
	"myapp/function"
)

// 定义减去8小时的毫秒数
const (
	hoursToSubtract        = 8
	millisecondsInAnHour   = 3600000
	millisecondsToSubtract = hoursToSubtract * millisecondsInAnHour
)

type HistoryOrderResponse struct {
	Success bool             `json:"success"`
	Code    int              `json:"code"`
	Data    HistoryOrderData `json:"data"`
}

type HistoryOrderData struct {
	PageSize    int                `json:"pageSize"`
	TotalCount  int                `json:"totalCount"`
	TotalPage   int                `json:"totalPage"`
	CurrentPage int                `json:"currentPage"`
	ResultList  []HistoryOrderInfo `json:"resultList"`
}

type HistoryOrderInfo struct {
	Symbol          string          `json:"symbol"`
	Side            int             `json:"side"`
	DealAvgPriceStr string          `json:"dealAvgPriceStr"`
	DealVol         int             `json:"dealVol"`
	TakerFee        decimal.Decimal `json:"takerFee"`
	MakerFee        decimal.Decimal `json:"makerFee"`
	Profit          decimal.Decimal `json:"profit"`
	State           int             `json:"state"`
	CreateTime      int64           `json:"createTime"`
	UpdateTime      int64           `json:"updateTime"`
}

// CalcHistoryOrderProfitLoss 分别计算历史订单盈利和亏损
func CalcHistoryOrderProfitLoss() (error, decimal.Decimal, decimal.Decimal) {
	profit := decimal.Zero
	loss := decimal.Zero
	currentPage := 1

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%s&page_num=%d", config.HistoryOrderUrl, currentPage)

		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var historyOrderResponse HistoryOrderResponse
		if err := json.Unmarshal(responseTest, &historyOrderResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero, decimal.Zero
		}

		// 检查响应是否成功
		if historyOrderResponse.Success {
			// 遍历结果列表
			for _, item := range historyOrderResponse.Data.ResultList {
				if item.Profit.GreaterThan(decimal.Zero) {
					profit = profit.Add(item.Profit)
				} else {
					loss = loss.Add(item.Profit)
				}
			}

			// 检查是否还有更多的页面
			if historyOrderResponse.Data.CurrentPage < historyOrderResponse.Data.TotalPage {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), decimal.Zero, decimal.Zero
		}
	}

	return nil, profit, loss
}

type HistoryTradeResponse struct {
	Success bool             `json:"success"`
	Code    int              `json:"code"`
	Data    HistoryTradeData `json:"data"`
}

type HistoryTradeData struct {
	PageSize    int                `json:"pageSize"`
	TotalCount  int                `json:"totalCount"`
	TotalPage   int                `json:"totalPage"`
	CurrentPage int                `json:"currentPage"`
	ResultList  []HistoryTradeInfo `json:"resultList"`
}

type HistoryTradeInfo struct {
	Symbol    string          `json:"symbol"`
	Side      int             `json:"side"`
	Vol       decimal.Decimal `json:"vol"`
	Price     decimal.Decimal `json:"price"`
	Fee       decimal.Decimal `json:"fee"`
	Profit    decimal.Decimal `json:"profit"`
	Timestamp int64           `json:"timestamp"`
	Taker     bool            `json:"taker"`
}

// CalcHistoryTradeProfitLoss 分别计算成交历史盈利和亏损
func CalcHistoryTradeProfitLoss() (error, decimal.Decimal, decimal.Decimal) {
	profit := decimal.Zero
	loss := decimal.Zero
	currentPage := 1

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%s&page_num=%d", config.HistoryTradeUrl, currentPage)

		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var historyTradeResponse HistoryTradeResponse
		if err := json.Unmarshal(responseTest, &historyTradeResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero, decimal.Zero
		}

		// 检查响应是否成功
		if historyTradeResponse.Success {
			// 遍历结果列表
			for _, item := range historyTradeResponse.Data.ResultList {
				if item.Profit.GreaterThan(decimal.Zero) {
					profit = profit.Add(item.Profit)
				} else {
					loss = loss.Add(item.Profit)
				}
			}

			// 检查是否还有更多的页面
			if historyTradeResponse.Data.CurrentPage < historyTradeResponse.Data.TotalPage {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), decimal.Zero, decimal.Zero
		}
	}

	return nil, profit, loss
}

type TradeVolumeData struct {
	TradeNumber      decimal.Decimal `json:"tradeNumber"`
	OpenLongVolume   decimal.Decimal `json:"open_long_volume"`
	OpenShortVolume  decimal.Decimal `json:"open_short_volume"`
	CloseLongVolume  decimal.Decimal `json:"close_long_volume"`
	CloseShortVolume decimal.Decimal `json:"close_short_volume"`
	TotalVolume      decimal.Decimal `json:"total_volume"`
	OpenLongAmount   decimal.Decimal `json:"open_long_amount"`
	OpenShortAmount  decimal.Decimal `json:"open_short_amount"`
	CloseLongAmount  decimal.Decimal `json:"close_long_amount"`
	CloseShortAmount decimal.Decimal `json:"close_short_amount"`
	TotalAmount      decimal.Decimal `json:"total_amount"`
	TakerFee         decimal.Decimal `json:"taker_fee"`
	MakerFee         decimal.Decimal `json:"maker_fee"`
	TotalFee         decimal.Decimal `json:"total_fee"`
	Timestamp        int64           `json:"timestamp"`
}

// CalcTradeVolume 统计历史交易数据
func CalcTradeVolume(startTime, endTime int64, symbol string, symbols []*SymbolInfo) (error, *TradeVolumeData) {
	TradeNumber := decimal.Zero
	OpenLongVolume := decimal.Zero
	OpenShortVolume := decimal.Zero
	CloseLongVolume := decimal.Zero
	CloseShortVolume := decimal.Zero
	TotalVolume := decimal.Zero
	OpenLongAmount := decimal.Zero
	OpenShortAmount := decimal.Zero
	CloseLongAmount := decimal.Zero
	CloseShortAmount := decimal.Zero
	TotalAmount := decimal.Zero
	TakerFee := decimal.Zero
	MakerFee := decimal.Zero
	TotalFee := decimal.Zero
	currentPage := 1

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%s&page_num=%d&symbol=%s", config.HistoryTradeUrl, currentPage, symbol)

		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var historyTradeResponse HistoryTradeResponse
		if err := json.Unmarshal(responseTest, &historyTradeResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
		}

		// 检查响应是否成功
		if historyTradeResponse.Success {
			// 遍历结果列表
			for _, item := range historyTradeResponse.Data.ResultList {
				//adjustedTimestamp := item.Timestamp - millisecondsToSubtract
				//if adjustedTimestamp >= startTime && adjustedTimestamp <= endTime {
				for _, symbolData := range symbols {
					if symbolData.Symbol == item.Symbol {
						if item.Side == 1 {
							OpenLongVolume = OpenLongVolume.Add(item.Vol)
							OpenLongAmount = OpenLongAmount.Add(item.Price.Mul(item.Vol).Mul(symbolData.Fv))
						} else if item.Side == 2 {
							CloseShortVolume = CloseShortVolume.Add(item.Vol)
							CloseShortAmount = CloseShortAmount.Add(item.Price.Mul(item.Vol).Mul(symbolData.Fv))
						} else if item.Side == 3 {
							OpenShortVolume = OpenShortVolume.Add(item.Vol)
							OpenShortAmount = OpenShortAmount.Add(item.Price.Mul(item.Vol).Mul(symbolData.Fv))
						} else if item.Side == 4 {
							CloseLongVolume = CloseLongVolume.Add(item.Vol)
							CloseLongAmount = CloseLongAmount.Add(item.Price.Mul(item.Vol).Mul(symbolData.Fv))
						}

						if item.Taker {
							TakerFee = TakerFee.Add(item.Fee)
						} else {
							MakerFee = MakerFee.Add(item.Fee)
						}
						TradeNumber = TradeNumber.Add(One)
						TotalVolume = TotalVolume.Add(item.Vol)
						TotalAmount = TotalAmount.Add(item.Price.Mul(item.Vol).Mul(symbolData.Fv))
						TotalFee = TotalFee.Add(item.Fee)
					}
				}
				//}
			}

			// 检查是否还有更多的页面
			if historyTradeResponse.Data.CurrentPage < historyTradeResponse.Data.TotalPage {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), nil
		}
	}

	tradeVolumeData := &TradeVolumeData{
		TradeNumber:      TradeNumber,
		OpenLongVolume:   OpenLongVolume,
		OpenShortVolume:  OpenShortVolume,
		CloseLongVolume:  CloseLongVolume,
		CloseShortVolume: CloseShortVolume,
		TotalVolume:      TotalVolume,
		OpenLongAmount:   OpenLongAmount,
		OpenShortAmount:  OpenShortAmount,
		CloseLongAmount:  CloseLongAmount,
		CloseShortAmount: CloseShortAmount,
		TotalAmount:      TotalAmount,
		TakerFee:         TakerFee,
		MakerFee:         MakerFee,
		TotalFee:         TotalFee,
	}

	return nil, tradeVolumeData
}

type LiqOrderResponse struct {
	Success bool         `json:"success"`
	Code    int          `json:"code"`
	Data    LiqOrderData `json:"data"`
}

type LiqOrderData struct {
	PageSize    int            `json:"pageSize"`
	TotalCount  int            `json:"totalCount"`
	TotalPage   int            `json:"totalPage"`
	CurrentPage int            `json:"currentPage"`
	ResultList  []LiqOrderInfo `json:"resultList"`
}

type LiqOrderInfo struct {
	Symbol          string          `json:"symbol"`
	DealAvgPriceStr string          `json:"dealAvgPriceStr"`
	DealVol         decimal.Decimal `json:"dealVol"`
	TakerFee        decimal.Decimal `json:"takerFee"`
	MakerFee        decimal.Decimal `json:"makerFee"`
	Profit          decimal.Decimal `json:"profit"`
	CreateTime      int64           `json:"createTime"`
	UpdateTime      int64           `json:"updateTime"`
}

// CalcLiqVolume 统计每个币对的强平数据
func CalcLiqVolume(symbol string, startTime, endTime int64) (error, decimal.Decimal, decimal.Decimal) {
	profit := decimal.Zero
	totalVol := decimal.Zero
	currentPage := 1

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%s&page_num=%d&symbol=%s", config.LiqHistoryOrderUrl, currentPage, symbol)

		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var liqOrderResponse LiqOrderResponse
		if err := json.Unmarshal(responseTest, &liqOrderResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero, decimal.Zero
		}

		// 检查响应是否成功
		if liqOrderResponse.Success {
			// 遍历结果列表
			for _, item := range liqOrderResponse.Data.ResultList {
				//adjustedTimestamp := item.CreateTime - millisecondsToSubtract
				//if adjustedTimestamp >= startTime && adjustedTimestamp <= endTime {
				profit = profit.Add(item.Profit)
				totalVol = totalVol.Add(item.DealVol)
				//}
			}

			// 检查是否还有更多的页面
			if liqOrderResponse.Data.CurrentPage < liqOrderResponse.Data.TotalPage {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), decimal.Zero, decimal.Zero
		}
	}

	return nil, profit, totalVol
}
