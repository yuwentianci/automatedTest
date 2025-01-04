package futureBackend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/futureBackend"
	"myapp/function"
)

type CommissionResponse struct {
	Code    int            `json:"code"`
	Data    CommissionData `json:"data"`
	Message string         `json:"msg"`
	Success bool           `json:"success"`
}

type CommissionData struct {
	ResultList  []CommissionInfo `json:"resultList"`
	TotalResult int              `json:"totalResult"`
}

type CommissionInfo struct {
	RewardUID int64  `json:"reward_uid"`
	Amount    string `json:"amount"`
	Status    int64  `json:"status"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func CalcCommission(pageSize, uid, rewardUID, oid, roundID int, startTime, endTime int64) (error, decimal.Decimal) {
	commission := decimal.Zero
	currentPage := 1
	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%spageIndex=%d&pageSize=%d&uid=%d&reward_uid=%d&oid=%d&round_id=%d&startTime=%d&endTime=%d", futureBackend.CommissionUrl, currentPage, pageSize, uid, rewardUID, oid, roundID, startTime, endTime)
		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var commissionResponse CommissionResponse
		if err := json.Unmarshal(responseTest, &commissionResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
		}

		// 检查响应是否成功
		if commissionResponse.Code == 200 {
			// 遍历结果列表
			for _, item := range commissionResponse.Data.ResultList {
				if item.Status == 1 || item.Status == 0 {
					amountDec, err := decimal.NewFromString(item.Amount)
					if err != nil {
						return errors.New("金额转换错误:" + err.Error()), decimal.Zero
					}

					commission = commission.Add(amountDec)
				}
			}

			// 检查是否还有更多的页面
			totalCurrentPage := pageSize * currentPage
			if totalCurrentPage < commissionResponse.Data.TotalResult {
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
