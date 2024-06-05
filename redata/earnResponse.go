package redata

import (
	"github.com/shopspring/decimal"
	"time"
)

// BuyEarn 购买理财
type BuyEarn struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  *buyEarnResult `json:"result"`
}

// buyEarnResult 购买理财result数据
type buyEarnResult struct {
	SubscriptionAmount   string `json:"subscription_amount"`
	EarnType             string `json:"earn_type"`
	SubscriptionDate     int64  `json:"subscription_date"`
	ValueDate            int64  `json:"value_date"`
	InterestDistribution int64  `json:"interest_distribution"`
	InterestEndDate      int64  `json:"interest_end_date"`
	RedemptionDate       int64  `json:"redemption_date"`
}

// MyEarn 我的理财
type MyEarn struct {
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
	} `json:"result"`
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

// MyEarnAssets 我的理财资产
type MyEarnAssets struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Result  MyEarnAssetsResult `json:"result"`
}

// MyEarnAssetsResult 我的理财资产result数据
type MyEarnAssetsResult struct {
	TotalDeposited    float64 `json:"total_deposited"`
	TotalEarnings     float64 `json:"total_earnings"`
	YesterdayEarnings float64 `json:"yesterday_earninds"`
}

// EarnProduct 理财产品
type EarnProduct struct {
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

// EarnHistory 理财申购记录
type EarnHistory struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Result  earnHistoryResult `json:"result"`
}

// EarnSavingResult 理财申购记录result数据
type earnHistoryResult struct {
	Data []earnHistoryData `json:"data"`
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
