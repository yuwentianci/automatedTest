package earn

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/earn"
	"myapp/function"
	"sync"
	"time"
)

// BuyEarnData 购买理财
type BuyEarnData struct {
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

// BuyEarn 购买理财产品
func BuyEarn(earnID string, amount1, amount2 decimal.Decimal, step float64) {
	var wg sync.WaitGroup

	// 启动5个goroutines并发发送购买理财请求
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			stepDec := decimal.NewFromFloat(step)
			// 购买某一理财产品从20到23进行数量循环
			for amount := amount1; amount.Cmp(amount2) <= 0; amount = amount.Add(stepDec) {

				// 将amount转换为字符串类型
				amountStr := amount.String()

				// 创建要发送的 JSON 数据
				requestData := map[string]interface{}{
					"earn_saving_id": earnID,
					"amount":         amountStr,
				}

				// 将 JSON 数据序列化为字节数组
				jsonData, err := json.Marshal(requestData)
				if err != nil {
					fmt.Println("JSON 编码失败:", err)
					return
				}

				// 创建一个POST请求，将rawBytes作为请求主体数据
				responseText, err := function.PostByteDetails(earn.BuyEarnUrl, jsonData)

				// 解析JSON响应或处理raw格式响应，具体取决于API的响应类型
				var buyEarn BuyEarnData
				if err := json.Unmarshal(responseText, &buyEarn); err != nil {
					now := time.Now()
					formattedTime := now.Format("2006-01-02 15:04:05")
					fmt.Println(formattedTime, earnID, "购买失败，失败原因:", buyEarn.Message)
					//fmt.Println("解析JSON响应时发生错误:", err)
				}

				entry := buyEarn.Result
				if entry != nil && buyEarn.Code == 0 {
					// 将Unix时间戳转换为24小时制
					ctime := function.To24H(entry.SubscriptionDate)
					valueDate := function.To24H(entry.ValueDate)
					interestDistribution := function.To24H(entry.InterestDistribution)
					interestEndDate := function.To24H(entry.InterestEndDate)
					redemptionDate := function.To24H(entry.RedemptionDate)
					earnType := function.EarnTypeMap(entry.EarnType)
					CNFormat := "%s 购买%s理财数量: %s 起息日: %s 发息日: %s 计息结束日: %s 赎回日: %s\n"

					fmt.Printf(CNFormat, ctime, earnType, entry.SubscriptionAmount, valueDate, interestDistribution, interestEndDate, redemptionDate)
				}
			}
		}()
	}

	// 等待所有goroutines完成
	wg.Wait()
}
