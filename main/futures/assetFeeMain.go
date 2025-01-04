package main

import (
	"fmt"
	"myapp/future"
	"time"
)

func main() {
	symbol := "PEOPLE_USDT"
	startDate := "2024-06-13"
	assetType := "FUNDING"

	// 解析开始日期
	startTime, err := time.ParseInLocation("2006-01-02", startDate, time.Local)
	if err != nil {
		fmt.Println("解析开始日期时出错:", err)
		return
	}

	// 设置开始时间为当日午夜
	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.Local)
	startTimeUnix := startTime.UnixNano() / int64(time.Millisecond)

	// 设置结束时间为当天的23:59:59
	endTime := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 23, 59, 59, 999999999, time.Local)
	endTimeUnix := endTime.UnixNano() / int64(time.Millisecond)

	err, assetFee := future.CalcFee(symbol, startTimeUnix, endTimeUnix, assetType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(assetType, "总费用为:", assetFee, startTimeUnix, endTimeUnix)
}
