package main

import (
	"fmt"
	"myapp/future"
)

func main() {
	//err, totalPNL := future.TotalRealizedPNL("TRB_USDT")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("totalPNL:", totalPNL)

	// 获取持仓
	err, position := future.OpenPosition()
	if err != nil {
		fmt.Println(err)
	}

	// 获取行情
	ticker := future.Ticker()

	// 获取交易对详情
	err, symbol := future.SymbolDetails()
	if err != nil {
		fmt.Println(err)
	}

	shortPNL, longPNL := future.UnrealizedPNL(position, ticker, symbol)
	fmt.Println("shortPNL:", shortPNL, "longPNL:", longPNL)

}
