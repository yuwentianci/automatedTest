package agentBackend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/agentBackend"
	"myapp/function"
)

type TradeHistoryResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Result  TradeHistoryData `json:"result"`
}

type TradeHistoryData struct {
	Item  []TradeHistoryInfo `json:"item"`
	Total int                `json:"total"`
}

type TradeHistoryInfo struct {
	UID        int    `json:"uid"`
	InviterUID int    `json:"inviter_uid"`
	PNL        string `json:"pnl"`
	Commission string `json:"commission"`
}

func TradeHistory(uId, inviterUid int, startTime, endTime int64) (error, decimal.Decimal, decimal.Decimal) {
	profit := decimal.Zero
	loss := decimal.Zero
	currentPage := 1

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%s&uId=%d&inviter_uid=%d&pageNumber=%d&startTime=%d&endTime=%d", agentBackend.HistoryTradeUrl, uId, inviterUid, currentPage, startTime, endTime)

		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var tradeHistoryResponse TradeHistoryResponse
		if err := json.Unmarshal(responseTest, &tradeHistoryResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero, decimal.Zero
		}

		// 检查响应是否成功
		if tradeHistoryResponse.Message == "success" {
			// 遍历结果列表
			for _, item := range tradeHistoryResponse.Result.Item {
				pnl, err := decimal.NewFromString(item.PNL)
				if err != nil {
					return err, decimal.Zero, decimal.Zero
				}

				if pnl.GreaterThan(decimal.Zero) {
					profit = profit.Add(pnl)
				} else {
					loss = loss.Add(pnl)
				}
			}

			// 检查是否还有更多的页面
			totalCurrentPage := 50 * currentPage
			if totalCurrentPage < tradeHistoryResponse.Result.Total {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), decimal.Zero, decimal.Zero
		}
	}

	return nil, profit, loss
}

func TradeHistoryCommission(uId, inviterUid int, startTime, endTime int64) (error, decimal.Decimal) {
	commission := decimal.Zero
	currentPage := 1

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%s&uId=%d&inviter_uid=%d&pageNumber=%d&startTime=%d&endTime=%d", agentBackend.HistoryTradeUrl, uId, inviterUid, currentPage, startTime, endTime)

		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var tradeHistoryResponse TradeHistoryResponse
		if err = json.Unmarshal(responseTest, &tradeHistoryResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
		}

		// 检查响应是否成功
		if tradeHistoryResponse.Message == "success" {
			// 遍历结果列表
			for _, item := range tradeHistoryResponse.Result.Item {
				commissionDec, err := decimal.NewFromString(item.Commission)
				if err != nil {
					return err, decimal.Zero
				}
				commission = commission.Add(commissionDec)
			}

			// 检查是否还有更多的页面
			totalCurrentPage := 50 * currentPage
			if totalCurrentPage < tradeHistoryResponse.Result.Total {
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

	return nil, commission
}
