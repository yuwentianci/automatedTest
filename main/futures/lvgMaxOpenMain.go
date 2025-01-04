package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/future"
)

func main() {

	err, Symbol := future.SymbolDetails()
	if err != nil {
		fmt.Println(err)
	}

	leverage := decimal.NewFromInt(int64(50))
	level := future.LeverageRiskLevel(leverage, Symbol)
	fmt.Println(level)
}
