package main

import (
	"fmt"
	"myapp/future"
	"time"
)

func main() {
	startDate := "2024-05-27"
	symbol := "BTC_USDT"

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

}
