package future

import (
	"encoding/json"
	"errors"
	"github.com/shopspring/decimal"
	config "myapp/config/future"
	"myapp/function"
)

type AssetData struct {
	Data []AssetInfo `json:"data"`
}

type AssetInfo struct {
	PositionMargin   decimal.Decimal `json:"positionMargin"`
	Equity           decimal.Decimal `json:"equity"`
	Unrealized       decimal.Decimal `json:"unrealized"`
	Bonus            decimal.Decimal `json:"bonus"`
	Cur              string          `json:"cur"`
	AvailableBalance decimal.Decimal `json:"avlBal"`
	CanWithdraw      decimal.Decimal `json:"canWithdraw"`
	FrozenBalance    decimal.Decimal `json:"frzBal"`
}

func UserAsset() (error, *AssetInfo) {
	responseTest, err := function.GetDetails(config.AsserUrl)
	if err != nil {
		return err, nil
	}

	var assetData AssetData
	if err := json.Unmarshal(responseTest, &assetData); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
	}

	if len(assetData.Data) > 0 {
		for _, asset := range assetData.Data {
			if asset.Cur == "USDT" {
				return nil, &asset
			}
		}
	}

	return err, nil
}
