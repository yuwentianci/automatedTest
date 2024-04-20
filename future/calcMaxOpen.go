package future

import (
	"github.com/shopspring/decimal"
	config "myapp/config/future"
)

// MaxOpen 最大可开多张数
func MaxOpen(price decimal.Decimal, positionType int, position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo, lvg []*LvgInfo, openOrder []*OpenOrderInfo, asset *AssetInfo) decimal.Decimal {
	// 余额最大可开
	availableMaxOpen := AvailableMaxOpen(price, positionType, position, ticker, symbol, lvg, asset)
	// 剩余最大可开
	remainMaxOpen := RemainMaxOpen(positionType, position, lvg, openOrder)

	if availableMaxOpen.Cmp(remainMaxOpen) >= 0 {
		return remainMaxOpen
	}

	return availableMaxOpen
}

// AvailableMaxOpen 余额最大可开
func AvailableMaxOpen(price decimal.Decimal, positionType int, position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo, lvg []*LvgInfo, asset *AssetInfo) decimal.Decimal {
	maxOpen := decimal.Zero

	for _, symbols := range symbol {
		for _, lvgs := range lvg {
			if symbols.Symbol == config.Symbol && symbols.Symbol == lvgs.Symbol {
				//余额最大可开
				availableMargin := AvailableMargin(position, ticker, symbol, asset)

				if lvgs.PositionType == positionType && positionType == 1 {
					//多仓: 可用保证金 / ( 委托价格 * 面值 *（imr + 2 * taker))
					maxOpen = availableMargin.Div(price.Mul(symbols.Fv).Mul(One.Div(lvgs.Leverage).Add(Two.Mul(symbols.Tfr))))
				} else if lvgs.PositionType == positionType && positionType == 2 {
					//空仓: 可用保证金 / (委托价格 * 面值 * ((imr + taker) * (1 + taker) + (taker * (1 - mmr))))
					maxOpen = availableMargin.Div(price.Mul(symbols.Fv).Mul(One.Div(lvgs.Leverage).Add(symbols.Tfr).Mul(One.Add(symbols.Tfr)).Add(symbols.Tfr.Mul(One.Sub(lvgs.Mmr)))))
				}
			}
		}
	}
	return maxOpen
}

// RemainMaxOpen 剩余最大可持仓
func RemainMaxOpen(positionType int, position []*PositionInfo, lvg []*LvgInfo, openOrder []*OpenOrderInfo) decimal.Decimal {
	leverageMaxVol := LeverageMaxVol(positionType, lvg)
	totalHeldPosition := TotalHeldPosition(positionType, position)
	openOrders := TotalOpenOrder(positionType, openOrder)

	//剩余最大可持仓: lvg接口中的 maxVol - 持仓中的totalVol - 所有限价委托
	remainMaxOpen := leverageMaxVol.Sub(totalHeldPosition).Sub(openOrders)
	return remainMaxOpen
}

// LeverageMaxVol 当前杠杆最大持仓
func LeverageMaxVol(positionType int, lvg []*LvgInfo) decimal.Decimal {
	for _, lvgS := range lvg {
		if lvgS.PositionType == positionType {
			return lvgS.MaxVol
		}
	}

	return decimal.Zero
}

// TotalHeldPosition 当前交易对总持仓
func TotalHeldPosition(positionType int, position []*PositionInfo) decimal.Decimal {
	for _, positions := range position {
		if positions.Symbol == config.Symbol && positions.PositionType == positionType {
			return positions.HoldVol
		}
	}

	return decimal.Zero
}

// TotalOpenOrder 当前交易对所有限价委托单数量
func TotalOpenOrder(positionType int, openOrder []*OpenOrderInfo) decimal.Decimal {
	total := decimal.Zero

	for _, openOrders := range openOrder {
		sidePositionTypes := sidePositionType(openOrders.Side)
		if openOrders.Symbol == config.Symbol && sidePositionTypes == positionType {
			total = total.Add(openOrders.Vol)
		}
	}

	return total
}

func sidePositionType(side int) int {
	if side == 3 {
		return 2
	}

	return side
}
