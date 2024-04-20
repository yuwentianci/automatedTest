package future

import (
	"github.com/shopspring/decimal"
)

// CrossUnrealizedPNL 计算全仓所有亏损仓位的总未实现盈亏
func CrossUnrealizedPNL(position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo) decimal.Decimal {
	totalLoss := decimal.Zero

	// 遍历持仓
	for _, positions := range position {
		// 查找对应的行情
		var Tickers *TickerInfo
		for _, tickers := range ticker {
			if positions.Symbol == tickers.Symbol {
				Tickers = tickers
				break
			}
		}

		// 查找对应的交易对详情
		var Symbols *SymbolInfo
		for _, symbols := range symbol {
			if positions.Symbol == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		if Tickers == nil || Symbols == nil {
			return decimal.Zero

		}

		// 计算该持仓的未实现盈亏
		if positions.OpenType == 2 { // 检查持仓是否开仓并且亏损
			var positionUnrealizedLoss decimal.Decimal
			if positions.PositionType == 1 { // 多头持仓
				positionUnrealizedLoss = Tickers.FairPrice.Sub(positions.HoldAvgPrice).Mul(positions.HoldVol).Mul(Symbols.Fv)
			} else { // 空头持仓
				positionUnrealizedLoss = positions.HoldAvgPrice.Sub(Tickers.FairPrice).Mul(positions.HoldVol).Mul(Symbols.Fv)
			}
			if positionUnrealizedLoss.LessThan(decimal.Zero) {
				totalLoss = totalLoss.Add(positionUnrealizedLoss.Abs())
			}
		}
	}

	return totalLoss
}
