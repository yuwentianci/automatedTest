package text

import (
	"encoding/json"
	"fmt"
	"myapp/function"
	"myapp/redata"
	"strconv"
	"sync"
	"time"
)

type PostEarnCreateData struct {
}

// EarnProduct 查询理财产品
func (t *PostEarnCreateData) EarnProduct(offset, limit int, asset string) {

	// 创建一个 GET 请求
	url := fmt.Sprintf("https://www.biconomy.com/api/activity/earn_saving/allearnterms?offset=%d&limit=%d&asset=%s", offset, limit, asset)
	responseText, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
	}

	var EarnProduct = redata.EarnProduct{}
	if err := json.Unmarshal(responseText, &EarnProduct); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}

	if EarnProduct.Result != nil {
		dataCount := len(EarnProduct.Result.Data)
		earnSaveListCount := 0

		for _, resultData := range EarnProduct.Result.Data {
			earnSaveListCount += len(resultData.EarnSaveList)
			termDescStr := function.EarnTypeMap(resultData.TermDesc)
			APYRate := function.FormatPercentage(resultData.APYRate)
			fmt.Printf("%s最高年利率: %s 类型含有: %s\n", resultData.Name, APYRate, termDescStr)
			for _, earnSaving := range resultData.EarnSaveList {
				earnTypeStr := function.EarnTypeMap(earnSaving.Type)
				earnDays := ""
				if earnTypeStr == "定期" {
					earnDays = fmt.Sprintf("投资周期:%d天\n", earnSaving.TermAmount)
				}

				quotient, _ := function.Divide(earnSaving.TotalCurrentInput, earnSaving.TotalInputMax)
				Progress := function.FormatPercentage(quotient)
				Rate := function.FormatPercentage(earnSaving.Rate)

				fmt.Printf("%s产品ID: %d\n", earnSaving.Asset, earnSaving.EarnSavingID)
				fmt.Printf("类型: %s\n", earnTypeStr)
				fmt.Printf("年利率: %s\n", Rate)
				fmt.Printf("%s", earnDays)
				fmt.Printf("起投⾦额: %g\n", earnSaving.InputMin)
				fmt.Printf("封顶⾦额: %g\n", earnSaving.TotalInputMax)
				fmt.Printf("已购金额: %g\n", earnSaving.TotalCurrentInput)
				fmt.Printf("单人最高⾦额: %d\n", earnSaving.UserMaxAmount)
				fmt.Printf("单笔最高⾦额: %d\n", earnSaving.OrderMaxAmount)
				fmt.Printf("进度: %s\n", Progress)
				fmt.Printf("活动类型: %s\n", earnSaving.ActivityType)
				fmt.Println()
			}
		}
		fmt.Printf("币种总计: %d\n", dataCount)
		fmt.Printf("产品总计: %d\n", earnSaveListCount)
	} else {
		fmt.Println("没有数据")
	}
}

// BuyEarn 购买理财产品
func (t *PostEarnCreateData) BuyEarn(amount1, amount2 int) {
	var wg sync.WaitGroup

	// 启动5个goroutines并发发送购买理财请求
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 购买某一理财产品从20到23进行数量循环
			for amount := amount1; amount <= amount2; amount++ {

				// 将amount转换为字符串类型
				amountStr := strconv.Itoa(amount)

				// 创建要发送的 JSON 数据
				requestData := map[string]interface{}{
					"earn_saving_id": "95",
					"amount":         amountStr,
				}

				// 将 JSON 数据序列化为字节数组
				jsonData, err := json.Marshal(requestData)
				if err != nil {
					fmt.Println("JSON 编码失败:", err)
					return
				}

				// 创建一个POST请求，将rawBytes作为请求主体数据
				url := "https://www.biconomy.com/api/activity/earn_saving/order/create"
				responseText, err := function.PostByteDetails(url, jsonData)

				// 解析JSON响应或处理raw格式响应，具体取决于API的响应类型
				var BuyEarn = redata.BuyEarn{}
				if err := json.Unmarshal(responseText, &BuyEarn); err != nil {
					if BuyEarn.Result != nil {
					}
					fmt.Println("购买理财失败，失败原因:", BuyEarn.Message)
					fmt.Println("解析JSON响应时发生错误:", err)
					return
				}

				entry := BuyEarn.Result

				// 将Unix时间戳转换为24小时制
				ctime := function.To24H(entry.SubscriptionDate)
				valueDate := function.To24H(entry.ValueDate)
				interestDistribution := function.To24H(entry.InterestDistribution)
				interestEndDate := function.To24H(entry.InterestEndDate)
				redemptionDate := function.To24H(entry.RedemptionDate)

				earnType := function.EarnTypeMap(entry.EarnType)
				CNFormat := "%s 购买%s理财数量: %s 起息日: %s 发息日: %s 计息结束日: %s 赎回日: %s\n"
				if entry != nil {
					fmt.Printf(CNFormat, ctime, earnType, entry.SubscriptionAmount, valueDate, interestDistribution, interestEndDate, redemptionDate)
				} else {
					fmt.Printf("amount值为：%d，响应内容codedata: %d，messagedata: %s\n", amount, BuyEarn.Code, BuyEarn.Message)
				}
			}
		}()
	}

	// 等待所有goroutines完成
	wg.Wait()
}

// MyEarn 查询我的理财产品
func (t *PostEarnCreateData) MyEarn(offset, limit int, asset string) []string {

	// 创建一个 GET 请求
	url := fmt.Sprintf("https://www.biconomy.com/api/activity/earn_saving/myearn?offset=%d&limit=%d&asset=%s", offset, limit, asset)
	responseText, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
	}

	var MyEarn = redata.MyEarn{}
	if err := json.Unmarshal(responseText, &MyEarn); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return nil
	}

	idList := make([]string, 0)
	if len(MyEarn.Result.Data) > 0 {
		fmt.Println(len(MyEarn.Result.Data))
		for _, entry := range MyEarn.Result.Data {
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
	return idList
}

// MyEarnAssetsDetails 查询我的理财资产详情
func (t *PostEarnCreateData) MyEarnAssetsDetails(idList []string) {
	var MyEarnDetails = redata.MyEarnDetails{}
	//client := http.Client{} // 创建一个HTTP客户端

	// 创建一个 GET 请求
	for i := 0; i < len(idList); i++ {
		url := "https://www.biconomy.com/api/activity/earn_saving/instancedetail?instance_id=" + idList[i]
		responseText, err := function.GetDetails(url)
		if err != nil {
			fmt.Println(err)
		}

		err = function.ParseJsonRe(responseText, &MyEarnDetails)
		if err != nil {
			return
		}

		if MyEarnDetails.Result != nil {
			entry := MyEarnDetails.Result
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

// MyEarnAssets 查询我的理财资产
func (t *PostEarnCreateData) MyEarnAssets() {

	// 创建一个 GET 请求
	url := "https://www.biconomy.com/api/activity/earn_saving/overview"
	responseText, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
	}

	var MyEarnAssets = redata.MyEarnAssets{}
	if err := json.Unmarshal(responseText, &MyEarnAssets); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}

	if MyEarnAssets.Result != (redata.MyEarnAssetsResult{}) {
		entry := MyEarnAssets.Result
		fmt.Printf("购买理财总计:%f ", entry.TotalDeposited)
		fmt.Println("累计收益:", entry.TotalEarnings, "昨日收益:", entry.YesterdayEarnings)
	} else {
		fmt.Println("没有数据")
	}
}

// EarnSaving 查询理财购买记录
func (t *PostEarnCreateData) EarnSaving() {

	// 创建一个 GET 请求
	url := "https://www.biconomy.com/api/activity/earn_saving/earnSaving/list?offset=1&limit=10"
	responseText, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
	}

	var EarnSaving = redata.EarnHistory{}
	if err := json.Unmarshal(responseText, &EarnSaving); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}
	if EarnSaving.Result.Data != nil {
		for _, entry := range EarnSaving.Result.Data {
			createAtUnix := function.To24H(entry.CreateAtUnix)
			termType := function.EarnTypeMap(entry.TermType)
			fmt.Println("购买时间:", createAtUnix, "类型:", termType,
				"产品:", entry.Token, "购买数量:", entry.Amount, "锁仓期限:", entry.LockPeriod)
		}
	} else {
		fmt.Println("没有数据")
	}
}

// Redeem 查询理财赎回记录
func (t *PostEarnCreateData) Redeem() {

	// 创建一个 GET 请求
	url := "https://www.biconomy.com/api/activity/earn_saving/redeem/list?offset=1&limit=10"
	responseText, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
	}

	var Redeem = redata.EarnHistory{}
	if err := json.Unmarshal(responseText, &Redeem); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}
	if Redeem.Result.Data != nil {
		for _, entry := range Redeem.Result.Data {
			createAtUnix := function.To24H(entry.CreateAtUnix)
			termType := function.EarnTypeMap(entry.TermType)
			fmt.Println("赎回时间:", createAtUnix, "类型:", termType,
				"产品:", entry.Token, "赎回数量:", entry.Amount)
		}
	} else {
		fmt.Println("没有数据")
	}
}

// Profit 查询理财利息记录
func (t *PostEarnCreateData) Profit() {

	// 创建一个 GET 请求
	url := "https://www.biconomy.com/api/activity/earn_saving/profit/list?offset=1&limit=10"
	responseText, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
	}

	var Profit = redata.EarnHistory{}
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
}
