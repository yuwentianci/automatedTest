package main

import (
	"fmt"
	"myapp/futureBackend"
	"time"
)

func main() {
	uid := 1190490
	pageSize := 20
	contractId := 10
	startDate := "2024-05-25"

	// 解析开始日期
	startTime, err := time.ParseInLocation("2006-01-02", startDate, time.Local)
	if err != nil {
		fmt.Println("解析开始日期时出错:", err)
		return
	}

	// 设置开始时间为当日午夜
	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.Local)
	startTimeUnix := startTime.UnixNano() / int64(time.Millisecond)

	// 设置结束时间为开始时间的23:59:59
	endTime := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 23, 59, 59, 999999999, time.Local)
	endTimeUnix := endTime.UnixNano() / int64(time.Millisecond)
	err, totalVol, mm := futureBackend.CalcLiqOrder(uid, pageSize, contractId, startTimeUnix, endTimeUnix)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("合约强平仓位数量:", totalVol, "合约强平平台盈亏(U):", mm)
}
