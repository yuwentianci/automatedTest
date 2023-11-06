package main

import (
	_ "github.com/go-sql-driver/mysql"
	"myapp/text"
	"time"
)

func main() {

	//dd := new(text.FuturesTrading)
	//var wg sync.WaitGroup
	//
	//// 使用WaitGroup等待goroutines完成。
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//	limitOpenLong(dd)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	limitOpenClose(dd)
	//}()
	//
	//// 在退出之前等待两个goroutines完成。
	//wg.Wait()
}

func limitOpenLong(dd *text.FuturesTrading) {
	if result := dd.LimitOpenLong("10", "41000", "100"); result == 0 {
		time.Sleep(3 * time.Second)
		marketOpenLong(dd)
	}
}

func marketOpenLong(dd *text.FuturesTrading) {
	if result := dd.MarketOpenLong("10", "100"); result == 0 {
		time.Sleep(2 * time.Second)
		limitOpenLong(dd)
	}
}

func limitOpenClose(dd *text.FuturesTrading) {
	if result := dd.LimitOpenClose("10", "40000", "100"); result == 0 {
		time.Sleep(2 * time.Second)
		marketOpenClose(dd)
	}
}

func marketOpenClose(dd *text.FuturesTrading) {
	if result := dd.MarketOpenClose("10", "100"); result == 0 {
		time.Sleep(3 * time.Second)
		limitOpenClose(dd)
	}
}
