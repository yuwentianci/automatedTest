package future

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	config "myapp/config/future"
	"myapp/function"
)

type OpenOrderData struct {
	Data []OpenOrderInfo `json:"data"`
}

type OpenOrderInfo struct {
	OrderID         string          `json:"orderId"`
	Symbol          string          `json:"symbol"`
	PositionID      int             `json:"positionId"`
	PriceStr        string          `json:"priceStr"`
	Vol             decimal.Decimal `json:"vol"`
	Side            int             `json:"side"`
	DealAvgPriceStr string          `json:"dealAvgPriceStr"`
	DealVolume      decimal.Decimal `json:"dealVol"`
	OrderMargin     decimal.Decimal `json:"orderMargin"`
	TakerFee        decimal.Decimal `json:"takerFee"`
	MakerFee        decimal.Decimal `json:"makerFee"`
	Profit          decimal.Decimal `json:"profit"`
	OpenType        int             `json:"openType"`
	Leverage        decimal.Decimal `json:"leverage"`
	DealMargin      decimal.Decimal `json:"dealMargin"`
}

func OpenOrder() (error, []*OpenOrderInfo) {
	responseTest, err := function.GetDetails(config.OpenOrderUrl)
	if err != nil {
		fmt.Println(err)
	}

	var openOrderData OpenOrderData
	if err := json.Unmarshal(responseTest, &openOrderData); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
	}

	var result []*OpenOrderInfo
	if len(openOrderData.Data) > 0 {
		for _, openOrder := range openOrderData.Data {
			newOpenOrder := openOrder
			result = append(result, &newOpenOrder)
		}
	}
	return nil, result

}
