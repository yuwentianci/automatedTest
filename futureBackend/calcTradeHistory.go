package futureBackend

import (
	"errors"
	"github.com/shopspring/decimal"
	"myapp/config/futureBackend"
	"myapp/function"
)

type LiqOrderResponse struct {
	Success bool         `json:"success"`
	Code    int          `json:"code"`
	Msg     string       `json:"msg"`
	Data    LiqOrderData `json:"data"`
}

type LiqOrderData struct {
	CurrentPage int            `json:"currentPage"`
	ShowCount   int            `json:"showCount"`
	ResultList  []LiqOrderInfo `json:"resultList"`
	TotalResult int            `json:"totalResult"`
}

type LiqOrderInfo struct {
	ContractID  int    `json:"contractId"`
	Vol         string `json:"vol"`
	ClearAmount string `json:"clearAmount"`
	UpdateTime  int64  `json:"updateTime"`
	CreateTime  int64  `json:"createTime"`
}

func CalcLiqOrder(uid, pageSize, contractId int, startTime, endTime int64) (error, decimal.Decimal, decimal.Decimal) {
	holdVol := decimal.Zero
	mm := decimal.Zero
	currentPage := 1

	for {
		// 构建 rawData
		rawData := map[string]interface{}{
			"startTime":  startTime,
			"endTime":    endTime,
			"pageIndex":  currentPage,
			"pageSize":   pageSize,
			"uid":        uid,
			"contractId": contractId,
		}

		var liqOrderResponse LiqOrderResponse
		if err := function.PostByteDetailsComplete(futureBackend.LiqListUrl, rawData, &liqOrderResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero, decimal.Zero
		}

		// 检查响应是否成功
		if liqOrderResponse.Msg == "success" {
			// 遍历结果列表
			for _, item := range liqOrderResponse.Data.ResultList {
				vol, err := decimal.NewFromString(item.Vol)
				if err != nil {
					return err, decimal.Zero, decimal.Zero
				}

				clearAmount, err := decimal.NewFromString(item.ClearAmount)
				if err != nil {
					return err, decimal.Zero, decimal.Zero
				}

				holdVol = holdVol.Add(vol)
				mm = mm.Add(clearAmount)
			}

			// 检查是否还有更多的页面
			totalCurrentPage := 20 * currentPage
			if totalCurrentPage < liqOrderResponse.Data.TotalResult {
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

	return nil, holdVol, mm
}
