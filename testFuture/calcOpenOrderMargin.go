package testFuture

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/testFuture"
	"myapp/function"
)

type OrderData struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    []OrderInfo `json:"data"`
}

type OrderInfo struct {
	Symbol      string  `json:"symbol"`
	PriceStr    string  `json:"priceStr"`
	Vol         int     `json:"vol"`
	Side        int     `json:"side"`
	OrderMargin float64 `json:"orderMargin"`
	OpenType    int     `json:"openType"`
	Leverage    int     `json:"leverage"`
	Type        int     `json:"type"`
	DealMargin  float64 `json:"dealMargin"`
}

func CalcOrderMargin() (error, decimal.Decimal) {
	totalOrderMargin := decimal.Zero
	responseTest, err := function.GetDetails(testFuture.OpenOrdersUrl)
	if err != nil {
		fmt.Println(err)
	}

	var orderData OrderData
	if err := json.Unmarshal(responseTest, &orderData); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
	}

	// 检查响应是否成功
	if orderData.Success {
		for _, item := range orderData.Data {
			orderMargin := decimal.NewFromFloat(item.OrderMargin)
			dealMargin := decimal.NewFromFloat(item.DealMargin)
			totalOrderMargin = totalOrderMargin.Add(orderMargin).Sub(dealMargin)
		}
	} else {
		return errors.New("请求失败"), decimal.Zero
	}

	return nil, totalOrderMargin
}
