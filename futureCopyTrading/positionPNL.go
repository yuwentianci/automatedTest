package futureCopyTrading

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/future"
)

// DaiDanCurrentPositionPNL 带单员当前仓位开仓手续费
func DaiDanCurrentPositionPNL(position []*DaiDanCurrentPositionInfo, symbol []*future.SymbolInfo) (error, decimal.Decimal) {
	fee := decimal.Zero
	for _, positions := range position {
		// 查找对应的交易对详情
		var Symbols *future.SymbolInfo
		for _, symbols := range symbol {
			if positions.ContractName == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
		if err != nil {
			return errors.New("金额转换错误:" + err.Error()), decimal.Zero
		}
		openFee := entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul(Symbols.Tfr)
		fee = fee.Add(openFee)
	}

	return nil, fee
}

// DaiDanHistoryPositionPNL 带单者历史仓位已实现盈亏
func DaiDanHistoryPositionPNL(position []*PositionInfo, symbol []*future.SymbolInfo) (error, decimal.Decimal) {
	PNL := decimal.Zero

	// 遍历持仓
	for _, positions := range position {

		// 查找对应的交易对详情
		var Symbols *future.SymbolInfo
		for _, symbols := range symbol {
			if positions.ContractName == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		// 计算持仓的未实现盈亏

		if positions.PositionType == 1 {
			closePriceDec, err := decimal.NewFromString(positions.ClosePrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			}
			entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			} // 多头持仓
			PNL = closePriceDec.Sub(entryPriceDec).Mul(positions.CloseVolume).Mul(Symbols.Fv).Sub(entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr)).Sub(closePriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr))
		} else { // 空头持仓
			closePriceDec, err := decimal.NewFromString(positions.ClosePrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			}
			entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			}
			PNL = entryPriceDec.Sub(closePriceDec).Mul(positions.CloseVolume).Mul(Symbols.Fv).Sub(entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr)).Sub(closePriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr))
		}
	}

	return nil, PNL
}

// DaiDanHistoryPositionPNLRoi 带单者历史仓位盈亏率
func DaiDanHistoryPositionPNLRoi(position []*PositionInfo, symbol []*future.SymbolInfo, lvg []*future.LvgInfo) (error, decimal.Decimal) {
	roi := decimal.Zero
	err, PNL := DaiDanHistoryPositionPNL(position, symbol)
	if err != nil {
		return err, decimal.Zero
	}
	for _, positions := range position {
		// 查找对应的交易对详情
		var Symbols *future.SymbolInfo
		for _, symbols := range symbol {
			if positions.ContractName == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		// 查找对应的杠杆
		var Lvgs *future.LvgInfo
		for _, lvgs := range lvg {
			if positions.ContractName == lvgs.Symbol {
				Lvgs = lvgs
				break
			}
		}

		if positions.PositionType == 1 {
			entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			}
			imr := future.One.Div(positions.Leverage)
			positionMargin := entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(imr.Add(Symbols.Tfr.Mul(future.Two)))
			roi = PNL.Div(positionMargin)
			fmt.Println("盈亏:", PNL, "仓位保证金：", positionMargin)
		} else if positions.PositionType == 2 && positions.PositionType == Lvgs.PositionType {
			entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			}
			imr := future.One.Div(positions.Leverage)
			positionMargin := entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul((imr.Add(Symbols.Tfr)).Mul(future.One.Add(Symbols.Tfr)).Add(Symbols.Tfr.Mul(future.One.Sub(Lvgs.Mmr))))
			roi = PNL.Div(positionMargin)
			fmt.Println("盈亏:", PNL, "仓位保证金：", positionMargin)
		}
	}

	return nil, roi
}

// DaiDanHistoryPositionTotalPNL 带单员历史仓位总已实现盈亏
func DaiDanHistoryPositionTotalPNL(pageSize, trader, positionID int) (error, decimal.Decimal) {
	pnl := decimal.Zero
	err, position := DaiDanHistoryPosition(pageSize, trader, positionID)
	if err != nil {
		fmt.Println("错误信息：", err)
	}

	for _, positions := range position {
		releasePNL, err := decimal.NewFromString(positions.ReleasePNL)
		if err != nil {
			return errors.New("金额转换错误:" + err.Error()), decimal.Zero
		}
		pnl = pnl.Add(releasePNL)
	}

	return nil, pnl
}

// GenDanHistoryPositionPNL 跟单者历史仓位已实现盈亏
func GenDanHistoryPositionPNL(position []*PositionInfo, symbol []*future.SymbolInfo) (error, decimal.Decimal) {
	PNL := decimal.Zero

	// 遍历持仓
	for _, positions := range position {

		// 查找对应的交易对详情
		var Symbols *future.SymbolInfo
		for _, symbols := range symbol {
			if positions.ContractName == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		if positions.PositionType == 1 {
			closePriceDec, err := decimal.NewFromString(positions.ClosePrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			}
			entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			} // 多头持仓
			PNL = closePriceDec.Sub(entryPriceDec).Mul(positions.CloseVolume).Mul(Symbols.Fv).Sub(entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr)).Sub(closePriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr))
		} else { // 空头持仓
			closePriceDec, err := decimal.NewFromString(positions.ClosePrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			}
			entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			}
			PNL = entryPriceDec.Sub(closePriceDec).Mul(positions.CloseVolume).Mul(Symbols.Fv).Sub(entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr)).Sub(closePriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr))
		}
	}

	return nil, PNL
}

// GenDanHistoryPositionPNLRoi 跟单者历史仓位盈亏率
func GenDanHistoryPositionPNLRoi(position []*PositionInfo, symbol []*future.SymbolInfo) (error, decimal.Decimal) {
	roi := decimal.Zero
	err, PNL := GenDanHistoryPositionPNL(position, symbol)
	if err != nil {
		return err, decimal.Zero
	}
	for _, positions := range position {
		// 查找对应的交易对详情
		var Symbols *future.SymbolInfo
		for _, symbols := range symbol {
			if positions.ContractName == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		if positions.PositionType == 1 {
			entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			}
			imr := future.One.Div(positions.Leverage)
			positionMargin := entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(imr.Add(Symbols.Tfr.Mul(future.Two)))
			roi = PNL.Div(positionMargin)
			fmt.Println("盈亏:", PNL, "仓位保证金：", positionMargin)
		} else if positions.PositionType == 2 {
			entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
			if err != nil {
				return errors.New("金额转换错误:" + err.Error()), decimal.Zero
			}
			imr := future.One.Div(positions.Leverage)
			mmr := decimal.NewFromFloat(0.04)
			positionMargin := entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul((imr.Add(Symbols.Tfr)).Mul(future.One.Add(Symbols.Tfr)).Add(Symbols.Tfr.Mul(future.One.Sub(mmr))))
			roi = PNL.Div(positionMargin)
			fmt.Println("盈亏:", PNL, "仓位保证金：", positionMargin)
		}
	}

	return nil, roi
}

func GenDanHistoryPositionTotalPNL(position []*PositionInfo) (error,decimal.Decimal){
	totalPNL := decimal.Zero
	for _, positions := range position {
		releasePNLDec, err := decimal.NewFromString(positions.ReleasePNL)
		if err != nil {
			return errors.New("金额转换错误:" + err.Error()), decimal.Zero
		}
		totalPNL = totalPNL.Add(releasePNLDec)
	}

	return nil, totalPNL
}
