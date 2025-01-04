package secondFuture

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	config "myapp/config/future"
	"myapp/config/secondFuture"
	"myapp/function"
)

type SecondHistoryTradeResponse struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data SecondHistoryTradeData `json:"data"`
}

type SecondHistoryTradeData struct {
	List  []SecondHistoryTradeInfo `json:"list"`
	Total int                      `json:"total"`
}

type SecondHistoryTradeInfo struct {
	Amount     string `json:"amount"`
	ClosePrice string `json:"close_price"`
	Expiration int    `json:"expiration"`
	Fee        string `json:"fee"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	OpenPrice  string `json:"open_price"`
	Profit     string `json:"profit"`
	RoundID    int    `json:"round_id"`
	Side       int    `json:"side"`
	Start      int    `json:"start"`
}

func RealizedPNL() (error, decimal.Decimal) {
	profit := decimal.Zero
	currentPage := 1
	pageSize := 10

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%sunderlying=%s&page_index=%d&page_size=%d", secondFuture.SecondHistoryTradeUrl, config.Symbol, currentPage, pageSize)

		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var secondHistoryTradeResponse SecondHistoryTradeResponse
		if err := json.Unmarshal(responseTest, &secondHistoryTradeResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
		}

		// 检查响应是否成功
		if secondHistoryTradeResponse.Code == 200 {
			// 遍历结果列表
			for _, item := range secondHistoryTradeResponse.Data.List {
				profit = profit.Add(decimal.RequireFromString(item.Profit))
			}

			// 检查是否还有更多的页面
			if secondHistoryTradeResponse.Data.Total > currentPage*pageSize {
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

	return nil, profit

}
