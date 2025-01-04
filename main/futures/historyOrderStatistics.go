package main

import (
	"fmt"
	"myapp/future"
	"time"
)

func main() {
	startDate := "2024-08-10"
	symbol := "ETH_USDT"

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

	// 获取交易对详情
	err, symbols := future.SymbolDetails()
	if err != nil {
		fmt.Println(err)
	}

	// 统计历史交易数据
	err, tradeVolumeData := future.CalcTradeVolume(startTimeUnix, endTimeUnix, symbol, symbols)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("成交次数:", tradeVolumeData.TradeNumber)
	fmt.Println("开多成交量(张):", tradeVolumeData.OpenLongVolume)
	fmt.Println("开空成交量(张):", tradeVolumeData.OpenShortVolume)
	fmt.Println("平多成交量(张):", tradeVolumeData.CloseLongVolume)
	fmt.Println("平空成交量(张):", tradeVolumeData.CloseShortVolume)
	fmt.Println("总交易量(张):", tradeVolumeData.TotalVolume)
	fmt.Println("开多成交额(U):", tradeVolumeData.OpenLongAmount)
	fmt.Println("开空成交额(U):", tradeVolumeData.OpenShortAmount)
	fmt.Println("平多成交额(U):", tradeVolumeData.CloseLongAmount)
	fmt.Println("平空成交额(U):", tradeVolumeData.CloseShortAmount)
	fmt.Println("总交易额(U):", tradeVolumeData.TotalAmount)
	fmt.Println("TAKER手续费(U):", tradeVolumeData.TakerFee)
	fmt.Println("MAKER手续费(U):", tradeVolumeData.MakerFee)
	fmt.Println("总交易手续费(U):", tradeVolumeData.TotalFee)
	fmt.Println(startTimeUnix, endTimeUnix)

	//err, TradeProfit, TradeLess := future.CalcHistoryTradeProfitLoss()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("TradeProfit:", TradeProfit, "TradeLess:", TradeLess)

	//err, liqProfit, liqVol := future.CalcLiqVolume(symbol, startTimeUnix, endTimeUnix)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("合约强平仓位数量:", liqVol, "合约强平平台盈亏(U):", liqProfit)

	//symbol := "BTC_USDT"
	//
	//// 1分钟
	////startTime := "2024-07-23 16:30:00"
	////endTime := "2024-07-23 16:30:59"
	//
	//// 5分钟
	////startTime := "2024-07-22 15:50:00"
	////endTime := "2024-07-22 15:54:59"
	//
	//// 15分钟
	////startTime := "2024-07-22 15:15:00"
	////endTime := "2024-07-22 15:29:59"
	//
	//// 30分钟
	////startTime := "2024-07-20 15:30:00"
	////endTime := "2024-07-20 15:59:59"
	//
	//// 1小时
	////startTime := "2024-07-20 15:00:00"
	////endTime := "2024-07-20 15:59:59"
	//
	//// 4小时
	//startTime := "2024-08-02 08:00:00"
	//endTime := "2024-08-02 11:59:59"
	//
	//// 1天
	////startTime := "2024-07-23 00:00:00"
	////endTime := "2024-07-23 23:59:59"
	//
	//// 1周
	////startTime := "2024-07-08 00:00:00"
	////endTime := "2024-07-14 23:59:59"
	//
	//// 1月
	////startTime := "2024-07-01 00:00:00"
	////endTime := "2024-07-31 23:59:59"
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
	//startTimestamp := start.UnixNano() / int64(time.Millisecond)
	//endTimestamp := end.UnixNano() / int64(time.Millisecond)
	//
	//err, longAvgPrice, shortAvgPrice, longAmount, shortAmount := future.CalcHistoryOrderAvgPrice(symbol, startTimestamp, endTimestamp)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("开始时间戳:", startTimestamp)
	//fmt.Println("结束时间戳:", endTimestamp)
	//fmt.Println("多仓平均成交价:", longAvgPrice, "多仓成交数:", longAmount)
	//fmt.Println("空仓平均成交价:", shortAvgPrice, "空仓成交数:", shortAmount)
}
