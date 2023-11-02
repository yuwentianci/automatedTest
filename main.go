package main

import (
	_ "github.com/go-sql-driver/mysql"
	"myapp/text"
	"strconv"
)

func main() {

	futureTrades := text.FuturesTrading{}
	startLong := 40000           //开多盘口(高)
	endLong := 35000             //开多峰值(低)
	startClose := endLong + 6000 //开空盘口(低)
	endClose := startLong + 6000 //开空峰值(高)

	for longPrice := endLong; longPrice <= startLong; longPrice += 1000 {
		longPriceStr := strconv.Itoa(longPrice)
		futureTrades.LimitOpenLong("10", longPriceStr, "10000")
	}
	for closePrice := startClose; closePrice <= endClose; closePrice += 1000 {
		closePriceStr := strconv.Itoa(closePrice)
		futureTrades.LimitOpenClose("10", closePriceStr, "10000")
	}

	//dd.OpenOrders()
	//dd.MarginAndLeverage()
}
