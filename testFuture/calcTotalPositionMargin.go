package testFuture

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/testFuture"
	"myapp/function"
)

type PositionData struct {
	Success bool           `json:"success"`
	Code    int            `json:"code"`
	Data    []PositionInfo `json:"data"`
}

type PositionInfo struct {
	Symbol       string  `json:"symbol"`
	OpenType     int     `json:"openType"`
	HoldAvgPrice float64 `json:"holdAvgPrice"`
	OIM          float64 `json:"oim"`
	IM           float64 `json:"im"`
	PositionType int     `json:"positionType"`
	TotalVol     int     `json:"totalVol"`
	Leverage     int     `json:"leverage"`
}

func CalcPositionMargin() (error, decimal.Decimal) {
	TotalPositionMargin := decimal.Zero
	responseTest, err := function.GetDetails(testFuture.OpenPositionUrl)
	if err != nil {
		fmt.Println(err)
	}

	var positionData PositionData
	if err := json.Unmarshal(responseTest, &positionData); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
	}

	if positionData.Success == true && len(positionData.Data) > 0 {
		for _, positionInfo := range positionData.Data {
			TotalPositionMargin = TotalPositionMargin.Add(decimal.NewFromFloat(positionInfo.IM))
		}
	} else {
		return errors.New("仓位为空或仓位接口请求失败！！！"), decimal.Zero
	}

	return nil, TotalPositionMargin
}
