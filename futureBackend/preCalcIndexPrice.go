package futureBackend

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/futureBackend"
	"myapp/function"
)

type PreCalcIndexPriceResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		MarketList []PreCalcIndexPriceInfo `json:"marketList"`
	} `json:"data"`
}

type PreCalcIndexPriceInfo struct {
	MarketId   int    `json:"marketId"`
	IndexPrice string `json:"indexPrice"`
	Reason     string `json:"reason"`
}

func PreCalcIndexPrice(symbolId int, symbol, priceDecimal string) (decimal.Decimal, error) {
	rawData := map[string]interface{}{
		"indexSymbolId": symbolId,
		"baseCoin":      symbol,
		"quoteCoin":     "USDT",
		"priceDecimal":  priceDecimal,
		"algorithm":     "AUTO_WEIGHTED_MEAN",
		"marketList": []map[string]interface{}{
			{"marketId": 6, "weight": 30},
			{"marketId": 11, "weight": 10},
			{"marketId": 16, "weight": 30},
			{"marketId": 13, "weight": 15},
			{"marketId": 14, "weight": 15},
		},
	}

	var preCalcIndexPriceResponse PreCalcIndexPriceResponse
	if err := function.PostByteDetailsComplete(futureBackend.PreCalcIndexPriceUrl, rawData, &preCalcIndexPriceResponse); err != nil {
		return decimal.Zero, fmt.Errorf("error making API call: %w", err)
	}

	if !preCalcIndexPriceResponse.Success || preCalcIndexPriceResponse.Message != "success" {
		return decimal.Zero, errors.New("request failed")
	}

	indexPrice := decimal.Zero
	for _, market := range preCalcIndexPriceResponse.Data.MarketList {
		marketWeight, _ := decimal.NewFromString(fmt.Sprintf("%f", 30/100.0))
		calcIndexPrice, err := decimal.NewFromString(market.IndexPrice)
		if err != nil {
			return decimal.Zero, fmt.Errorf("error parsing index price: %w", err)
		}
		indexPrice = indexPrice.Add(calcIndexPrice.Mul(marketWeight))
	}

	return indexPrice, nil
}
