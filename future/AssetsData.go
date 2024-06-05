package future

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	config "myapp/config/future"
	"myapp/function"
)

type BalanceData struct {
	Success bool          `json:"success"`
	Code    int           `json:"code"`
	Data    []BalanceInfo `json:"data"`
}

type BalanceInfo struct {
	PositionMargin float64 `json:"positionMargin"`
	Equity         float64 `json:"equity"`
	Unrealized     float64 `json:"unrealized"`
	Bonus          float64 `json:"bonus"`
	Currency       string  `json:"cur"`
	AvailableBal   float64 `json:"avlBal"`
	CanWithdraw    float64 `json:"canWithdraw"`
	FrozenBal      float64 `json:"frzBal"`
}

func AssetsData() (error, decimal.Decimal) {
	responseTest, err := function.GetDetails(config.AsserUrl)
	if err != nil {
		fmt.Println(err)
	}

	var balanceData BalanceData
	if err := json.Unmarshal(responseTest, &balanceData); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
	}

	// 检查响应是否成功
	if balanceData.Success {
		avlBal := decimal.NewFromFloat(balanceData.Data[1].AvailableBal)
		return nil, avlBal
	} else {
		return errors.New("请求失败"), decimal.Zero
	}
}
