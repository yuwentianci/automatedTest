package KLine

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/Kline"
	"myapp/function"
	"myapp/future"
)

type FutureKLineResponse struct {
	Success bool                    `json:"success"`
	Code    int                     `json:"code"`
	Data    FutureKLineResponseInfo `json:"data"`
}

type FutureKLineResponseInfo struct {
	C []decimal.Decimal `json:"c"`
}

// FutureKLine 获取合约K线数据
func FutureKLine(symbol, interval string, start, end int) (error, []decimal.Decimal) {
	url := fmt.Sprintf("%s%s?end=%d&start=%d&interval=%s", Kline.FutureKLineUrl, symbol, end, start, interval)
	responseTest, err := function.GetDetails(url)
	if err != nil {
		return fmt.Errorf("请求错误: %w", err), nil
	}

	var futureKLineResponse FutureKLineResponse
	if err := json.Unmarshal(responseTest, &futureKLineResponse); err != nil {
		return fmt.Errorf("解析JSON响应时发生错误: %w", err), nil
	}

	// 检查响应是否成功
	if futureKLineResponse.Code == 0 {
		return nil, futureKLineResponse.Data.C // 返回收盘价数组
	}

	return fmt.Errorf("响应失败, Code: %d", futureKLineResponse.Code), nil
}

func CalculateEMA(closePrices []decimal.Decimal, period float64) []decimal.Decimal {
	// 计算α
	periodDec := decimal.NewFromFloat(period)
	alpha := future.Two.Div(periodDec.Add(future.One))

	// 初始化EMA数组
	ema := make([]decimal.Decimal, len(closePrices))

	// 初始EMA值为第一个价格
	ema[0] = closePrices[0]

	// 循环计算EMA
	for i := 1; i < len(closePrices); i++ {
		ema[i] = alpha.Mul(closePrices[i].Sub(ema[i-1])).Add(ema[i-1])
	}

	return ema
}

func OutputEMA(ema []decimal.Decimal, place int32) []decimal.Decimal {
	roundedEMA := make([]decimal.Decimal, len(ema))

	// 遍历EMA数组并保留两位小数
	for i, value := range ema {
		roundedEMA[i] = value.Round(place)
	}

	return roundedEMA
}

type SpotKLineResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Result  [][]interface{} `json:"result"`
}

type SpotKLineInfo struct {
	ClosePrice string `json:"close_price"` // 收盘价
}

func SpotKLine(symbol string, start, end, interval int) (error, []decimal.Decimal) {
	url := fmt.Sprintf("%ssymbol=%s&start=%d&end=%d&interval=%d", Kline.SpotKLineUrl, symbol, start, end, interval)
	responseTest, err := function.GetDetails(url)
	if err != nil {
		return fmt.Errorf("请求错误: %w", err), nil
	}

	var spotKLineResponse SpotKLineResponse
	if err := json.Unmarshal(responseTest, &spotKLineResponse); err != nil {
		return fmt.Errorf("解析JSON响应时发生错误: %w", err), nil
	}

	var closePrices []decimal.Decimal
	// 检查响应是否成功
	if spotKLineResponse.Code == 0 {
		for _, item := range spotKLineResponse.Result {
			closePrice, _ := decimal.NewFromString(item[2].(string))
			closePrices = append(closePrices, closePrice)
		}
	}

	return nil, closePrices
}
