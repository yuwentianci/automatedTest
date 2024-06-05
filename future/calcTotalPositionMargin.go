package future

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	config "myapp/config/future"
	"myapp/function"
)

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
