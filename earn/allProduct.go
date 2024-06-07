package earn

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/earn"
	"myapp/function"
	"time"
)

// ProductData 理财产品
type ProductData struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Result  *earnProductResult `json:"result"`
}

// earnProductResult 理财产品result数据
type earnProductResult struct {
	Data []struct {
		Name             string         `json:"name"`
		DepositAsset     string         `json:"deposit_asset"`
		DepositAssetName string         `json:"deposit_asset_name"`
		APYRate          float64        `json:"apy_rate"`
		Weigh            int            `json:"weigh"`
		TermDesc         string         `json:"term_desc"`
		Icon             string         `json:"icon"`
		EarnSaveList     []earnSaveList `json:"earn_save_list"`
	} `json:"data"`
}

// earnSaveList 理财产品result的earn_save_list数据
type earnSaveList struct {
	EarnSavingID      int             `json:"earn_saving_id"`
	Type              string          `json:"type"`
	TermType          string          `json:"term_type"`
	Rate              decimal.Decimal `json:"rate"`
	RateDate          time.Time       `json:"rate_date"`
	TermAmount        int             `json:"term_amount"`
	TermUnit          string          `json:"term_unit"`
	InputMin          decimal.Decimal `json:"input_min"`
	TotalInputMax     decimal.Decimal `json:"total_input_max"`
	TotalCurrentInput decimal.Decimal `json:"total_current_input"`
	UserMaxAmount     decimal.Decimal `json:"user_max_amount"`
	OrderMaxAmount    decimal.Decimal `json:"order_max_amount"`
	Progress          decimal.Decimal `json:"progress"`
	APYRateLine       struct {
		XTime []interface{} `json:"x_time"`
		YRate []interface{} `json:"y_rate"`
	} `json:"apy_rate_line"`
	Schedule struct {
		Subscription             string `json:"subscription"`
		InterestAccrual          string `json:"interest_accrual"`
		InterestDistribution     string `json:"interest_distribution"`
		Expiration               string `json:"expiration"`
		RedemptionPeriod         string `json:"redemption_period"`
		ArrivalDate              string `json:"arrival_date"`
		SubscriptionUnix         int    `json:"subscription_unix"`
		InterestAccrualUnix      int    `json:"interest_accrual_unix"`
		InterestDistributionUnix int    `json:"interest_distribution_unix"`
		ExpirationUnix           int    `json:"expiration_unix"`
		RedemptionPeriodUnix     int    `json:"redem_ption_period_unix"`
		ArrivalDateUnix          int    `json:"arrival_date_unix"`
	} `json:"schedule"`
	Icon                 string      `json:"icon"`
	Asset                string      `json:"asset"`
	Name                 string      `json:"name"`
	InstanceList         interface{} `json:"instance_list"`
	StartTime            string      `json:"start_time"`
	EndTime              string      `json:"end_time"`
	Top                  int         `json:"top"`
	FlagName             string      `json:"flag_name"`
	FlagIcon             string      `json:"flag_icon"`
	ActivityType         string      `json:"activity_type"`
	ProfitStartTime      string      `json:"profit_start_time"`
	ProfitEndTime        string      `json:"profit_end_time"`
	EarnSavingConfigList interface{} `json:"earn_saving_config_list"`
	StartTimeUnix        int         `json:"start_time_unix"`
	EndTimeUnix          int         `json:"end_time_unix"`
	StartTimeDiffUnix    int         `json:"start_time_diff_unix"`
	EndTimeDiffUnix      int         `json:"end_time_diff_unix"`
	CurrentTimestampUnix int         `json:"current_timestamp_unix"`
	ProfitStartTimeUnix  int         `json:"profit_start_time_unix"`
	ProfitEndTimeUnix    int         `json:"profit_end_time_unix"`
	RateDateUnix         int         `json:"rate_date_unix"`
	SubscribeAtUnix      int         `json:"subscribe_at_unix"`
	InterestPaid         string      `json:"interest_paid"`
	LockPeriod           string      `json:"lock_period"`
	InterestEndDay       int         `json:"interest_end_day"`
	AccrueDays           string      `json:"accrue_days"`
	CumulativeInterest   string      `json:"cumulative_interest"`
	RedemptionDate       int         `json:"redemption_date"`
	Amount               string      `json:"amount"`
}

// AllEarnProduct 查询全部理财产品
func AllEarnProduct(offset, limit int, asset string) {
	// 创建一个 GET 请求
	url := fmt.Sprintf(earn.AllEarnProductUrl+"offset=%d&limit=%d&asset=%s", offset, limit, asset)
	responseText, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
	}

	var earnProduct ProductData
	if err := json.Unmarshal(responseText, &earnProduct); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}

	if earnProduct.Result != nil {
		dataCount := len(earnProduct.Result.Data)
		earnSaveListCount := 0

		for _, resultData := range earnProduct.Result.Data {
			earnSaveListCount += len(resultData.EarnSaveList)
			termDescStr := function.EarnTypeMap(resultData.TermDesc)
			APYRate := function.FormatPercentage(resultData.APYRate)
			fmt.Printf("%s最高年利率: %s 类型含有: %s\n", resultData.Name, APYRate, termDescStr)
			for _, earnSaving := range resultData.EarnSaveList {
				earnTypeStr := function.EarnTypeMap(earnSaving.Type)
				earnDays := ""
				if earnTypeStr == "定期" {
					earnDays = fmt.Sprintf("投资周期: %d天", earnSaving.TermAmount)
				}

				fmt.Printf("%s产品ID: %d\n", earnSaving.Asset, earnSaving.EarnSavingID)
				fmt.Printf("类型: %s %s 活动类型: %s\n", earnTypeStr, earnDays, earnSaving.ActivityType)
				fmt.Printf("起投⾦额: %s\n", earnSaving.InputMin)
				fmt.Printf("已购金额: %s\n", earnSaving.TotalCurrentInput)
				fmt.Printf("单人最高⾦额: %s，单笔最高⾦额: %s\n", earnSaving.UserMaxAmount, earnSaving.OrderMaxAmount)
				fmt.Println()
			}
		}
		fmt.Printf("币种总计: %d\n", dataCount)
		fmt.Printf("产品总计: %d\n", earnSaveListCount)
	} else {
		fmt.Println("没有数据")
	}
}
