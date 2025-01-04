package future

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	config "myapp/config/future"
	"myapp/function"
)

// CalcPositionMargin 所有保证金模式的仓位保证金
func CalcPositionMargin() (error, decimal.Decimal) {
	TotalPositionMargin := decimal.Zero
	responseTest, err := function.GetDetails(config.OpenPositionUrl)
	if err != nil {
		fmt.Println(err)
	}

	var positionData PositionData
	if err := json.Unmarshal(responseTest, &positionData); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
	}

	if positionData.Success == true && len(positionData.Data) > 0 {
		for _, positionInfo := range positionData.Data {
			TotalPositionMargin = TotalPositionMargin.Add(positionInfo.IM)
		}
	} else {
		return errors.New("仓位为空或仓位接口请求失败！！！"), decimal.Zero
	}

	return nil, TotalPositionMargin
}

// CalcCrossPositionMargin 所有全仓的仓位保证金
func CalcCrossPositionMargin() (error, decimal.Decimal) {
	TotalPositionMargin := decimal.Zero
	responseTest, err := function.GetDetails(config.OpenPositionUrl)
	if err != nil {
		fmt.Println(err)
	}

	var positionData PositionData
	if err := json.Unmarshal(responseTest, &positionData); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
	}

	if positionData.Success == true && len(positionData.Data) > 0 {
		for _, positionInfo := range positionData.Data {
			if positionInfo.OpenType == 2 {
				TotalPositionMargin = TotalPositionMargin.Add(positionInfo.IM)
			}
		}
	} else {
		return errors.New("全仓仓位为空或仓位接口请求失败！！！"), decimal.Zero
	}

	return nil, TotalPositionMargin
}
