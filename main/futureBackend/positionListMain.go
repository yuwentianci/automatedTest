package main

import (
	"fmt"
	"myapp/futureBackend"
)

func main() {
	//startTime := "2024-08-21 08:00:00"
	//endTime := "2024-08-22 07:59:59"
	//timeZone := "Asia/Shanghai" // 替换为你想要的时区，例如 "America/New_York"
	//
	//// 定义时间格式
	//layout := "2006-01-02 15:04:05"
	//
	//// 加载时区
	//location, err := time.LoadLocation(timeZone)
	//if err != nil {
	//	fmt.Println("加载时区错误:", err)
	//	return
	//}
	//
	//// 解析开始时间
	//start, err := time.ParseInLocation(layout, startTime, location)
	//if err != nil {
	//	fmt.Println("解析开始时间错误:", err)
	//	return
	//}
	//
	//// 解析结束时间
	//end, err := time.ParseInLocation(layout, endTime, location)
	//if err != nil {
	//	fmt.Println("解析结束时间错误:", err)
	//	return
	//}
	//
	//// 转换为毫秒级时间戳
	////uid, pageSize, startTimestamp, endTimestamp := 1190653, 10, start.UnixNano()/int64(time.Millisecond), end.UnixNano()/int64(time.Millisecond)
	////err, d := futureBackend.CalcTradingVolume(uid, pageSize, startTimestamp, endTimestamp)
	////if err != nil {
	////	fmt.Println("获取数据错误:", err)
	////}
	////fmt.Println("获取数据成功:", d)
	//
	//// 计算某位用户看涨看跌盈亏
	//uid, pageSize, roundId, startTimestamp, endTimestamp := 1190533, 100, 5909, start.UnixNano()/int64(time.Millisecond), end.UnixNano()/int64(time.Millisecond)
	//err, risePnl, fallPnl := futureBackend.SecondPNL(uid, pageSize, roundId, startTimestamp, endTimestamp)
	//if err != nil {
	//	fmt.Println("获取数据错误:", err)
	//}
	//fmt.Println("看涨盈亏:", risePnl, "看跌盈亏:", fallPnl)

	err, HoldVol := futureBackend.CalcPositionVolume()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(HoldVol)
}
