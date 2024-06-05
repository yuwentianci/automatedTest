package future

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	config "myapp/config/future"
	"myapp/function"
)

type HistoryPositionResponse struct {
	Success bool                `json:"success"`
	Code    int                 `json:"code"`
	Data    HistoryPositionData `json:"data"`
}

type HistoryPositionData struct {
	PageSize    int                   `json:"pageSize"`
	TotalCount  int                   `json:"totalCount"`
	TotalPage   int                   `json:"totalPage"`
	CurrentPage int                   `json:"currentPage"`
	ResultList  []HistoryPositionInfo `json:"resultList"`
}

type HistoryPositionInfo struct {
	Symbol   string          `json:"symbol"`
	Realised decimal.Decimal `json:"realised"`
}

func TotalRealizedPNL(symbol string) (error, decimal.Decimal) {
	totalRealizedPNL := decimal.Zero
	currentPage := 1

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%s&page_num=%d&symbol=%s", config.HistoryPositionUrl, currentPage, symbol)

		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var historyPositionResponse HistoryPositionResponse
		if err := json.Unmarshal(responseTest, &historyPositionResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
		}

		// 检查响应是否成功
		if historyPositionResponse.Success {
			// 遍历结果列表
			for _, item := range historyPositionResponse.Data.ResultList {
				totalRealizedPNL = totalRealizedPNL.Add(item.Realised)
			}

			// 检查是否还有更多的页面
			if historyPositionResponse.Data.CurrentPage < historyPositionResponse.Data.TotalPage {
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

	return nil, totalRealizedPNL
}
