package main

import (
	"fmt"
	config "myapp/config/future"

	//config "myapp/config/future"
	"myapp/future"
)

func main() {
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

	// 获取仓位维持保证金率
	err, tradingPair := future.PositionMmrDetails()
	if err != nil {
		fmt.Println(err)
	}

	// 用户available
	err, asset := future.UserAsset()
	if err != nil {
		fmt.Println(err)
	}

	err, crossPositionMargin := future.CalcCrossPositionMargin()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("所有全仓仓位保证金为:", crossPositionMargin)

	totalMm := future.TotalCrossPositionMm(position, symbol, tradingPair)
	fmt.Println("所有全仓维持保证金为:", totalMm)

	fmt.Println("可用余额:", asset.AvailableBalance)

	crossUnrealizedPNL := future.CrossUnrealizedPNL(position, ticker, symbol)
	fmt.Println("除", config.Symbol, "外其他全仓仓位的总未实现盈亏:", crossUnrealizedPNL)

	calcCloseFee := future.CalcCloseFee(position, ticker, symbol)
	fmt.Println("除", config.Symbol, "外其他全仓仓位的总平仓手续费:", calcCloseFee)

	availableMargin := future.AvailableAsset(position, ticker, symbol, asset, tradingPair)
	fmt.Println("余额为:", availableMargin)

	liquidationPrice := future.CrossLiquidate(position, ticker, symbol, asset, tradingPair)
	fmt.Println(config.Symbol, "强平价为:", liquidationPrice)
}
