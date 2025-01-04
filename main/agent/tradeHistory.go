package main

import (
	"fmt"
	"myapp/agentBackend"
	"time"
)

func main() {
	// 定义时间格式
	layout := "2006-01-02 15:04:05"

	// 定义时间字符串
	startDate := "2024-09-12 00:00:00"
	endDate := "2024-09-13 23:59:59"

	// 解析时间字符串
	startTime, err := time.Parse(layout, startDate)
	if err != nil {
		fmt.Println("Error parsing start time:", err)
		return
	}

	endTime, err := time.Parse(layout, endDate)
	if err != nil {
		fmt.Println("Error parsing end time:", err)
		return
	}

	// 转换为秒时间戳
	startTimestamp := startTime.Unix()
	endTimestamp := endTime.Unix()

	// 输出结果
	fmt.Printf("Start time: %v\n", startTimestamp)
	fmt.Printf("End time: %v\n", endTimestamp)

	uid := 0
	inviterUid := 51405736
	err, commission := agentBackend.TradeHistoryCommission(uid, inviterUid, startTimestamp, endTimestamp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("commission:", commission)
	//
	//err, TradeProfit, TradeLess := agentBackend.TradeHistory(uid, inviterUid, startTimestamp, endTimestamp)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("TradeProfit:", TradeProfit, "TradeLess:", TradeLess)
	//fmt.Println("总盈亏:", TradeProfit.Add(TradeLess))
}
