package main

import (
	"fmt"
	"myapp/text"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	aa := new(text.PostSpotCreateData)

	operations := []func(){
		func() { aa.MarketBuy("PIT_USDT", 13.1, 15) },
		func() { aa.LimitBuy("PIT_USDT", "0.0000000004885", 21470829068.6, 21470829088.6) },
		func() { aa.MarketSell("PIT_USDT", 21470829068.6, 21470829088.6) },
		func() { aa.LimitSell("PIT_USDT", "0.0000000004905", 21470829068.6, 21470829088.6) },
	}

	for _, operation := range operations {
		wg.Add(1)
		go executeOperation(&wg, operation)
	}

	// 等待所有goroutine完成
	wg.Wait()

	fmt.Println("所有操作已完成。")

	//aa.LimitBuy("BTC_USDT", "18582.89", 0.0006, 0.00066)
	//aa.LimitBuy("ETH_USDT", "980.65", 0.0102, 0.015)
	//aa.LimitBuy("BIT_USDT", "0.0000026182", 4000000, 4000008)

	//aa.MarketBuy("PIT_USDT", 10.1, 12)
	//aa.MarketSell("PIT_USDT", 22737115534.4, 45737115534.4)
	//aa.LimitSell("PIT_USDT", "0.000000000462", 22737115534.4, 45737115534.4)
}

func executeOperation(wg *sync.WaitGroup, operation func()) {
	defer wg.Done()
	operation()
}
