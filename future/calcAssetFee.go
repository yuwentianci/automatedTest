package future

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	config "myapp/config/future"
	"myapp/function"
)

type AssetRecord struct {
	Success bool            `json:"success"`
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    AssetRecordData `json:"data"`
}

type AssetRecordData struct {
	CurrentPage int               `json:"currentPage"`
	PageSize    int               `json:"pageSize"`
	TotalNum    int               `json:"totalNum"`
	TotalPage   int               `json:"totalPage"`
	ResultList  []AssetRecordInfo `json:"resultList"`
}

type AssetRecordInfo struct {
	Type       string  `json:"type"`
	Direction  string  `json:"direction"`
	Currency   string  `json:"c"`
	Amount     float64 `json:"a"`
	UpdateTime int64   `json:"updateTime"`
	Symbol     string  `json:"s"`
}

func CalcFee(symbol string, startTime, endTime int64, assetType string) (error, decimal.Decimal) {
	totalFundingAmount := decimal.Zero
	currentPage := 1

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%s&symbol=%s&start_time=%d&end_time=%d&type=%s&currentPage=%d", config.AssetRecordUrl, symbol, startTime, endTime, assetType, currentPage)

		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var assetRecord AssetRecord
		if err := json.Unmarshal(responseTest, &assetRecord); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
		}

		// 检查响应是否成功
		if assetRecord.Success {
			// 遍历结果列表
			for _, item := range assetRecord.Data.ResultList {
				if item.Type == assetType {
					amount := decimal.NewFromFloat(item.Amount)
					totalFundingAmount = totalFundingAmount.Add(amount)
				}
			}

			// 检查是否还有更多的页面
			if assetRecord.Data.CurrentPage < assetRecord.Data.TotalPage {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败:" + assetRecord.Message), decimal.Zero
		}
	}

	return nil, totalFundingAmount
}

func CalcCloseFee(position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo) decimal.Decimal {
	TotalFee := decimal.Zero
	// 遍历持仓
	for _, positions := range position {
		// 查找对应的行情
		var Tickers *TickerInfo
		for _, tickers := range ticker {
			if positions.Symbol == tickers.Symbol {
				Tickers = tickers
				break
			}
		}

		// 查找对应的交易对详情
		var Symbols *SymbolInfo
		for _, symbols := range symbol {
			if positions.Symbol == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		if Tickers == nil || Symbols == nil {
			return decimal.Zero
		}

		// 计算持仓的除自己交易对外其他全仓仓位的平仓手续费
		if positions.Symbol != config.Symbol {
			TotalFee = TotalFee.Add(positions.HoldVol.Mul(Tickers.FairPrice).Mul(Symbols.Fv).Mul(Symbols.Tfr))
		}
	}

	return TotalFee
}
