package future

import (
	"github.com/shopspring/decimal"
	config "myapp/config/future"
)

func LeverageRiskLevel(leverage decimal.Decimal, Symbol []*SymbolInfo) decimal.Decimal {
	level := decimal.Zero
	for _, symbol := range Symbol {

		if symbol.Symbol == config.Symbol {
			// level：1 + (1 / riskIncrImr * (1 / lvg - imr))，向下取整
			level = One.Add(One.Div(symbol.Rim).Mul(One.Div(leverage).Sub(symbol.Imr)))
			if level.Cmp(symbol.Rll) >= 0 {
				return symbol.Rll
			}
		}
	}

	return level.Floor()
}

//// PositionRiskLevel 计算持仓的风险等级
//func PositionRiskLevel(position []*PositionInfo, symbol []*SymbolInfo) decimal.Decimal {
//	riskLevel := decimal.Zero
//
//	// 遍历持仓
//	for _, positions := range position {
//		// 查找对应的交易对详情
//		var sym *SymbolInfo
//		for _, symbols := range symbol {
//			if positions.Symbol == symbols.Symbol {
//				sym = symbols
//				break
//			}
//		}
//
//		// 如果没有找到对应的 SymbolInfo，跳过该持仓
//		if sym == nil {
//			continue
//		}
//
//		// 计算风险等级
//		// 如果持仓量小于或等于 rbv，则风险等级为 0
//		if positions.HoldVol.Cmp(sym.Rbv) <= 0 {
//			riskLevel = decimal.Zero
//		} else {
//			// 持仓量大于 rbv 时，计算风险等级
//			diff := positions.HoldVol.Sub(sym.Rbv)
//			level := diff.Div(sym.Riv).Floor() // (rbv < 持仓量 <= (rbv + riv * level))，取整
//			riskLevel = level
//		}
//	}
//
//	return riskLevel
//}
