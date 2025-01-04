package genDan

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/futureCopyTrading"
	"myapp/function"
)

type GenDanResponse struct {
	Code    int        `json:"code"`
	Data    GenDanData `json:"data"`
	Msg     string     `json:"msg"`
	Success bool       `json:"success"`
}

type GenDanData struct {
	ResultList  []GenDanResult `json:"resultList"`
	TotalResult int            `json:"totalResult"`
}

type GenDanResult struct {
	ID                 int     `json:"id"`
	UID                int     `json:"uid"`
	Trader             int     `json:"trader"`
	FollowMode         int     `json:"follow_mode"`
	ProfitSharingRatio string  `json:"profit_sharing_ratio"`
	WithholdingFunding string  `json:"withholding_funding"`
	WithholdingStatus  int     `json:"withholding_status"`
	CloseVol1          *string `json:"close_vol1"`
	ContractID         int     `json:"contract_id"`
	Vol                int     `json:"vol"`
	Margin             string  `json:"margin"`
	Leverage           int     `json:"leverage"`
	LiquidatePrice     string  `json:"liquidate_price"`
	OpenPrice          string  `json:"open_price"`
	MarginMode         int     `json:"margin_mode"`
	Chengben           *string `json:"chengben"`
	PositionType       int     `json:"position_type"`
	ClosePrice         string  `json:"close_price"`
	CloseVol           int     `json:"close_vol"`
	RealisedPnl        string  `json:"realised_pnl"`
	Rate               *string `json:"rate"`
	PositionState      int     `json:"position_state"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
}

func GenDanPNL(pageSize, uid int) (error, decimal.Decimal) {
	totalRealisedPnl := decimal.Zero
	currentPage := 1
	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%spageIndex=%d&pageSize=%d&trader=%d", futureCopyTrading.GenDanPNLUrl, currentPage, pageSize, uid)
		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var genDanResponse GenDanResponse
		if err := json.Unmarshal(responseTest, &genDanResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
		}

		// 检查响应是否成功
		if genDanResponse.Code == 200 {
			if genDanResponse.Data.ResultList != nil {
				for _, gendan := range genDanResponse.Data.ResultList {
					if gendan.UpdatedAt < "2024-10-14 00:00:00" {
						realisedPnlDec, err := decimal.NewFromString(gendan.RealisedPnl)
						if err != nil {
							return errors.New("金额转换错误:" + err.Error()), decimal.Zero
						}
						totalRealisedPnl = totalRealisedPnl.Add(realisedPnlDec)
					}
				}
			}

			// 检查是否还有更多的页面
			totalCurrentPage := pageSize * currentPage
			if totalCurrentPage < genDanResponse.Data.TotalResult {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), decimal.Zero
		}
	}

	return nil, totalRealisedPnl
}
