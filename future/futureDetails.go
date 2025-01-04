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

type positionMmrData map[string][]TradingPair

type TradingPair struct {
	Level        int             `json:"level"`        // 风险等级
	MaxVol       int             `json:"maxVol"`       // 最大持仓量
	MMR          decimal.Decimal `json:"mmr"`          // 维持保证金率
	IMR          decimal.Decimal `json:"imr"`          // 初始保证金率
	MaxLeverage  int             `json:"maxLeverage"`  // 最大杠杆
	Symbol       string          `json:"symbol"`       // 交易对
	PositionType int             `json:"positionType"` // 开仓类型
	OpenType     int             `json:"openType"`     // 保证金模式
	Leverage     int             `json:"leverage"`     // 杠杆倍数
	LimitByAdmin bool            `json:"limitByAdmin"`
}

// PositionMmrDetails 持仓量对应维持保证率详情
func PositionMmrDetails() (error, []*TradingPair) {
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
	var result []*TradingPair

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
		return errors.New("解析get请求SymbolDetails错误信息:" + err.Error()), nil
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
