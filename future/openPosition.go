package future

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/future"
	"myapp/function"
)

type Position struct {
	Data []PositionInfo `json:"data"`
}

type PositionInfo struct {
	Symbol         string          `json:"symbol"`
	PositionType   int             `json:"positionType"` // 1多 2空
	OpenType       int             `json:"openType"`     // 1逐仓 2全仓
	HoldVol        decimal.Decimal `json:"totalVol"`
	HoldAvgPrice   decimal.Decimal `json:"holdAvgPrice"`
	LiquidatePrice decimal.Decimal `json:"liqPri"`
	Oim            decimal.Decimal `json:"oim"`
	Leverage       int             `json:"leverage"`
}

// OpenPosition 持仓详情
func OpenPosition() (error, []*PositionInfo) {
	responseTest, err := function.GetDetails(config.OpenPositionUrl)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	var positionData Position
	if err := json.Unmarshal(responseTest, &positionData); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return err, nil
	}

	var positions []*PositionInfo
	if positionData.Data != nil {
		for _, position := range positionData.Data {
			newPositions := position
			positions = append(positions, &newPositions)
		}
	}
	return nil, positions
}
