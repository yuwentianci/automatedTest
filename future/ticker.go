package future

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config"
	"myapp/function"
)

type TickerData struct {
	Data []TickerInfo `json:"data"`
}

type TickerInfo struct {
	Ask1         decimal.Decimal `json:"ask1"`
	Bid1         decimal.Decimal `json:"bid1"`
	CurrentPrice decimal.Decimal `json:"curPrice"`
	FairPrice    decimal.Decimal `json:"fairPrice"`
	FundingRate  decimal.Decimal `json:"fundingRate"`
	IndexPrice   decimal.Decimal `json:"indexPrice"`
	Symbol       string          `json:"symbol"`
}

func Ticker() []*TickerInfo {
	subscribeRequest := []byte(`{"method": "subscribe.tickers"}`)
	message := function.WsDetails(subscribeRequest)

	var tickerData TickerData
	if err := json.Unmarshal(message, &tickerData); err != nil {
		fmt.Println("解析消息时出错:", err)
		return nil
	}

	var result []*TickerInfo
	if len(tickerData.Data) > 0 {
		for _, ticker := range tickerData.Data {
			newTicker := ticker
			result = append(result, &newTicker)
		}
	}

	return result
}

func WsLogin() {

	aaa := fmt.Sprintf("{\"method\":\"login\",\"params\":{\"token\": %s}}", config.Token)
	subscribeRequest := []byte(aaa)
	function.WsDetails(subscribeRequest)
}
