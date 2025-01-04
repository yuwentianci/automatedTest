package main

import (
	"fmt"
	"myapp/KLine"
)

func main() {
	start := 862444800
	end := 1729814400
	symbol := "BTC_USDT"
	interval := "Month1"
	//spotInterval := 86400
	// 获取K线数据中的收盘价
	err, futureClosePrices := KLine.FutureKLine(symbol, interval, start, end)
	if err != nil {
		fmt.Println("获取K线数据出错:", err)
		return
	}

	// 设定计算EMA的周期
	period := 9.0
	futureEMA := KLine.CalculateEMA(futureClosePrices, period)
	futureRoundedEMA := KLine.OutputEMA(futureEMA, 1)
	fmt.Println(futureRoundedEMA)

	//err, spotClosePrices := KLine.SpotKLine(symbol, start, end, spotInterval)
	//if err != nil {
	//	fmt.Println("获取K线数据出错:", err)
	//	return
	//}
	//
	//spotEMA := KLine.CalculateEMA(spotClosePrices, period)
	//spotRoundedEMA := KLine.OutputEMA(spotEMA)
	//fmt.Println(spotRoundedEMA)

}
