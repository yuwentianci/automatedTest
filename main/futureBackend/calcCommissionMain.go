package main

import (
	"fmt"
	"myapp/futureBackend"
	"time"
)

func main() {

	//一周
	//startTime := "2024-08-03 08:00:00"
	//endTime := "2024-08-10 07:59:59"

	//当天
	startTime := "2024-08-14 00:00:00"
	endTime := "2024-08-14 23:59:59"
	timeZone := "Asia/Shanghai" // 替换为你想要的时区，例如 "America/New_York"

	// 定义时间格式
	layout := "2006-01-02 15:04:05"

	// 加载时区
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		fmt.Println("加载时区错误:", err)
		return
	}

	// 解析开始时间
	start, err := time.ParseInLocation(layout, startTime, location)
	if err != nil {
		fmt.Println("解析开始时间错误:", err)
		return
	}

	// 解析结束时间
	end, err := time.ParseInLocation(layout, endTime, location)
	if err != nil {
		fmt.Println("解析结束时间错误:", err)
		return
	}

	// 转换为毫秒级时间戳
	pageSize, uid, rewardUID, oid, roundID, startTimestamp, endTimestamp := 100, 0, 1190740, 0, 3826, start.UnixNano()/int64(time.Millisecond), end.UnixNano()/int64(time.Millisecond)
	err, d := futureBackend.CalcCommission(pageSize, uid, rewardUID, oid, roundID, startTimestamp, endTimestamp)
	if err != nil {
		fmt.Println("获取数据错误:", err)
	}
	fmt.Println("获取数据成功:", d)
}
