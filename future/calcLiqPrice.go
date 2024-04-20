package future

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
//
//// positionMargin 仓位保证金
//func positionMargin(openAvgPrice, holdVol decimal.Decimal, s SymbolData) {
//	FvDc := decimal.NewFromFloat(s.Fv)
//	TfrDec := decimal.NewFromFloat(s.Tfr)
//	// 开仓均价 * 持仓数量 * 面值 * (起始保证金率+流动性提取方费率)
//	PositionMargin := openAvgPrice.Mul(holdVol).Mul(FvDc).Mul()
//}
