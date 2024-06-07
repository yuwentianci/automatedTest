package earn

import (
	"encoding/json"
	"fmt"
	"myapp/config/earn"
	"myapp/function"
)

// EarnHistory 理财申购记录
type EarnHistory struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Result  earnHistoryResult `json:"result"`
}

// EarnSavingResult 理财申购记录result数据
type earnHistoryResult struct {
	Total  int               `json:"total"`
	Offset int               `json:"offset"`
	Limit  int               `json:"limit"`
	Data   []earnHistoryData `json:"data"`
}

// earnHistoryData 理财申购记录result数据的data数据
type earnHistoryData struct {
	CreateAtUnix       int64   `json:"create_at_unix"`
	CreateAtMilliUnix  int64   `json:"create_at_milli_unix"`
	Time               string  `json:"time"`
	MobileTime         string  `json:"mobile_time"`
	TermType           string  `json:"term_type"`
	Token              string  `json:"token"`
	Amount             string  `json:"amount"`
	Type               string  `json:"type"`
	Status             string  `json:"status"`
	Icon               string  `json:"icon"`
	LockPeriod         int     `json:"lock_period"`
	OriginalAmount     string  `json:"original_amount"`
	PrincipalRedeemed  string  `json:"principal_redeemed"`
	InterestPaid       string  `json:"interest_paid"`
	Rate               float64 `json:"rate"`
	RateDate           string  `json:"rate_date"`
	ActivityType       string  `json:"activity_type"`
	EarnSavingID       string  `json:"earn_saving_id"`
	RedemptionDateUnix int64   `json:"redemption_date_unix"`
}

// MyEarnHistory 查询理财购买记录
func MyEarnHistory(limit int) {
	offset := 1

	for {
		// 创建一个 GET 请求
		url := fmt.Sprintf(earn.MyEarnHistoryUrl+"offset=%d&limit=%d", offset, limit)
		responseText, err := function.GetDetails(url)
		if err != nil {
			fmt.Println(err)
		}

		var earnSaving EarnHistory
		if err := json.Unmarshal(responseText, &earnSaving); err != nil {
			fmt.Println("解析JSON响应时发生错误:", err)
			return
		}

		if earnSaving.Result.Data != nil {
			for _, entry := range earnSaving.Result.Data {
				createAtUnix := function.To24H(entry.CreateAtUnix)
				termType := function.EarnTypeMap(entry.TermType)
				fmt.Println("购买时间:", createAtUnix, "类型:", termType,
					"产品:", entry.Token, "购买数量:", entry.Amount, "锁仓期限:", entry.LockPeriod)
			}
		} else {
			fmt.Println("没有数据")
		}

		if offset*limit <= earnSaving.Result.Total {
			// 更新当前页数
			offset++
		} else {
			// 所有页面都已遍历完成
			break
		}
	}
}

// Redeem 查询理财赎回记录
func Redeem(limit int) {
	offset := 1

	for {
		// 创建一个 GET 请求
		url := fmt.Sprintf(earn.RedeemUrl+"offset=%d&limit=%d", offset, limit)
		responseText, err := function.GetDetails(url)
		if err != nil {
			fmt.Println(err)
		}

		var redeem EarnHistory
		if err := json.Unmarshal(responseText, &redeem); err != nil {
			fmt.Println("解析JSON响应时发生错误:", err)
			return
		}

		if redeem.Result.Data != nil {
			for _, entry := range redeem.Result.Data {
				createAtUnix := function.To24H(entry.CreateAtUnix)
				termType := function.EarnTypeMap(entry.TermType)
				fmt.Println("赎回时间:", createAtUnix, "类型:", termType,
					"产品:", entry.Token, "赎回数量:", entry.Amount)
			}
		} else {
			fmt.Println("没有数据")
		}

		if offset*limit <= redeem.Result.Total {
			// 更新当前页数
			offset++
		} else {
			// 所有页面都已遍历完成
			break
		}
	}
}

// Profit 查询理财利息记录
func Profit(limit int) {
	offset := 1
	for {
		// 创建一个 GET 请求
		url := fmt.Sprintf(earn.ProfitUrl+"offset=%d&limit=%d", offset, limit)
		responseText, err := function.GetDetails(url)
		if err != nil {
			fmt.Println(err)
		}

		var Profit EarnHistory
		if err := json.Unmarshal(responseText, &Profit); err != nil {
			fmt.Println("解析JSON响应时发生错误:", err)
			return
		}
		if Profit.Result.Data != nil {
			for _, entry := range Profit.Result.Data {
				createAtUnix := function.To24H(entry.CreateAtUnix)
				termType := function.EarnTypeMap(entry.TermType)
				fmt.Println("利息发放时间:", createAtUnix, "类型:", termType,
					"产品:", entry.Token, "利息:", entry.Amount)
			}
		} else {
			fmt.Println("没有数据")
		}

		if offset*limit <= Profit.Result.Total {
			// 更新当前页数
			offset++
		} else {
			// 所有页面都已遍历完成
			break
		}
	}
}
