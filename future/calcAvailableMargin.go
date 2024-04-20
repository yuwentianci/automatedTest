package future

import (
	"github.com/shopspring/decimal"
)

var One = decimal.NewFromInt(int64(1))
var Two = decimal.NewFromInt(int64(2))

func AvailableMargin(position []*PositionInfo, ticker []*TickerInfo, symbol []*SymbolInfo, asset *AssetInfo) decimal.Decimal {

	// 全部全仓的未实现盈亏(亏损)
	crossUnrealizedPNL := CrossUnrealizedPNL(position, ticker, symbol)

	// 用户available - 全部全仓的未实现盈亏(亏损) - 体验金
	availableMargin := asset.AvailableBalance.Sub(crossUnrealizedPNL)

	return availableMargin
}
