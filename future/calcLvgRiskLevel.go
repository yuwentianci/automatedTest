package future

import (
	"github.com/shopspring/decimal"
	config "myapp/config/future"
)

func RiskLevel(leverage decimal.Decimal, Symbol []*SymbolInfo) decimal.Decimal {
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
