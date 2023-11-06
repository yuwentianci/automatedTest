package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"myapp/text"
)

const (
	closeSide         = 1
	longSide          = 2
	closeAmount       = 9000
	closeAveragePrice = 29999
	closePrice        = 0
)

func main() {
	text.UpdateBalance(closeAveragePrice, closeAmount)
	comparativeClosePosition(longSide, closePrice, closeAmount) // side = 1 空仓; side = 2 多仓
}

// 限价平仓
func limitClosePosition(side int, price, volume float64) {
	closePrice := decimal.NewFromFloat(price).String()
	closeAmount := decimal.NewFromFloat(volume).String()
	dd := new(text.FuturesTrading)
	if side == 1 {
		dd.LimitCloseShort("5", closePrice, closeAmount)
		closeBalance()
	} else {
		dd.LimitCloseLong("5", closePrice, closeAmount)
		closeBalance()
	}
}

// 市价平仓
func marketClosePosition(side int, volume float64) {
	closeAmount := decimal.NewFromFloat(volume).String()
	dd := new(text.FuturesTrading)
	if side == 1 {
		dd.MarketCloseShort("5", closeAmount)
		closeBalance()
	} else {
		dd.MarketCloseLong("5", closeAmount)
		closeBalance()
	}
}

// closeBalance 平仓后资产详情
func closeBalance() {
	dd := text.FuturesTrading{}
	available, freeze, err := dd.FuturesBalance()
	balance := available + freeze
	if err != nil {
		fmt.Println(err)
		return
	}

	Available := decimal.NewFromFloat(available).String()
	Freeze := decimal.NewFromFloat(freeze).String()
	Balance := decimal.NewFromFloat(balance).String()
	fmt.Println("平仓后可用资产:", Available, "冻结资产", Freeze, "总资产:", Balance)
}

// comparativeClosePosition 平仓对比
func comparativeClosePosition(side int, closePrice, closeAmount float64) {
	if closePrice > 0 {
		limitClosePosition(side, closePrice, closeAmount) // side = 1 空仓; side = 2 多仓
	} else {
		marketClosePosition(side, closeAmount)
	}
}
