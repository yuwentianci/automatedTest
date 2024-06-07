package earn

import (
	"encoding/json"
	"fmt"
	"myapp/config/earn"
	"myapp/function"
)

// MyEarnAssetsData 我的理财资产
type MyEarnAssetsData struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Result  myEarnAssetsResult `json:"result"`
}

// myEarnAssetsResult 我的理财资产result数据
type myEarnAssetsResult struct {
	TotalDeposited    float64 `json:"total_deposited"`
	TotalEarnings     float64 `json:"total_earnings"`
	YesterdayEarnings float64 `json:"yesterday_earninds"`
}

// MyEarnAssets 查询我的理财资产
func MyEarnAssets() {
	// 创建一个 GET 请求
	responseText, err := function.GetDetails(earn.MyEarnAssetsUrl)
	if err != nil {
		fmt.Println(err)
	}

	var myEarnAssets MyEarnAssetsData
	if err := json.Unmarshal(responseText, &myEarnAssets); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}

	if myEarnAssets.Result != (myEarnAssetsResult{}) {
		entry := myEarnAssets.Result
		fmt.Printf("购买理财总计:%f ", entry.TotalDeposited)
		fmt.Println("累计收益:", entry.TotalEarnings, "昨日收益:", entry.YesterdayEarnings)
	} else {
		fmt.Println("没有数据")
	}
}
