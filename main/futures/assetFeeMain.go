package main

import (
	"fmt"
	"myapp/future"
	"time"
)

func main() {
	startDate := "2024-05-01"
	startTime, err := time.ParseInLocation("2006-01-02", startDate, time.Local)
	assetType := "FUNDING"
	if err != nil {
		fmt.Println("解析开始日期时出错:", err)
		return
	}

	// 设置开始时间为当日午夜
	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.Local)

	startTimeUnix := startTime.UnixNano() / int64(time.Millisecond)
	endTime := time.Now().UnixNano() / int64(time.Millisecond)

	err, assetFee := future.CalcFee(startTimeUnix, endTime, assetType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(assetType, "总费用为:", assetFee, startTimeUnix)
}
