package future

import (
	"github.com/shopspring/decimal"
	config "myapp/config/future"
)

// TotalCrossUnrealizedPNL 计算所有全仓亏损仓位的总未实现盈亏
func TotalCrossUnrealizedPNL(position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo) decimal.Decimal {
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

// CrossUnrealizedPNL 计算除自己外其他全仓仓位的总未实现盈亏
func CrossUnrealizedPNL(position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo) decimal.Decimal {
	totalPNL := decimal.Zero

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

		if positions.OpenType == 2 {
			var positionUnrealizedPNL decimal.Decimal
			if positions.PositionType == 1 { // 多头持仓
				positionUnrealizedPNL = Tickers.FairPrice.Sub(positions.HoldAvgPrice).Mul(positions.HoldVol).Mul(Symbols.Fv)
			} else { // 空头持仓
				positionUnrealizedPNL = positions.HoldAvgPrice.Sub(Tickers.FairPrice).Mul(positions.HoldVol).Mul(Symbols.Fv)
			}
			if positions.Symbol != config.Symbol {
				totalPNL = totalPNL.Add(positionUnrealizedPNL)
			}
		}
	}

	return totalPNL
}

// UnrealizedPNL 计算所有保证金模式的多空仓未实现盈亏
func UnrealizedPNL(position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo) (decimal.Decimal, decimal.Decimal) {
	CrossShortPNL := decimal.Zero
	CrossLongPNL := decimal.Zero
	IsolatedLongPNL := decimal.Zero
	IsolatedShortPNL := decimal.Zero
	shortPNL := decimal.Zero
	longPNL := decimal.Zero

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
			return decimal.Zero, decimal.Zero
		}

		// 计算持仓的未实现盈亏
		if positions.OpenType == 1 {
			if positions.PositionType == 1 { // 多头持仓
				IsolatedLongPNL = IsolatedLongPNL.Add(Tickers.FairPrice.Sub(positions.HoldAvgPrice).Mul(positions.HoldVol).Mul(Symbols.Fv))
			} else { // 空头持仓
				IsolatedShortPNL = IsolatedShortPNL.Add(positions.HoldAvgPrice.Sub(Tickers.FairPrice).Mul(positions.HoldVol).Mul(Symbols.Fv))
			}
		} else if positions.OpenType == 2 {
			if positions.PositionType == 1 { // 多头持仓
				CrossLongPNL = CrossLongPNL.Add(Tickers.FairPrice.Sub(positions.HoldAvgPrice).Mul(positions.HoldVol).Mul(Symbols.Fv))
			} else { // 空头持仓
				CrossShortPNL = CrossShortPNL.Add(positions.HoldAvgPrice.Sub(Tickers.FairPrice).Mul(positions.HoldVol).Mul(Symbols.Fv))
			}
		}
	}
	shortPNL = IsolatedShortPNL.Add(CrossShortPNL)
	longPNL = IsolatedLongPNL.Add(CrossLongPNL)

	return shortPNL, longPNL
}
