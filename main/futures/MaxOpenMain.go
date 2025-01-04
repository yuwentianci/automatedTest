package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/future"
)

func main() {

	// 通过委托价计算最大可开
	position, ticker, symbol, lvg, openOrder, asset := futureInfo()

	positionType := 2
	price := decimal.NewFromFloat(61232.3)

	crossUnrealizedPNL := future.TotalCrossUnrealizedPNL(position, ticker, symbol)
	fmt.Println(crossUnrealizedPNL)

	availableMaxOpen := future.AvailableMaxOpen(price, positionType, position, ticker, symbol, lvg, asset)
	fmt.Println(availableMaxOpen)

	maxOpen := future.MaxOpen(price, positionType, position, ticker, symbol, lvg, openOrder, asset)
	fmt.Println(maxOpen)
}

func futureInfo() ([]*future.PositionInfo, []*future.TickerInfo, []*future.SymbolInfo, []*future.LvgInfo, []*future.OpenOrderInfo, *future.AssetInfo) {
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

	// 获取杠杆详情
	err, lvg := future.LvgMmrDetails()
	if err != nil {
		fmt.Println(err)
	}

	// 获取限价委托单
	err, openOrder := future.OpenOrder()
	if err != nil {
		fmt.Println(err)
	}

	// 用户available
	err, asset := future.UserAsset()
	if err != nil {
		fmt.Println(err)
	}

	return position, ticker, symbol, lvg, openOrder, asset
}
