package earn

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/earn"
	"myapp/function"
	"time"
)

// MyEarnData 我的理财
type MyEarnData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		Data []struct {
			Token         string          `json:"token"`
			APYRate       float64         `json:"apy_rate"`
			TermType      string          `json:"term_type"`
			Duration      int             `json:"duration"`
			Amount        decimal.Decimal `json:"amount"`
			TotalEarnings string          `json:"total_earnings"`
			Status        string          `json:"status"`
			InstanceID    string          `json:"instance_id"`
			EarnSavingID  int             `json:"earn_saving_id"`
		} `json:"data"`
		Total  int `json:"total"`
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"result"`
}

// MyFixedEarn 查询我的理财定期产品
func MyFixedEarn(limit int, asset string) {
	offset := 1
	totalAmounts := make(map[string]decimal.Decimal)

	for {
		// 构造请求URL
		url := fmt.Sprintf(earn.MyEarnUrl+"offset=%d&limit=%d&asset=%s&startTime=&endTime=&earnType=2", offset, limit, asset)
		responseText, err := function.GetDetails(url)
		if err != nil {
			fmt.Println(err)
		}

		var myEarn MyEarnData
		if err := json.Unmarshal(responseText, &myEarn); err != nil {
			fmt.Println("解析JSON响应时发生错误:", err)
			return
		}

		if len(myEarn.Result.Data) > 0 {
			for _, entry := range myEarn.Result.Data {
				totalAmounts[entry.Token] = totalAmounts[entry.Token].Add(entry.Amount)
			}
		} else {
			fmt.Println("没有数据，可能是没有购买过理财产品或者已购买的理财产品已被全部赎回")
		}

		// 检查是否还有更多的页面
		if offset*limit <= myEarn.Result.Total {
			// 更新当前页数
			offset++
		} else {
			// 所有页面都已遍历完成
			break
		}
	}

	for token, total := range totalAmounts {
		fmt.Printf("币名: %s 总数: %s\n", token, total)
	}
}

// MyEarn 查询我的理财产品
func MyEarn(limit int, asset string) []string {
	offset := 1
	idList := make([]string, 0)

	for {
		// 创建一个 GET 请求
		url := fmt.Sprintf(earn.MyEarnUrl+"offset=%d&limit=%d&asset=%s", offset, limit, asset)
		responseText, err := function.GetDetails(url)
		if err != nil {
			fmt.Println(err)
		}

		var myEarn MyEarnData
		if err := json.Unmarshal(responseText, &myEarn); err != nil {
			fmt.Println("解析JSON响应时发生错误:", err)
			return nil
		}

		if len(myEarn.Result.Data) > 0 {
			for _, entry := range myEarn.Result.Data {
				idList = append(idList, entry.InstanceID)
				//earnType := function.EarnTypeMap(entry.TermType)
				//Amount := function.FormatFloat(entry.Amount)
				// 将 APY 格式化为百分比
				//arp := fmt.Sprintf("%.2f%%", entry.APYRate*100)
				////fmt.Printf("已购:产品id:%d %s 年利率:%s %d天 %s 产品数量:%s 已产生收益:%s 理财产品id:%s\n",
				//	entry.EarnSavingID, earnType, arp, entry.Duration, entry.Token, Amount, entry.TotalEarnings, entry.InstanceID)
			}
		} else {
			fmt.Println("没有数据，可能是没有购买过理财产品或者已购买的理财产品已被全部赎回")
		}
		if offset*limit <= myEarn.Result.Total {
			// 更新当前页数
			offset++
		} else {
			// 所有页面都已遍历完成
			break
		}
	}
	return idList
}

// MyEarnDetails 我的理财详情
type MyEarnDetails struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Result  *MyEarnDetailsResult `json:"result"`
}

// MyEarnDetailsResult 我的理财详情result数据
type MyEarnDetailsResult struct {
	EarnSavingID         int         `json:"earn_saving_id"`
	Type                 string      `json:"type"`
	TermType             string      `json:"term_type"`
	Rate                 float64     `json:"rate"`
	RateDate             string      `json:"rate_date"`
	TermAmount           int         `json:"term_amount"`
	TermUnit             string      `json:"term_unit"`
	InputMin             float64     `json:"input_min"`
	TotalInputMax        float64     `json:"total_input_max"`
	TotalCurrentInput    float64     `json:"total_current_input"`
	UserMaxAmount        float64     `json:"user_max_amount"`
	OrderMaxAmount       float64     `json:"order_max_amount"`
	Progress             float64     `json:"progress"`
	APYRateLine          apyRateLine `json:"apy_rate_line"`
	Schedule             schedule    `json:"schedule"`
	Icon                 string      `json:"icon"`
	Asset                string      `json:"asset"`
	Name                 string      `json:"name"`
	InstanceList         []instance  `json:"instance_list"`
	StartTime            string      `json:"start_time"`
	EndTime              string      `json:"end_time"`
	Top                  int         `json:"top"`
	FlagName             string      `json:"flag_name"`
	FlagIcon             string      `json:"flag_icon"`
	ActivityType         string      `json:"activity_type"`
	ProfitStartTime      string      `json:"profit_start_time"`
	ProfitEndTime        string      `json:"profit_end_time"`
	EarnSavingConfigList interface{} `json:"earn_saving_config_list"`
	StartTimeUnix        int64       `json:"start_time_unix"`
	EndTimeUnix          int64       `json:"end_time_unix"`
	StartTimeDiffUnix    int64       `json:"start_time_diff_unix"`
	EndTimeDiffUnix      int64       `json:"end_time_diff_unix"`
	CurrentTimestampUnix int64       `json:"current_timestamp_unix"`
	ProfitStartTimeUnix  int64       `json:"profit_start_time_unix"`
	ProfitEndTimeUnix    int64       `json:"profit_end_time_unix"`
	RateDateUnix         int64       `json:"rate_date_unix"`
	SubscribeAtUnix      int64       `json:"subscribe_at_unix"`
	InterestPaid         string      `json:"interest_paid"`
	LockPeriod           string      `json:"lock_period"`
	InterestEndDay       int64       `json:"interest_end_day"`
	AccrueDays           string      `json:"accrue_days"`
	CumulativeInterest   string      `json:"cumulative_interest"`
	RedemptionDate       int64       `json:"redemption_date"`
	Amount               string      `json:"amount"`
}

// apyRateLine 我的理财详情result的apy_rate_line数据
type apyRateLine struct {
	XTime []string  `json:"x_time"`
	YRate []float64 `json:"y_rate"`
}

// schedule 我的理财详情result的schedule数据
type schedule struct {
	Subscription             string `json:"subscription"`
	InterestAccrual          string `json:"interest_accrual"`
	InterestDistribution     string `json:"interest_distribution"`
	Expiration               string `json:"expiration"`
	RedemptionPeriod         string `json:"redemption_period"`
	ArrivalDate              string `json:"arrival_date"`
	SubscriptionUnix         int64  `json:"subscription_unix"`
	InterestAccrualUnix      int64  `json:"interest_accrual_unix"`
	InterestDistributionUnix int64  `json:"interest_distribution_unix"`
	ExpirationUnix           int64  `json:"expiration_unix"`
	RedemptionPeriodUnix     int64  `json:"redemption_period_unix"`
	ArrivalDateUnix          int64  `json:"arrival_date_unix"`
}

// instance 我的理财详情result的instance数据
type instance struct {
	InstanceId   string  `json:"instance_id"`
	CurrentInput float64 `json:"current_input"`
	SubscribeAt  string  `json:"subscribe_at"`
	Status       string  `json:"status"`
	Asset        string  `json:"asset"`
	TotalProfit  float64 `json:"total_profit"`
	ProfitAsset  string  `json:"profit_asset"`
}

// MyEarnAssetsDetails 查询我的理财资产详情
func MyEarnAssetsDetails(idList []string) {
	var myEarnDetails MyEarnDetails

	// 创建一个 GET 请求
	for i := 0; i < len(idList); i++ {
		url := earn.MyEarnAssetsDetailsUrl + idList[i]
		responseText, err := function.GetDetails(url)
		if err != nil {
			fmt.Println(err)
		}

		err = function.ParseJsonRe(responseText, &myEarnDetails)
		if err != nil {
			return
		}

		if myEarnDetails.Result != nil {
			entry := myEarnDetails.Result
			earnType := function.EarnTypeMap(entry.Type)
			originalTime, err := time.Parse(time.RFC3339, entry.RateDate)
			if err != nil {
				fmt.Println("解析日期时间字符串出错:", err)
				return
			}

			// 重新格式化为24小时制的字符串
			RateDateStr := originalTime.Format("2006-01-02 15:04:05")
			redemptionDate := function.To24H(entry.RedemptionDate)
			interestEndDay := function.To24H(entry.InterestEndDay)
			subscribeAtUnix := function.To24H(entry.SubscribeAtUnix)
			fmt.Println("购买时间:", subscribeAtUnix, "利息发放结束时间:", interestEndDay, "赎回时间:", redemptionDate)
			fmt.Printf("理财ID:%d 理财类型:%s 投资周期:%d ", entry.EarnSavingID, earnType, entry.TermAmount)
			fmt.Println("理财年利率:", entry.Rate, "利率生效时间:", RateDateStr, "收益:", entry.InterestPaid, "已购天数:", entry.AccrueDays, "购买数量:", entry.Amount)
		} else {
			fmt.Println("没有数据")
		}
	}
}
