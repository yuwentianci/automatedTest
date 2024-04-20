package future

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/future"
	"myapp/function"
)

type DetailsLeverage struct {
	Data []LvgInfo `json:"data"`
}

type LvgInfo struct {
	MaxVol       decimal.Decimal `json:"maxVol"`
	Mmr          decimal.Decimal `json:"mmr"`
	Imr          decimal.Decimal `json:"imr"`
	Symbol       string          `json:"symbol"`
	PositionType int             `json:"positionType"`
	OpenType     int             `json:"openType"`
	CurrentMmr   decimal.Decimal `json:"currentMmr"`
	Leverage     decimal.Decimal `json:"leverage"`
}

// LvgMmrDetails 杠杆对应维持保证率详情
func LvgMmrDetails() (error, []*LvgInfo) {
	lvgUrl := config.LvgUrl + config.Symbol
	responseTest, err := function.GetDetails(lvgUrl)
	if err != nil {
		return err, nil
	}

	var lvgData DetailsLeverage
	if err := json.Unmarshal(responseTest, &lvgData); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
	}

	var result []*LvgInfo
	if len(lvgData.Data) > 0 {
		for _, lvgDetails := range lvgData.Data {
			newLvgDetails := lvgDetails
			result = append(result, &newLvgDetails)
		}
	}
	return nil, result
}

type DetailsPosition struct {
	Data positionMmrData `json:"data"`
}

type positionMmrData map[string][]tradingPairs

type tradingPairs struct {
	Imr       float64 `json:"imr"`
	Symbol    string  `json:"symbol"`
	Direction int     `json:"direction"`
	OpenType  int     `json:"openType"`
	Leverage  int     `json:"lvg"`
}

// PositionMmrDetails 持仓量对应维持保证率详情
func PositionMmrDetails() (error, []*tradingPairs) {
	responseTest, err := function.GetDetails(config.RiskLimitUrl)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	var Leverages DetailsPosition
	if err := json.Unmarshal(responseTest, &Leverages); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return err, nil
	}
	var result []*tradingPairs

	if Leverages.Data != nil {
		for _, tradingPair := range Leverages.Data {
			for _, pair := range tradingPair {
				newPair := pair
				result = append(result, &newPair)
			}
		}
	}
	return nil, result

}

type futureDetails struct {
	Data []SymbolInfo `json:"data"`
}

type SymbolInfo struct {
	Symbol string          `json:"symbol"`
	Fv     decimal.Decimal `json:"cs"`
	MaxLvg decimal.Decimal `json:"maxL"`
	Tfr    decimal.Decimal `json:"tfr"`
	Mfr    decimal.Decimal `json:"mfr"`
	Mmr    decimal.Decimal `json:"mmr"`
	Imr    decimal.Decimal `json:"imr"`
	Rbv    decimal.Decimal `json:"rbv"`
	Riv    decimal.Decimal `json:"riv"`
	Rim    decimal.Decimal `json:"rim"`
	Rii    decimal.Decimal `json:"rii"`
	Rll    decimal.Decimal `json:"rll"`
	State  decimal.Decimal `json:"state"`
}

func SymbolDetails() (error, []*SymbolInfo) {
	requestTest, err := function.GetDetails(config.DetailUrl)
	if err != nil {
		return err, nil
	}

	var SymbolData futureDetails
	if err := json.Unmarshal(requestTest, &SymbolData); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
	}

	var result []*SymbolInfo
	if len(SymbolData.Data) > 0 {
		for _, symbol := range SymbolData.Data {
			newSymbol := symbol
			result = append(result, &newSymbol)
		}
	}
	return nil, result

}
