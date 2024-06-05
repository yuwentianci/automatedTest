package agentBackend

import (
	"encoding/json"
	"errors"
	"github.com/shopspring/decimal"
	"myapp/config/agentBackend"
	"myapp/function"
)

type TradeInfo struct {
	ID          int    `json:"id"`
	UID         int    `json:"uid"`
	Nickname    string `json:"nickname"`
	UserType    string `json:"user_type"`
	TradingPair string `json:"trading_pair"`
	Type        string `json:"Type"`
	Fee         string `json:"fee"`
	Cashback    string `json:"cashback"`
	PNL         string `json:"pnl"`
}

type TradeResult struct {
	Item  []TradeInfo `json:"item"`
	Total int         `json:"total"`
}

type TradeResponse struct {
	Result TradeResult `json:"result"`
}

func CalcFee() (error, decimal.Decimal) {
	requestTest, err := function.GetDetails(agentBackend.HistoryTradeUrl)
	if err != nil {
		return err, decimal.Zero
	}

	var tradeResponse TradeResponse
	if err := json.Unmarshal(requestTest, &tradeResponse); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
	}

	totalFee := decimal.Zero

	for _, tradeInfo := range tradeResponse.Result.Item {
		fee, err := decimal.NewFromString(tradeInfo.Fee)
		if err != nil {
			return errors.New("无效的费用值:" + err.Error()), decimal.Zero
		}
		totalFee = totalFee.Add(fee)
	}

	return nil, totalFee
}
