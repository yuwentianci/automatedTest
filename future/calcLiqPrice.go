package future

import (
	"github.com/shopspring/decimal"
	config "myapp/config/future"
)

//// IsolatedLiquidate 逐仓强平价格
//func IsolatedLiquidate(openAvgPrice, holdVol decimal.Decimal, s SymbolData) {
//	var One = decimal.NewFromInt(1)
//	FvDc := decimal.NewFromFloat(s.Fv)
//	MmrDc := decimal.NewFromFloat(s.Mmr)
//	TfrDec := decimal.NewFromFloat(s.Tfr)
//	// 多仓强平价格 =（开仓均价*持仓数量*面值*(1+维持保证金率)-仓位保证金) / (持仓数量*面值*（1-流动性提取方费率))
//	liquidationPrice := ((openAvgPrice.Mul(holdVol).Mul(FvDc).Mul(One.Add(MmrDc))).Sub(仓位保证金)).Div(openAvgPrice.Mul(FvDc).Mul(One.Sub(TfrDec)))
//	// 空仓强平价格 =（开仓均价*持仓数量*面值*(1-维持保证金率)+仓位保证金) / (持仓数量*面值*（1+流动性提取方费率))
//	fmt.Println(liquidationPrice)
//}

// CrossLiquidate 全仓强平价格
func CrossLiquidate(position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo, asset *AssetInfo, tradingPair []*TradingPair) decimal.Decimal {
	liquidationPrice := decimal.Zero
	longPositionVales := decimal.Zero
	shortPositionVales := decimal.Zero
	longPositionSize := decimal.Zero
	shortPositionSize := decimal.Zero
	availableMargin := AvailableAsset(position, ticker, symbol, asset, tradingPair)
	for _, positions := range position {
		if positions.Symbol == config.Symbol {
			// 查找对应的交易对详情
			var Symbols *SymbolInfo
			for _, symbols := range symbol {
				if positions.Symbol == symbols.Symbol {
					Symbols = symbols
					break
				}
			}

			if positions.OpenType == 2 && positions.PositionType == 1 {
				longPositionVales = positions.HoldVol.Mul(positions.HoldAvgPrice).Mul(Symbols.Fv)
				longPositionSize = positions.HoldVol
			} else if positions.OpenType == 2 && positions.PositionType == 2 {
				shortPositionVales = positions.HoldVol.Mul(positions.HoldAvgPrice).Mul(Symbols.Fv)
				shortPositionSize = positions.HoldVol
			}
			// 强平价格 = (可用保证金 - 多仓仓位价值 + 空仓仓位价值) / (面值 * ((流动性提取方费率- 1) * 多仓持仓数量 + (流动性提取方费率 + 1) * 空仓持仓数量))
			liquidationPrice = availableMargin.Sub(longPositionVales).Add(shortPositionVales).Div(Symbols.Fv.Mul(Symbols.Tfr.Sub(One).Mul(longPositionSize).Add(Symbols.Tfr.Add(One).Mul(shortPositionSize))))
		}
	}
	return liquidationPrice
}
