package main

import (
	"fmt"
	"myapp/agentBackend"
)

func main() {
	err, TradeProfit, TradeLess := agentBackend.TradeHistory(1190476)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("TradeProfit:", TradeProfit, "TradeLess:", TradeLess)
}
