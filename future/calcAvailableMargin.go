package future

import (
	"fmt"
	"github.com/shopspring/decimal"
)

var One = decimal.NewFromInt(int64(1))
var Two = decimal.NewFromInt(int64(2))

// AvailableMargin 用户用户available资产展示
func AvailableMargin(position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo, asset *AssetInfo) decimal.Decimal {

	// 全部全仓的未实现盈亏(亏损)
	crossUnrealizedPNL := TotalCrossUnrealizedPNL(position, ticker, symbol)

	// 用户available - 全部全仓的未实现盈亏(亏损) - 体验金
	availableMargin := asset.AvailableBalance.Sub(crossUnrealizedPNL)

	return availableMargin
}

// AvailableAsset 计算全仓强平价时的余额
func AvailableAsset(position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo, asset *AssetInfo, tradingPair []*TradingPair) decimal.Decimal {
	// 除自己交易对其他全仓仓位的未实现盈亏(亏损)
	crossUnrealizedPNL := CrossUnrealizedPNL(position, ticker, symbol)

	err, crossPositionMargin := CalcCrossPositionMargin()
	if err != nil {
		fmt.Println("所有全仓仓位保证金报错:", err)
	}

	calcCloseFee := CalcCloseFee(position, ticker, symbol)

	err, totalOrderMargin := CalcOrderMargin()
	if err != nil {
		fmt.Println("所有订单保证金报错:", err)
	}

	totalMm := TotalCrossPositionMm(position, symbol, tradingPair)

	// 余额 = 所有全仓仓位保证金 - 所有全仓仓位维持保证金 + 可用余额 - 除自己交易对(未实现盈亏 + 平仓手续费) - 订单保证金
	availableMargin := crossPositionMargin.Sub(totalMm).Add(asset.AvailableBalance).Add(crossUnrealizedPNL).Sub(calcCloseFee).Sub(totalOrderMargin)

	return availableMargin
}

// TotalCrossPositionMm 所有全仓的维持保证金
func TotalCrossPositionMm(position []*PositionInfo, symbol []*SymbolInfo, tradingPair []*TradingPair) decimal.Decimal {
	totalMm := decimal.Zero

	// 遍历持仓
	for _, positions := range position {

		// 查找对应的交易对详情
		var Symbols *SymbolInfo
		for _, symbols := range symbol {
			if positions.Symbol == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		var TradingPairs *TradingPair
		for _, tradingPairs := range tradingPair {
			if positions.Symbol == tradingPairs.Symbol && positions.PositionType == tradingPairs.PositionType {
				TradingPairs = tradingPairs
				break
			}
		}

		if TradingPairs == nil || Symbols == nil {
			return decimal.Zero
		}

		if positions.OpenType == 2 {
			totalMm = totalMm.Add(positions.HoldVol.Mul(positions.HoldAvgPrice).Mul(Symbols.Fv).Mul(TradingPairs.MMR))
		}
	}

	return totalMm
}
